# -*- ruby -*-

require 'rubygems'
require 'hoe'
require './lib/jsonpretty.rb'

h = Hoe.new('jsonpretty', Jsonpretty::VERSION) do |p|
  p.rubyforge_name = 'caldersphere'
  p.developer('Nick Sieger', 'nick@nicksieger.com')
  p.url = 'http://github.com/nicksieger/jsonpretty'
  p.extra_deps << ['json', '> 0'] if ENV["VERSION"]
end
spec = h.spec
spec.dependencies.delete_if { |dep| dep.name == "hoe" }
def spec.to_ruby
  additional_src = %{
  if defined?(JRUBY_VERSION)
    s.add_dependency('json_pure', '> 0')
  else
    s.add_dependency('json', '> 0')
  end
}
  super.sub(/end\n\Z/m, "#{additional_src}\nend\n")
end

task :gemspec do
  File.open("jsonpretty.gemspec", "w") {|f| f << spec.to_ruby }
end
