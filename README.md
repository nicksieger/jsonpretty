# jsonpretty

- http://github.com/nicksieger/jsonpretty

## DESCRIPTION

Command-line JSON pretty-printer. Pipe any JSON output to pretty-print it with 2-space indent.

## FEATURES

- Parse and pretty-print JSON/JSONP either from stdin or from command-line arguments.
- All arguments are concatenated together in a single string for pretty-printing.
- Use `@filename` as an argument to include the contents of the file.
- Use `-` or `@-` as an argument (or use no arguments) to read stdin.
- Detects HTTP response/headers, prints them untouched, and skips to the body (for use with `curl -i').

## SYNOPSIS

```
curl -i http://api.com/json | jsonpretty
```

## REQUIREMENTS

- For the Ruby version: any Ruby version >= 2 installed
- For the Go version: none for the release binary, or Go 1.17 for `go install`

## INSTALL

Install the Ruby version with rubygems:

```
gem install jsonpretty
```

Install the Ruby version globally with Homebrew and brew-gem:

```
brew install brew-gem
brew gem install jsonpretty
```

Install the Go version with `go install` (requires Go 1.17):

```
go install github.com/nicksieger/jsonpretty
```

Or install the Go version from a release with curl+tar:

```
VERSION=1.2.0
OS=$(uname -s)
ARCH=$(uname -m)
curl -L https://github.com/nicksieger/jsonpretty/releases/download/v${VERSION}/jsonpretty_${VERSION}_${OS}_${ARCH}.tar.gz | sudo tar -C /usr/local/bin -zxf - jsonpretty
```

## LICENSE

See [LICENSE](LICENSE.txt).
