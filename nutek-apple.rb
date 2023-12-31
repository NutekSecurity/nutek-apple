#!/usr/bin/env ruby

# update this version number when updating the script
# this is used to check for updates
# tag the commit with the version number at the same time
$this_version = "0.1.7"

$attack = [
  "bettercap",
  "ettercap",
  "hashcat",
  "hydra",
  "john-jumbo",
  #"metasploit",
  "mitmproxy",
  "ncrack",
  "socat",
  "sqlmap",
  "xh",
]

$code = [
  "adns",
  "bvi",
  "capstone",
  "cflow",
  "cppcheck",
  "fmt",
  "gitleaks",
  "grpc",
  "hyperscan",
  "jansson",
  "radare2",
  "rats",
  "zzuf",
]

$knowledge = [
  "amass",
  "arp-scan",
  "bmon",
  "dnsx",
  "feroxbuster",
  "ffuf",
  "fping",
  "gau",
  "gobuster",
  "httpx",
  "httrack",
  "masscan",
  "mtr",
  "nethogs",
  "ngrep",
  "nikto",
  "nmap",
  "nuclei",
  "p0f",
  "rustscan",
  "smap",
  "tcpdump",
  "tcpflow",
  "termshark",
  "testssl",
  "zmap",
]

$utility = [
  "argon2",
  "bat",
  "bottom",
  "discount",
  "dos2unix",
  "dust",
  "exa",
  "expect",
  "hexyl",
  "htmlq",
  "ipcalc",
  "irssi",
  "jq",
  "jql",
  "macchina",
  "mdcat",
  "monolith",
  "neovim",
  "openvpn",
  "ouch",
  "pandoc",
  "podman",
  "podman-compose",
  "ripgrep-all",
  "sd",
  "tmux",
  "viu",
  "w3m",
  "whatmask",
]

$mini = [
  "mitmproxy",
  "httpie",
  "smap",
  "termshark",
  "gau",
  "httpx",
  "feroxbuster",
  "ffuf",
  "amass",
  "ripgrep-all",
  "sd",
  "ouch",
  "jq",
  "htmlq",
  "exa",
  "bat",
  "warp",
  "alacritty",
  "tmux",
  "fd",
  "httrack",
  "monolith",
  "podman-desktop",
  "nmap",
  "nuclei",
  "w3m",
  "felinks",
]

def load_programs(file_name)
  programs = []
  File.open(file_name, "r") do |file|
      programs = file.readlines
  end
  return programs
end

def read_programs(programs)
  programs.each do |program|
      puts program
  end
end

def install_program(program, progressbar, dry_run)
  if dry_run
    dry_run = "--dry-run"
  else
    dry_run = ""
  end
  if program == "metasploit"
    system("brew tap homebrew/cask") do |output|
      print output
    end
    system("brew install #{dry_run} --cask #{program.chomp}") do |output|
      print output
    end
    puts "✅ #{program.chomp} installed!"
  end
  if program == "amass"
    system("brew tap caffix/amass") do |output|
      print output
    end
    system("brew install #{dry_run} #{program.chomp}") do |output|
      print output
    end
    puts "✅ #{program.chomp} installed!"
  elsif program == "browsh"
    system("brew tap browsh-org/homebrew-browsh") do |output|
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
    system("brew untap owasp-amass/amass") do |output|
      print output
    end
  end
  if program == 'browsh'
    system("brew untap browsh-org/homebrew-browsh") do |output|
      print output
    end
  end
  puts "✅ #{program.chomp} uninstalled!"
end

def get_yes_no_input(prompt)
  print prompt
  answer = $stdin.gets
  answer = answer.chomp.downcase
  if answer == "yes" || answer == "y"
    return true
  else
    return false
  end
end

def lates_version()
  raw_varsion = `git -c 'versionsort.suffix=-' ls-remote --exit-code --refs --sort='version:refname' --tags https://github.com/nutek-terminal/nutek-apple.git '*.*.*' | tail --lines=1`
  version = raw_varsion.split("/")[2].chomp
  return version
