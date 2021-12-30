package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

type JsonArg struct {
	file    *os.File
	content []byte
}

func (j JsonArg) Close() error {
	if j.file != nil {
		return j.file.Close()
	}
	return nil
}

func usage(code int) {
	fmt.Printf("usage: %s [args|@filename|@- (stdin)]\n", os.Args[0])
	fmt.Printf("Parse and pretty-print JSON, either from stdin or from arguments concatenated together\n")
	os.Exit(code)
}

func versionInfo() string {
	if commit != "" && date != "" {
		return fmt.Sprintf("%s (%s on %s)", version, commit[0:7], date[0:10])
	}
	return version
}

func checkFlags(arg string) {
	if arg[0] == '-' && len(arg) > 1 {
		code := 1
		if arg == "-v" || strings.HasPrefix(arg, "--v") {
			fmt.Printf("jsonpretty version %s\n", versionInfo())
			os.Exit(0)
		} else if arg == "-h" || strings.HasPrefix(arg, "--h") {
			code = 0
		}
		if code == 1 {
			fmt.Printf("unrecognized flag %s\n", arg)
		}
		usage(code)
	}
}

func buildJsonSource(args []JsonArg) *bytes.Buffer {
	buffer := bytes.NewBuffer(nil)
	for _, arg := range args {
		if arg.file != nil {
			_, err := buffer.ReadFrom(arg.file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
			}
		} else {
			_, err := buffer.Write(arg.content)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error appending content: %v\n", err)
			}
		}
	}
	return buffer
}

func readArguments(args []string) *bytes.Buffer {
	jsonArgs := []JsonArg{}
	defer func() {
		for _, j := range jsonArgs {
			_ = j.Close()
		}
	}()
	for _, arg := range args {
		checkFlags(arg)

		if arg == "-" || arg == "@-" {
			jsonArgs = append(jsonArgs, JsonArg{file: os.Stdin})
		} else if strings.HasPrefix(arg, "@") {
			file, err := os.Open(arg[1:])
			if err != nil {
				fmt.Fprintf(os.Stderr, "skipping '%s' due to error: %v\n", arg, err)
				continue
			}
			jsonArgs = append(jsonArgs, JsonArg{file: file})
		} else {
			jsonArgs = append(jsonArgs, JsonArg{content: []byte(arg)})
		}
	}

	if len(jsonArgs) == 0 {
		jsonArgs = append(jsonArgs, JsonArg{file: os.Stdin})
	}
	return buildJsonSource(jsonArgs)
}

var httpRegex *regexp.Regexp = regexp.MustCompile("^HTTP/\\d")

func parseHeaders(src *bytes.Buffer) ([]byte, []byte, error) {
	headerBuffer := bytes.NewBuffer(nil)
	line, err := src.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return nil, nil, err
	}

	if httpRegex.Match(line) {
		for {
			_, err = headerBuffer.Write(line)
			if err != nil {
				return nil, nil, err
			}
			if (len(line) == 1 && line[0] == '\n') ||
				(len(line) == 2 && line[0] == '\r' && line[1] == '\n') {
				break
			}

			line, err = src.ReadBytes('\n')

			if err == io.EOF {
				_, err = headerBuffer.Write(line)
				if err != nil {
					return nil, nil, err
				}
				break
			}

			if err != nil && err != io.EOF {
				return nil, nil, err
			}
		}
	} else {
		newsrc := bytes.NewBuffer(line)
		_, err = newsrc.Write(src.Bytes())
		if err != nil {
			return nil, nil, err
		}
		src = newsrc
	}
	return headerBuffer.Bytes(), src.Bytes(), nil
}

var jsonpRegexp *regexp.Regexp = regexp.MustCompile("^([a-zA-Z$_][^ \t\r\n(]+)\\(")
var endingParen *regexp.Regexp = regexp.MustCompile("\\)[ \t\r\n]*$")

func cleanJsonp(jsonSrc []byte) ([]byte, string) {
	match := jsonpRegexp.FindSubmatch(jsonSrc)
	ending := endingParen.FindSubmatch(jsonSrc)
	if match != nil && ending != nil {
		return jsonSrc[len(match[0]) : len(jsonSrc)-len(ending[0])], string(match[1])
	}
	return jsonSrc, ""
}

func handleParseError(err error, jsonSrc []byte) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing json: %v\ninput:\n%s\n", err, string(jsonSrc))
		os.Exit(1)
	}
}

func main() {
	buffer := readArguments(os.Args[1:])
	headers, jsonSrc, err := parseHeaders(buffer)
	handleParseError(err, jsonSrc)

	jsonSrc, jsonpName := cleanJsonp(jsonSrc)

	if len(headers) > 0 {
		fmt.Print(string(headers))
	}

	if jsonpName != "" {
		fmt.Printf("jsonp method name: %s\n\n", jsonpName)
	}

	var jsonObj interface{}
	err = json.Unmarshal(jsonSrc, &jsonObj)
	handleParseError(err, jsonSrc)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err = enc.Encode(&jsonObj)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encoding json: %v\n", err)
		os.Exit(1)
	}
}
