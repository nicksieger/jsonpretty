# -*- encoding: utf-8 -*-

Gem::Specification.new do |s|
  s.name = %q{jsonpretty}
  s.version = "1.0.0"

  s.required_rubygems_version = Gem::Requirement.new(">= 0") if s.respond_to? :required_rubygems_version=
  s.authors = ["Nick Sieger"]
  s.date = %q{2009-02-20}
  s.default_executable = %q{jsonpretty}
  s.description = %q{Command-line JSON pretty-printer, using the json gem.}
  s.email = ["nick@nicksieger.com"]
  s.executables = ["jsonpretty"]
  s.extra_rdoc_files = ["History.txt", "Manifest.txt", "README.txt"]
  s.files = ["History.txt", "Manifest.txt", "README.txt", "Rakefile", "bin/jsonpretty", "lib/jsonpretty.rb"]
  s.has_rdoc = true
  s.homepage = %q{http://github.com/nicksieger/jsonpretty}
  s.rdoc_options = ["--main", "README.txt"]
  s.require_paths = ["lib"]
  s.rubyforge_project = %q{caldersphere}
  s.rubygems_version = %q{1.3.1}
  s.summary = %q{Command-line JSON pretty-printer, using the json gem.}

  if s.respond_to? :specification_version then
    current_version = Gem::Specification::CURRENT_SPECIFICATION_VERSION
    s.specification_version = 2

    if Gem::Version.new(Gem::RubyGemsVersion) >= Gem::Version.new('1.2.0') then
      s.add_development_dependency(%q<hoe>, [">= 1.8.2"])
    else
      s.add_dependency(%q<hoe>, [">= 1.8.2"])
    end
  else
    s.add_dependency(%q<hoe>, [">= 1.8.2"])
  end

  if defined?(JRUBY_VERSION)
    s.add_dependency('json_pure', '> 0')
  else
    s.add_dependency('json', '> 0')
  end

end