end

def update(args, github_version)
  if not args.include?("--no-update")
    puts "Checking for updates..."
    if github_version != $this_version
      puts "New version available: #{github_version}"
      puts "Updating..."
      # check if we're in a git repo
      if File.directory?(".git")
        # check if we're in the right repo
        if `git remote get-url origin` == "https://github.com/nutek-terminal/nutek-apple.git"
          system("git pull origin main")
          puts "Updated to version #{github_version}"
          exit
        else
          puts "❌ Error: Not in the right repo, not updating."
        end
      else
        puts "❌ Error: Not in a git repo, not updating."
      end
    else
      puts "✅ Already up to date.\n"
    end
  end
end

def get_command_line_arguments
  args = ARGV
  uninstall_argument = false
  github_version = lates_version()
  if args.length == 0 || args.include?("--help") || args.include?("-h")
    print("\033[1;31m");
    print("::::    ::: :::    ::: ::::::::::: :::::::::: :::    :::\n");
    print("\033[1;31m");
    print(":+:+:   :+: :+:    :+:     :+:     :+:        :+:   :+:\n");
    print("\033[1;32m");
    print(":+:+:+  +:+ +:+    +:+     +:+     +:+        +:+  +:+\n");
    print("\033[1;32m");
    print("+#+ +:+ +#+ +#+    +:+     +#+     +#++:++#   +#++:++ \n");
    print("\033[1;32m");    
    print("+#+  +#+#+# +#+    +#+     +#+     +#+        +#+  +#+\n");
    print("\033[1;34m");
    print("#+#   #+#+# #+#    #+#     #+#     #+#        #+#   #+#\n");
    print("\033[1;34m");
    print("###    ####  ########      ###     ########## ###    ###\n");
    print("\033[0m");
    puts "Usage: ruby nutek-apple.rb [options]"
    puts "Automated installation of hacking command line programs on macOS - Nutek Security Platform. Requires Homebrew.\nCurated by Nutek Security Solutions\n\tand Szymon Błaszczyński."
    puts "Download the latest version from GitHub:"
    puts "https://github.com/nutek-terminal/nutek-apple"
    puts "This version: #{$this_version} GitHub version: #{github_version}"
    if $this_version != github_version
      puts "New version available: #{github_version} do you want to update?"
      if get_yes_no_input("Update? [Y/n] ")
        update(args, github_version)
      end
    end
    puts "\nOptions:"
    puts "  -h, --help\t\t\t\tShow this help message and exit"
    puts "  -i, --install\t\t\t\tInstall programs. Choose programs to install with --attack, --code, --knowledge, --utility or --all"
    puts "  -u, --uninstall\t\t\tUninstall programs. Choose programs to uninstall with --attack, --code, --knowledge, --utility or --all"
    puts "  --web, --safari\t\t\tOpen Safari and lookup Nutek Security Platform website with all programs and tools explained"
    puts "  --all\t\t\t\t\tInstall or uninstall all programs"
    puts "  --mini\t\t\t\tInstall or uninstall mini set of programs"
    puts "  --attack\t\t\t\tInstall or uninstall attack programs"
    puts "  --code\t\t\t\tInstall or uninstall code programs"
    puts "  --knowledge\t\t\t\tInstall or uninstall knowledge programs"
    puts "  --utility\t\t\t\tInstall or uninstall utility programs"
    puts "  --list\t\t\t\tList all programs"
    puts "  --list-mini\t\t\t\tList mini set of programs"
    puts "  --list-attack\t\t\t\tList attack programs"
    puts "  --list-code\t\t\t\tList code programs"
    puts "  --list-knowledge\t\t\tList knowledge programs"
    puts "  --list-utility\t\t\tList utility programs"
    puts "  --unattended\t\t\t\tUnattended mode. Install selected programs without asking for confirmation"
    puts "  --dry-run\t\t\t\tDry run. Show what would be installed without actually installing anything"
    puts "\nExamples:"
    puts "  ruby nutek-apple.rb --install --mini"
    puts "  ruby nutek-apple.rb --install --attack --utility --knowledge"
    puts "  ruby nutek-apple.rb --install --all"
    puts "  ruby nutek-apple.rb --uninstall --attack --code"
    puts "  ruby nutek-apple.rb --uninstall --all"
    puts "  ruby nutek-apple.rb --list"
    puts "  ruby nutek-apple.rb --list-attack"
    puts "  ruby nutek-apple.rb --unattended --install --attack"
    puts "  ruby nutek-apple.rb --unattended --install --all"
    puts "  ruby nutek-apple.rb --unattended --uninstall --attack"
    puts "  ruby nutek-apple.rb --unattended --uninstall --all"
    puts "\nFor more information, see"
    puts "https://nutek.neosb.net/docs/tools/bing-search/ and TOOLS section where you can find more information about each tool."
    exit
  end
  if args.include?("--safari") || args.include?("--web")
    system("open -a Safari https://nutek.neosb.net/")
    exit
  end
  update(args, github_version)
  if args.include?("--list")
    puts "Mini:"
    read_programs($mini)
    puts "\nAttack:"
    read_programs($attack)
    puts "\nCode:"
    read_programs($code)
    puts "\nKnowledge:"
    read_programs($knowledge)
    puts "\nUtility:"
    read_programs($utility)
    puts "\nFor more information, see"
    puts "https://nutek.neosb.net/docs/tools/bing-search/ and TOOLS section where you can find more information about each tool."
    exit
  end
  if args.include?("--list-mini")
    puts "Mini:"
    read_programs($mini)
    puts "\nFor more information, see"
    puts "https://nutek.neosb.net/docs/tools/bing-search/ and TOOLS section where you can find more information about each tool."
    exit
  end
  if args.include?("--list-attack")
    puts "Attack:"
    read_programs($attack)
    puts "\nFor more information, see"
    puts "https://nutek.neosb.net/docs/tools/bing-search/ and TOOLS section where you can find more information about each tool."
    exit
  end
  if args.include?("--list-code")
    puts "Code:"
    read_programs($code)
    puts "\nFor more information, see"
    puts "https://nutek.neosb.net/docs/tools/bing-search/ and TOOLS section where you can find more information about each tool."
    exit
  end
  if args.include?("--list-knowledge")
    puts "Knowledge:"
    read_programs($knowledge)
    puts "\nFor more information, see"
    puts "https://nutek.neosb.net/docs/tools/bing-search/ and TOOLS section where you can find more information about each tool."
    exit
  end
  if args.include?("--list-utility")
    puts "Utility:"
    read_programs($utility)
    puts "\nFor more information, see"
    puts "https://nutek.neosb.net/docs/tools/bing-search/ and TOOLS section where you can find more information about each tool."
    exit
  end
  if args.include?("--uninstall") || args.include?("-u")
    uninstall_argument = true
  end
  if args.include?("--install") || args.include?("-i")
    uninstall_argument = false
  end
  if args.include?("--uninstall") && args.include?("--install")
    puts "❌ Error: --uninstall and --install cannot be used together"
    exit false
  end
  if args.include?("-u") && args.include?("-i")
    puts "❌ Error: -u and -i cannot be used together"
    exit false
  end
  if args.include?("--uninstall") && args.include?("-i")
    puts "❌ Error: --uninstall and -i cannot be used together"
    exit false
  end
  if args.include?("-u") && args.include?("--install")
    puts "❌ Error: -u and --install cannot be used together"
    exit false
  end
  if args.include?("--uninstall") && args.include?("-u")
    puts "❌ Error: --uninstall and -u cannot be used together"
    exit false
  end
  if args.include?("-i") && args.include?("--install")
    puts "❌ Error: -i and --install cannot be used together"
    exit false
  end
  if args.include?("-u") && args.include?("--uninstall")
    puts "❌ Error: -u and --uninstall cannot be used together"
    exit false
  end
  if ( args.include?("-u") || args.include?("--uninstall") ) &&
    args.include?("--dry-run")
    puts "❌ Error: -u and --uninstall cannot be used with --dry-run"
    exit false
  end
  if (args.include?("--install") || args.include?("-i")) &&
    ( !args.include?("--all") && !args.include?("--attack") && 
    !args.include?("--utility") && !args.include?("--code") && 
    !args.include?("--knowledge") && !args.include?("--mini"))
    puts "❌ Error: -i and --install must be used with --all, --attack, --utility, --code, or --knowledge"
    exit false
  end
  if (args.include?("--uninstall") || args.include?("-u")) &&
    ( !args.include?("--all") && !args.include?("--attack") &&
    !args.include?("--utility") && !args.include?("--code") &&
    !args.include?("--knowledge") && !args.include?("--mini"))
    puts "❌ Error: -u and --uninstall must be used with --all, --attack, --utility, --code, or --knowledge"
    exit false
  end
  programs = []
  if args.include?("--mini")
    programs += $mini
  end
  if args.include?("--attack")
    programs += $attack
  end
  if args.include?("--utility")
    programs += $utility
  end
  if args.include?("--code")
    programs += $code
  end
  if args.include?("--knowledge")
    programs += $knowledge
  end
  if args.include?("--all")
    programs = $attack + $utility + $code + $knowledge + $mini
    # deduplicate programs
    programs = programs.uniq
  end
  if args.include?("--unattended")
    unattended = true
  else
    unattended = false
  end
  if args.include?("--dry-run")
    dry_run = true
  else
    dry_run = false
  end
  if programs.empty?
    puts "❌ Error: No programs selected. Please select at least one program group to install or uninstall.\nView help❓ with -h or --help for more information."
    exit false
  end
  return uninstall_argument, programs, unattended, dry_run
