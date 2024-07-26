#!/usr/bin/env ruby
# frozen_string_literal: true

require 'os'

require_relative 'cli'
require_relative 'gui'
require_relative 'input'

def load_programs(file_name)
  programs = []
  File.open(file_name, 'r') do |file|
    programs = file.readlines
  end
  programs
end

def read_programs(programs)
  programs.each do |program|
    puts program
  end
end

def install_program(program, progressbar, dry_run)
  command = "which #{program}"
  output = `#{command}`
  unless output.empty?
    puts "✅ #{program} already installed!"
    return
  end
  dry_run = if dry_run
              '--dry-run'
            else
              ''
            end
  if program == 'font-hack-nerd-font' && OS.linux?
    puts "'❌ Error: not available for Linux" 
    return
  end
  if program == 'podman-desktop' && OS.linux?
    if dry_run == '--dry-run'
      puts "✅ #{program} installed!"
      return
    end
    `flatpak install flathub io.podman_desktop.PodmanDesktop`
    puts "✅ #{program.chomp} installed!"
  end
  if %w[mitmproxy metasploit].include?(program)
    if OS.linux? && program == 'mitmproxy' and dry_run == ''
      `mkdir ~/mitmproxy`
      `curl -O -L https://downloads.mitmproxy.org/10.4.0/mitmproxy-10.4.0-linux-x86_64.tar.gz`
      `tar -xzf mitmproxy-10.4.0-linux-x86_64.tar.gz -C ~/mitmproxy`
      `rm mitmproxy-10.4.0-linux-x86_64.tar.gz`
      puts "✅ #{program.chomp} installed!"
      puts 'To run, run `ls ~/mitmproxy` and chose your way.'
      return
    elsif OS.linux? && program == 'mitmproxy' and dry_run == '--dry-run'
      puts "✅ #{program.chomp} installed!"
      return
    end
    if program == 'metasploit' && OS.linux? and dry_run == ''
      `curl https://raw.githubusercontent.com/rapid7/metasploit-omnibus/master/config/templates/metasploit-framework-wrappers/msfupdate.erb > msfinstall && \
  chmod 755 msfinstall && \
  ./msfinstall`
      `rm msfinstall`
    elsif program == 'metasploit' && OS.linux? and dry_run == '--dry-run'
      puts "✅ #{program.chomp} installed!"
      return
    end
    puts "✅ #{program.chomp} installed!"
    return
  end
  if program == 'amass'
    system('brew tap caffix/amass') do |output|
      print output
    end
    system("brew install #{dry_run} #{program.chomp}") do |output|
      print output
    end
    puts "✅ #{program.chomp} installed!"
    nil
  elsif program == 'browsh'
    system('brew tap browsh-org/homebrew-browsh') do |output|
      print output
    end
    system("brew install #{dry_run} #{program.chomp}") do |output|
      print output
    end
    puts "✅ #{program.chomp} installed!"
    nil
  elsif %w[alacritty imhex kitty wireshark nmap openvpn neovim].include?(program) && OS.linux?
    if dry_run == '--dry-run'
      puts "✅ #{program} installed!"
      return
    end
    command = 'which dnf'
    output = `#{command}`
    if output.empty?
      system("sudo apt install -y #{program}") do |output|
        print output
      end
      puts "✅ #{program} installed!"
      nil
    else
      system("sudo dnf install -y #{program}") do |output|
        print output
      end
      puts "✅ #{program} installed!"
      nil
    end
  elsif program == 'warp' && OS.linux?
    puts '❌ Error: Download Warp from https://www.warp.dev/ and install manually using dpkg (.deb) or dnf (.rpm).'
  elsif program == 'zed' && OS.linux?
    if dry_run == '--dry-run'
      puts "✅ #{program} installed!"
      return
    end
    `curl -f https://zed.dev/install.sh | sh`
    puts "✅ #{program} installed!"
  elsif program == 'ollama' && OS.linux?
    if dry_run == '--dry-run'
      puts "✅ #{program} installed!"
      return
    end
    `curl -fsSL https://ollama.com/install.sh | sh`
    puts "✅ #{program} installed!"
  else
    system("brew install #{dry_run} #{program.chomp}") do |output|
      print output
    end
    puts "✅ #{program.chomp} installed!"
  end
end

def uninstall_program(program, progressbar)
  if program == 'mitmproxy' && OS.linux?
    # Check if the directory exists and remove it
    `rm -rf ~/mitmproxy` if File.directory?('~/mitmproxy')
    puts "✅ #{program.chomp} uninstalled!"
    return
  end
  system("brew uninstall #{dry_run} #{program.chomp}") do |output|
    print output
  end
  if program == 'amass'
    system('brew untap owasp-amass/amass') do |output|
      print output
    end
  end
  if program == 'browsh'
    system('brew untap browsh-org/homebrew-browsh') do |output|
      print output
    end
  end
  puts "✅ #{program.chomp} uninstalled!"
