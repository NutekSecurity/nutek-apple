# frozen_string_literal: true

require 'os'

def programming
  if OS.linux?
    if `which dnf`.empty?
      system('sudo apt install build-essential') do |output|
        print output
      end
    else
      system('sudo dnf groupinstall "Development Tools"') do |output|
        print output
      end
    end
  end
  if `which pyenv`.empty?
    system('brew install pyenv') do |output|
      print output
    end
    puts 'look up for instructions to setup pyenv and press enter to continue...'
    $stdin.gets
    system('brew install pyenv-virtualenv') do |output|
      print output
    end
    puts 'look up for instructions to setup pyenv-virtualenv and press enter to continue...'
    $stdin.gets
  end

  # puts 'press enter to continue...'
  # $stdin.gets
  system('brew install node') do |output|
    print output
  end
end
