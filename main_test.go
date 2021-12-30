package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

type Testfile struct {
	name    string
	content string
	file    *os.File
}

var stdinSaved *os.File = os.Stdin

func Setup(contents []string, t *testing.T) []Testfile {
	var err error
	os.Stdin, err = os.CreateTemp("", "stdin")
	if err != nil {
		t.FailNow()
	}

	testfiles := make([]Testfile, len(contents))
	for i, c := range contents {
		testfiles[i] = Testfile{content: c}
		file, err := os.CreateTemp("", fmt.Sprintf("tf%d", i))
		if err != nil {
			t.Logf("unable to create tempfile for '%s'", c)
			t.FailNow()
		}
		_, err = file.Write([]byte(c))
		if err != nil {
			t.Logf("unable to write tempfile with '%s'", c)
			t.FailNow()
		}
		err = file.Close()
		if err != nil {
			t.Logf("unable to close tempfile for '%s'", c)
			t.FailNow()
		}
		testfiles[i].file = file
		testfiles[i].name = file.Name()
	}
	return testfiles
}

func TearDown(testfiles []Testfile) {
	for _, tf := range testfiles {
		_ = os.Remove(tf.name)
	}
	_ = os.Stdin.Close()
	os.Stdin = stdinSaved
}

func assertEqual(expect []byte, actual []byte, t *testing.T) {
	if bytes.Compare(expect, actual) != 0 {
		t.Errorf("bytes did not match:\n. expected:\n  %s\n  actual:\n  %s\n",
			string(expect), string(actual))
	}
}

func TestReadArguments(t *testing.T) {
	files := Setup([]string{"1", "2", "{\"foo\": true, \"bar\": false}"}, t)
	defer TearDown(files)

	tests := [][]Testfile{
		{}, // nothing
		{Testfile{content: "{\"hello\":\"world\"}"}},
		{Testfile{content: "["}, files[0], Testfile{content: ","}, files[1], Testfile{content: "]"}},
		{files[2]},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			expected := bytes.NewBuffer(nil)
			args := make([]string, len(test))
			for j, testfile := range test {
				if testfile.name != "" {
					args[j] = fmt.Sprintf("@%s", testfile.name)
				} else {
					args[j] = testfile.content
				}
				expected.WriteString(testfile.content)
			}
			expect := expected.Bytes()
			actual := readArguments(args)
			assertEqual(expect, actual.Bytes(), t)
		})
	}
}

func TestReadArgumentsStdin(t *testing.T) {
	f := Setup([]string{}, t)
	defer TearDown(f)

	expect := []byte("{\"foo\": true, \"bar\": false}")
	_, err := os.Stdin.Write(expect)
	if err != nil {
		t.Logf("unable to write to stdin\n")
		t.FailNow()
	}
	_ = os.Stdin.Close()

	reopen := func(t *testing.T) {
		f, err := os.Open(os.Stdin.Name())
		if err != nil {
			t.Logf("unable to reopen stdin")
			t.FailNow()
		}
		os.Stdin = f
	}

	t.Run("no args", func(t *testing.T) {
		reopen(t)
		actual := readArguments([]string{})
		assertEqual(expect, actual.Bytes(), t)
	})

	t.Run("-", func(t *testing.T) {
		reopen(t)
		actual := readArguments([]string{"-"})
		assertEqual(expect, actual.Bytes(), t)
	})

	t.Run("@-", func(t *testing.T) {
		reopen(t)
		actual := readArguments([]string{"@-"})
		assertEqual(expect, actual.Bytes(), t)
	})
}

func TestParseHeaders(t *testing.T) {
	tests := [][]string{
		{"", "", ""},
		{"{}", "", "{}"},
		{"not http\n\nhello\nworld", "", "not http\n\nhello\nworld"},
		{"HTTP/1.1 200 OK\nContent-Type: application/json\nContent-Length: 19\n\n{\"hello\":\"world\"}\n",
			"HTTP/1.1 200 OK\nContent-Type: application/json\nContent-Length: 19\n\n",
			"{\"hello\":\"world\"}\n"},
		{"HTTP/1.1 200 OK\nContent-Type: application/json\nContent-Length: 0\n\n",
			"HTTP/1.1 200 OK\nContent-Type: application/json\nContent-Length: 0\n\n",
			""},
		{"myfunc(1,2,3)", "", "myfunc(1,2,3)"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			buffer := bytes.NewBuffer(nil)
			buffer.WriteString(test[0])
			headers, jsonsrc, err := parseHeaders(buffer)
			if err != nil {
				t.Logf("error parseHeaders: %v", err)
				t.FailNow()
			}
			assertEqual([]byte(test[1]), headers, t)
			assertEqual([]byte(test[2]), jsonsrc, t)
		})
	}
}

func TestCleanJsonp(t *testing.T) {
	tests := [][]string{
		{"jsonp(true)", "true", "jsonp"},
		{"true", "true", ""},
		{"myJSFunc({\"hello\":\"world\"})", "{\"hello\":\"world\"}", "myJSFunc"},
		{"[1,2,3]", "[1,2,3]", ""},
		{"myfunc(1,2,3)", "1,2,3", "myfunc"},
		{"myfunc(1,2,3)\n", "1,2,3", "myfunc"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			jsonSrc, jsonpName := cleanJsonp([]byte(test[0]))
			assertEqual([]byte(test[1]), jsonSrc, t)
			assertEqual([]byte(test[2]), []byte(jsonpName), t)
		})
	}
}