end

def lates_version
  raw_varsion = `git -c 'versionsort.suffix=-' ls-remote --exit-code --refs --sort='version:refname' --tags https://github.com/NutekSecurity/nutek-apple.git '*.*.*' | tail --lines=1`
  raw_varsion.split('/')[2].chomp
end

def check_version_and_update
  response = `git status`
  if response.include? 'ahead'
    puts 'hint: You can run the script with --no-update to skip the update check.'
    on_your_own = get_yes_no_input "You're working on your own version of nutek-apple 🍎 Do you want to continue? (yes/no): "
    return if on_your_own

    puts 'Exit'
    exit
  end
  response = `git pull origin main --rebase`
  puts response
  if !$CHILD_STATUS.nil? && $CHILD_STATUS.success?
    if response.include?('error')
      puts '❌ Error: Could not update from the repository.'
      exit
    elsif response.include?('Already up to date.')
      puts '✅ Already up to date.'
    else
      puts '✅ Updated successfully. Please restart the script.'
      exit
    end
  else
    puts '❌ Error: Git command failed.'
    puts 'hint: You can run the script with --no-update to skip the update check.'
    exit
  end
end

def update(args)
  return if args.include?('--no-update')

  puts 'Updating...'
  # check if we're in a git repo
  if File.directory?('.git')
    # check if we're in the right repo
    if `git remote get-url origin`.chomp == 'https://github.com/NutekSecurity/nutek-apple.git' || `git remote get-url upstream`.chomp == 'https://github.com/NutekSecurity/nutek-apple.git'
      check_version_and_update
    else
      puts '❌ Error: Not in the right git repository, not updating.'
      puts 'hint: You can try to run the script with --no-update to skip the update check.'
      exit
    end
  else
    puts '❌ Error: Not in a git repoository, not updating.'
    puts 'hint: You can try to run the script with --no-update to skip the update check.'
    exit
  end
end

def get_command_line_arguments
  args = ARGV
  unless args.nil? 
    update(args)
  else
    update([])
  end
  uninstall_argument = false
  if args.empty? || args.include?('--help') || args.include?('-h') ||
     (args.length == 1 && args.include?('--no-update'))
    puts 'Usage: ruby nutek-apple.rb [options]'
    puts ''
    puts "Automated installation of hacking command line (and few GUI) programs on macOS (and Linux) - Nutek Security Platform. Requires Homebrew (also available for Linux).\nCurated by Nutek Security"
    puts "Auto-update enabled, just run the script from the repository directory.\nWill download the latest version from GitHub automatically:"
    puts 'https://github.com/NutekSecurity/nutek-apple'
    # Always update
    puts "\nOptions:"
    puts "  -h, --help\t\t\t\tShow this help message and exit"
    puts "  -i, --install\t\t\t\tInstall programs. Choose programs to install with --gui, --cli or --all"
    puts "  -u, --uninstall\t\t\tUninstall programs. Choose programs to uninstall with --gui, --cli or --all"
    puts "  --license\t\t\t\tShow license information"
    puts "  --all\t\t\t\t\tInstall or uninstall all programs"
    puts "  --cli\t\t\t\tInstall or uninstall cli set of programs"
    puts "  --gui\t\t\t\tInstall or uninstall GUI programs"
    puts "  --list\t\t\t\tList all programs"
    puts "  --list-cli\t\t\t\tList cli set of programs"
    puts "  --list-gui\t\t\t\tList gui programs"
    puts "  --dry-run\t\t\t\tDry run. Show what would be installed without actually installing anything and try to install without installing anything"
    puts "\nExamples:"
    puts '  ruby nutek-apple.rb --install --cli'
    puts '  ruby nutek-apple.rb --install --gui'
    puts '  ruby nutek-apple.rb --install --all'
    puts '  ruby nutek-apple.rb --uninstall --gui'
    puts '  ruby nutek-apple.rb --uninstall --all'
    puts '  ruby nutek-apple.rb --list-all'
    puts '  ruby nutek-apple.rb --list --gui'
    exit
  end
  if args.include?('--license')
    puts 'LICENSE'
    puts 'The user is granted a non-exclusive, perpetual license to use,'
    puts 'and distribute the software, subject to the following terms:'
    puts 'The user must display the following copyright notice in all copies of the'
    puts 'software:'
    puts ''
    puts 'Copyright (c) 2024 Szymon Błaszczyński'
    puts ''
    puts 'The user is not permitted to sublicense the software.'
    puts 'The user is not permitted to sell the software.'
    puts ''
    puts 'The above copyright notice and this permission notice shall be included in all'
    puts 'copies or substantial portions of the Software.'
    puts ''
    puts 'THE SOFTWARE IS PROVIDED \'AS IS\', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR'
    puts 'IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,'
    puts 'FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE'
    puts 'AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER'
    puts 'LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,'
    puts 'OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE'
    puts 'SOFTWARE.'
    exit
  end
  if args.include?('--list-all')
    puts 'Command Line Interface (terminal):'
    read_programs(cli)
    puts "\nGraphical User Interface (desktop):"
    read_programs(gui)
    exit
  end
  if args.include?('--list') && args.include?('--cli')
    puts 'Command Line Interface (terminal):'
    read_programs(cli)
    exit
  end
  if args.include?('--list') && args.include?('--gui')
    puts 'Graphical User Interface (desktop):'
    read_programs(gui)
    exit
  end
  uninstall_argument = true if args.include?('--uninstall') || args.include?('-u')
  uninstall_argument = false if args.include?('--install') || args.include?('-i')
  if args.include?('--uninstall') && args.include?('--install')
    puts '❌ Error: --uninstall and --install cannot be used together'
    exit false
  end
  if args.include?('-u') && args.include?('-i')
    puts '❌ Error: -u and -i cannot be used together'
    exit false
  end
  if args.include?('--uninstall') && args.include?('-i')
    puts '❌ Error: --uninstall and -i cannot be used together'
    exit false
  end
  if args.include?('-u') && args.include?('--install')
    puts '❌ Error: -u and --install cannot be used together'
    exit false
  end
  if args.include?('--uninstall') && args.include?('-u')
    puts '❌ Error: --uninstall and -u cannot be used together'
    exit false
  end
  if args.include?('-i') && args.include?('--install')
    puts '❌ Error: -i and --install cannot be used together'
    exit false
  end
  if args.include?('-u') && args.include?('--uninstall')
    puts '❌ Error: -u and --uninstall cannot be used together'
    exit false
  end
  if (args.include?('-u') || args.include?('--uninstall')) &&
     args.include?('--dry-run')
    puts '❌ Error: -u and --uninstall cannot be used with --dry-run'
    exit false
  end
  if (args.include?('--install') || args.include?('-i')) &&
     (!args.include?('--all') && !args.include?('--gui') &&
     !args.include?('--cli'))
    puts '❌ Error: -i and --install must be used with --all, --gui, --cli'
    exit false
  end
  if (args.include?('--uninstall') || args.include?('-u')) &&
     (!args.include?('--all') && !args.include?('--gui') &&
     !args.include?('--cli'))
    puts '❌ Error: -u and --uninstall must be used with --all, --gui, --cli'
    exit false
  end
  programs = []
  programs += cli if args.include?('--cli')
  programs += gui if args.include?('--gui')
  if args.include?('--all')
    programs = gui + cli
    # deduplicate programs
    programs = programs.uniq
  end
  unattended = false
  dry_run = if args.include?('--dry-run')
              true
            else
              false
            end
  if programs.empty?
    puts "❌ Error: No programs selected. Please select at least one program group to install or uninstall.\nView help❓ with -h or --help for more information."
    exit false
  end
  [uninstall_argument, programs, unattended, dry_run]
