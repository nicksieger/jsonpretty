= jsonpretty

* http://github.com/nicksieger/jsonpretty

== DESCRIPTION:

Command-line JSON pretty-printer, using the json gem.

== FEATURES/PROBLEMS:

- Parse and pretty-print JSON/JSONP either from stdin or from command-line
  arguments.
- All arguments are concatenated together in a single string for
  pretty-printing.
- Use '@filename' as an argument to include the contents of the file.
- Use '-' or '@-' as an argument (or use no arguments) to read stdin.
- Detects HTTP response/headers, prints them untouched, and skips to
  the body (for use with `curl -i').

== SYNOPSIS:


== REQUIREMENTS:

json or json_pure

== INSTALL:

gem install jsonpretty

== LICENSE:

(The MIT License)

Copyright (c) 2007-2009 Nick Sieger <nick@nicksieger.com>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
