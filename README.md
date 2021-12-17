# jsonpretty

- http://github.com/nicksieger/jsonpretty

## DESCRIPTION

Command-line JSON pretty-printer, using the json gem.

## FEATURES/PROBLEMS

- Parse and pretty-print JSON/JSONP either from stdin or from command-line
  arguments.
- All arguments are concatenated together in a single string for
  pretty-printing.
- Use '@filename' as an argument to include the contents of the file.
- Use '-' or '@-' as an argument (or use no arguments) to read stdin.
- Detects HTTP response/headers, prints them untouched, and skips to
  the body (for use with `curl -i').

## SYNOPSIS

```
curl -i http://api.com/json | jsonpretty
```

## REQUIREMENTS

json or json_pure

## INSTALL

```
gem install jsonpretty
```

or install globally with brew-gem:

```
brew install brew-gem
brew gem install jsonpretty
```

## LICENSE

See [LICENSE](LICENSE).