end

def check_if_homebrew_installed
  command = 'which brew'
  output = `#{command}`
  if output.empty?
    puts '❌ Homebrew is not installed. Please install Homebrew and try again. Visit https://brew.sh/ for more information.'
    exit false
  else
    puts '✅ Homebrew is installed.'
  end
end

def main
  uninstall_argument, programs,
  unattended, dry_run = get_command_line_arguments
  check_if_homebrew_installed
  # progressbar = ProgressBar.new(78)
  progressbar = nil
  if uninstall_argument
    answer = false
    if !unattended
      puts programs
      answer = get_yes_no_input 'Are you sure you want to uninstall all the above programs? (yes/no) '
    else
      answer = true
    end
    if answer
      programs.each do |program|
        puts "Uninstalling #{program}..."
        uninstall_program(program, progressbar)
        sleep(0.1)
      end
    else
      puts 'Uninstall aborted!'
    end
  else
    answer = false
    if !unattended
      puts programs
      answer = get_yes_no_input 'Are you sure you want to install all the above programs? (yes/no) '
    else
      answer = true
    end
    if answer
      programs.each do |program|
        puts "Installing #{program}..."
        install_program(program, progressbar, dry_run)
        sleep(0.1)
      end
    else
      puts 'Install aborted!'
    end
  end
  puts "\n Thank you for using nutek-apple 🍎 - the most important tools in hacker's backpack"
  puts ''
  puts 'https://nuteksecurity.com/'
  puts 'neosb@nuteksecurity.com'
end

main

####################
# End of Script
####################

####################
# Notes
####################
# This script is designed to be run on macOS, but as well may work on Linux.
# First run the script with --dry-run to see what would be installed, and to check if the script works on your system.
# Then run the script without --dry-run to install the programs.
# If you encounter any issues, please report them on GitHub:
# https://github.com/NutekSecurity/nutek-apple/issues
# Thank you for using Nutek Security Platform!
# https://nuteksecurity.com/
####################
# End of Notes
# End of nutek-apple.rb
####################
