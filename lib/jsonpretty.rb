require 'json'

class Jsonpretty
  def file_value(filename)
    file = if filename == '-'
             $stdin
           else
             File.open(filename)
           end
    lines = file.readlines
    if lines.first =~ /^HTTP\/\d/ # looks like an HTTP response; we just want the body
      index = lines.index("\r\n") || lines.index("\n")
      puts lines[0..index]
      lines[(index+1)..-1].join('')
    else
      lines.join('')
    end
  ensure
    file.close
  end

  def stdin_value
    file_value('-')
  end

  def clean_jsonp(input_string)
    match = input_string.match(/^(.+)\((.+)\)$/)
    if match
      puts "jsonp method name: #{match[1]}\n\n"
      match[2]
    else
      input_string
    end
  end

  def main
    if ARGV.length == 0
      ARGV.unshift stdin_value
    else
      ARGV.each_with_index do |v,i|
        case v
        when /^--?h/
          puts("usage: #{File.basename($0)} [args|@filename|@- (stdin)]",
               "Parse and pretty-print JSON, either from stdin or from arguments concatenated together")
          exit 0
        when /^--?v/
          puts "jsonpretty version #{VERSION}"
          exit 0
        else
          if v == '-'
            ARGV[i] = stdin_value
          elsif v =~ /^@/
            ARGV[i] = file_value(v[1..-1])
          end
        end
      end
    end

    input = clean_jsonp(ARGV.join(' '))
    json = JSON.parse(input)
    puts JSON.pretty_generate(json)
  rescue => e
    $stderr.puts "jsonpretty failed: #{e.message}"
    exit 1
  end
end