end

def check_if_homebrew_installed
  command = "which brew"
  output = `#{command}`
  if output.empty?
    puts "❌ Homebrew is not installed. Please install Homebrew and try again. Visit https://brew.sh/ for more information."
    exit false
  else
    puts "✅ Homebrew is installed."
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
      answer = get_yes_no_input "Are you sure you want to uninstall all the above programs? (yes/no) "
    else
      answer = true
    end
    if answer
      programs.each do |program|
        puts "Uninstalling #{program}..."
        uninstall_program(program, progressbar)
        sleep(0.1)
      end
      system("osascript -e 'display notification \"Uninstallation complete\" with title \"Nutek Security Platform\" sound name \"Boop\"' duration 1.5")
    else
      puts "Uninstall aborted!"
    end
  else
    answer = false
    if !unattended
      puts programs
      answer = get_yes_no_input "Are you sure you want to install all the above programs? (yes/no) "
    else
      answer = true
    end
    if answer
      programs.each do |program|
        puts "Installing #{program}..."
        install_program(program, progressbar, dry_run)
        sleep(0.1)
      end
      puts "\nFor more information, see"
      puts "https://nutek.neosb.net/docs/tools/bing-search/ and TOOLS section where you can find more information about each tool."
      system("osascript -e 'display notification \"Installation complete\" with title \"Nutek Security Platform\" sound name \"Boop\"' duration 1.5")
    else
      puts "Install aborted!"
    end
  end
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

####################
# License
####################
# MIT License
#
# This script is part of the Nutek Security Platform (https://nutek.neosb.net/) provided by Nutek Security Solutions and Szymon Błaszczyński.
# You can use this script for free. Please read the license agreement below.
# If you do not agree to this license, you cannot use this script.
#
# MIT License
# Permission is hereby granted, free of charge, to any person obtaining a copy of this script and associated documentation files (the "Script"), to deal in the Script without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Script, and to permit persons to whom the Script is furnished to do so, subject to the following conditions:
# The above license notice and this permission notice shall be included in all copies or substantial portions of the Script.
# THE SCRIPT IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SCRIPT OR THE USE OR OTHER DEALINGS IN THE SCRIPT.
####################
