#!/usr/bin/env ruby

def fuzz(input_file)
  line_number = File.readlines(input_file).size

  while line_number > 0
    line = File.readlines(input_file).take(line_number).last

    ruby_command = "ruby nutek-apple.rb #{line}"
    system(ruby_command)

    line_number -= 1
  end
end

input_file = ".input-r2"
fuzz(input_file)