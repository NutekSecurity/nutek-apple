#!/usr/bin/env ruby

# update this version number when updating the script
# this is used to check for updates
# tag the commit with the version number at the same time
$this_version = '0.1.9'

$gui = %w[
  podman-desktop
  imhex
  warp
  alacritty
]

$cli = %w[
  podman
  neovim
  openvpn
  irssi
  dos2unix
  ipcalc
  whatmask
  expect
  fd
  tmux
  lsd
  bat
  ripgrep-all
  sd
  termshark
  httpie
  smap
  nmap
  p0f
  masscan
  feroxbuster
  ffuf
  nuclei
  mitmproxy
  httpx
  amass
  jq
  htmlq
  tmux
  httrack
  monolith
  mdcat
  ouch
]

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
  dry_run = if dry_run
              '--dry-run'
            else
              ''
            end
  if program == 'metasploit'
    system('brew tap homebrew/cask') do |output|
      print output
    end
    system("brew install #{dry_run} --cask #{program.chomp}") do |output|
      print output
    end
    puts "✅ #{program.chomp} installed!"
  end
  if program == 'amass'
    system('brew tap caffix/amass') do |output|
      print output
    end
    system("brew install #{dry_run} #{program.chomp}") do |output|
      print output
    end
    puts "✅ #{program.chomp} installed!"
  elsif program == 'browsh'
    system('brew tap browsh-org/homebrew-browsh') do |output|
      print output
    end
    system("brew install #{dry_run} #{program.chomp}") do |output|
      print output
    end
    puts "✅ #{program.chomp} installed!"
  else
    system("brew install #{dry_run} #{program.chomp}") do |output|
      print output
    end
    puts "✅ #{program.chomp} installed!"
  end
end

def uninstall_program(program, progressbar)
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

def get_yes_no_input(prompt)
  print prompt
  answer = $stdin.gets
  answer = answer.chomp.downcase
  return true if %w[yes y].include?(answer)

  false
end

def lates_version
  raw_varsion = `git -c 'versionsort.suffix=-' ls-remote --exit-code --refs --sort='version:refname' --tags https://github.com/nutek-terminal/nutek-apple.git '*.*.*' | tail --lines=1`
  raw_varsion.split('/')[2].chomp
end

def update(args, github_version)
  return if args.include?('--no-update')

  puts 'Checking for updates...'
  if github_version != $this_version
    puts "New version available: #{github_version}"
    puts 'Updating...'
    # check if we're in a git repo
    if File.directory?('.git')
      # check if we're in the right repo
      if `git remote get-url origin` == 'https://github.com/nutek-terminal/nutek-apple.git'
        system('git pull origin main')
        puts "Updated to version #{github_version}"
        exit
      else
        puts '❌ Error: Not in the right repo, not updating.'
      end
    else
      puts '❌ Error: Not in a git repo, not updating.'
    end
  else
    puts "✅ Already up to date.\n"
  end
end

def get_command_line_arguments
  args = ARGV
  uninstall_argument = false
  github_version = lates_version
  if args.length == 0 || args.include?('--help') || args.include?('-h')
    puts 'Usage: ruby nutek-apple.rb [options]'
    puts "Automated installation of hacking command line programs on macOS - Nutek Security Platform. Requires Homebrew.\nCurated by Nutek Security Solutions\n\tand Szymon Błaszczyński."
    puts 'Download the latest version from GitHub:'
    puts 'https://github.com/NutekSecurity/nutek-apple'
    puts "This version: #{$this_version} GitHub version: #{github_version}"
    if $this_version != github_version
      puts "New version available: #{github_version} do you want to update?"
      update(args, github_version) if get_yes_no_input('Update? [Y/n] ')
    end
    puts "\nOptions:"
    puts "  -h, --help\t\t\t\tShow this help message and exit"
    puts "  -i, --install\t\t\t\tInstall programs. Choose programs to install with --gui, --cli or --all"
    puts "  -u, --uninstall\t\t\tUninstall programs. Choose programs to uninstall with --gui, --cli or --all"
    puts "  --web, --safari\t\t\tOpen Safari and lookup nuteksecurity.com"
    puts "  --license\t\t\t\tShow license information"
    puts "  --all\t\t\t\t\tInstall or uninstall all programs"
    puts "  --cli\t\t\t\tInstall or uninstall cli set of programs"
    puts "  --gui\t\t\t\tInstall or uninstall GUI programs"
    puts "  --list\t\t\t\tList all programs"
    puts "  --list-cli\t\t\t\tList cli set of programs"
    puts "  --list-gui\t\t\t\tList gui programs"
    puts "  --unattended\t\t\t\tUnattended mode. Install selected programs without asking for confirmation"
    puts "  --dry-run\t\t\t\tDry run. Show what would be installed without actually installing anything"
    puts "\nExamples:"
    puts '  ruby nutek-apple.rb --install --cli'
    puts '  ruby nutek-apple.rb --install --gui'
    puts '  ruby nutek-apple.rb --install --all'
    puts '  ruby nutek-apple.rb --uninstall --gui'
    puts '  ruby nutek-apple.rb --uninstall --all'
    puts '  ruby nutek-apple.rb --list'
    puts '  ruby nutek-apple.rb --list-gui'
    puts '  ruby nutek-apple.rb --unattended --install --gui'
    puts '  ruby nutek-apple.rb --unattended --install --all'
    puts '  ruby nutek-apple.rb --unattended --uninstall --gui'
    puts '  ruby nutek-apple.rb --unattended --uninstall --all'
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
  if args.include?('--safari') || args.include?('--web')
    system('open -a Safari https://nuteksecurity.com/')
    exit
  end
  update(args, github_version)
  if args.include?('--list')
    puts 'cli:'
    read_programs($cli)
    puts "\nGUI:"
    read_programs($gui)
    exit
  end
  if args.include?('--list-cli')
    puts 'cli:'
    read_programs($cli)
    exit
  end
  if args.include?('--list-gui')
    puts 'GUI:'
    read_programs($gui)
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
  programs += $cli if args.include?('--cli')
  programs += $gui if args.include?('--gui')
  if args.include?('--all')
    programs = $gui + $cli
    # deduplicate programs
    programs = programs.uniq
  end
  unattended = if args.include?('--unattended')
                 true
               else
                 false
               end
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
  puts "\nFor future development and security awarness help me with Monero:"
  puts 'Monero address: 87G8nLBPdwAEPycmWWAhUhZC8kUuuFgjX8zEUw1VjvNMPdkUWzxikocQyLtycwqzJfChR5bNVyXU87m5vT4Fy9gtS6Q5X8L'
  puts 'or Bitcoin:'
  puts 'Bitcoin address: 3AhSZUecGQDk97iCGtUtCq3kqCdndsZEF1'
end

main

####################
# End of Script
####################

####################
# Notes
####################
# This script is designed to be run on macOS.
# This script is tested to run with Ruby 3.2.2 / 2.6.10p210 and Homebrew 4.0.28
