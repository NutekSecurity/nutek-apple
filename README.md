# nutek-apple

Nutek Security Platform for macOS operating systems.

## What it does?

This - `nutek-apple.rb` - script for macOS operating systems installs
all the neccessairly tools for well prepared hacker, bug hunter or
a seasoned hobbiest willing to pursue a way of turning things inside out and
is not scared to use only command line tools. Why? Almost all of the fancy
GUI (graphical user interface) apps started as a command line tool, or are
in many ways similar, or even less capable, because in command line - you're
on top, you're left alone like Alice in Wonderland deep inside your dreams of power, and you know what? You can have this power! Just follow along.

### List of tools

* Attack
  - bettercap
  - ettercap
  - hashcat
  - hydra
  - john-jumbo
  - metasploit
  - mitmproxy
  - ncrack
  - socat
  - sqlmap
  - xh
* Code
  - adns
  - bvi
  - capstone
  - cflow
  - cppcheck
  - fmt
  - gitleaks
  - grpc
  - hyperscan
  - jansson
  - radare2
  - rats
  - zzuf
* Knowledge
  - amass
  - arp-scan
  - bmon
  - dnsx
  - feroxbuster
  - ffuf
  - fping
  - gau
  - gobuster
  - httpx
  - httrack
  - masscan
  - mtr
  - nethogs
  - ngrep
  - nikto
  - nmap
  - nuclei
  - p0f
  - rustscan
  - smap
  - tcpdump
  - tcpflow
  - termshark
  - testssl
  - zmap
* Utility
  - argon2
  - bat
  - bottom
  - discount
  - dos2unix
  - dust
  - exa
  - expect
  - hexyl
  - htmlq
  - ipcalc
  - irssi
  - jq
  - jql
  - macchina
  - mdcat
  - monolith
  - neovim
  - openvpn
  - ouch
  - pandoc
  - podman
  - podman-compose
  - ripgrep-all
  - sd
  - tmux
  - viu
  - w3m
  - whatmask

## Installation

Either clone, or download this repository to some safe location on your Apple computer. Make sure you have Ruby installed, and also Homebrew (that's from where we'll be fetching apps).

Run this command:

```shell
ruby nutek-apple.rb
```

or make this happen:

```shell
chmod 744 nutek-apple.rb
./nutek-apple.rb
```

This should present you with help message, and there you will find everything you need, but in case of clarity, I provide the transcript.

```
Usage: ruby nutek-apple.rb [options]
Automated installation of hacking command line programs on macOS - Nutek Security Platform. Requires Homebrew.
Curated by Nutek Security Solutions
	and Szymon Błaszczyński.

Options:
  -h, --help				Show this help message and exit
  -i, --install				Install programs. Choose programs to install with --attack, --code, --knowledge, --utility or --all
  -u, --uninstall			Uninstall programs. Choose programs to uninstall with --attack, --code, --knowledge, --utility or --all
  --all					Install or uninstall all programs
  --attack				Install or uninstall attack programs
  --code				Install or uninstall code programs
  --knowledge				Install or uninstall knowledge programs
  --utility				Install or uninstall utility programs
  --list				List all programs
  --list-attack				List attack programs
  --list-code				List code programs
  --list-knowledge			List knowledge programs
  --list-utility			List utility programs
  --unattended				Unattended mode. Install selected programs without asking for confirmation
  --dry-run				Dry run. Show what would be installed without actually installing anything

Examples:
  ruby nutek-apple.rb --install --attack --utility --knowledge
  ruby nutek-apple.rb --install --all
  ruby nutek-apple.rb --uninstall --attack --code
  ruby nutek-apple.rb --uninstall --all
  ruby nutek-apple.rb --list
  ruby nutek-apple.rb --list-attack
  ruby nutek-apple.rb --unattended --install --attack
  ruby nutek-apple.rb --unattended --install --all
  ruby nutek-apple.rb --unattended --uninstall --attack
  ruby nutek-apple.rb --unattended --uninstall --all

For more information, see
https://nutek.neosb.net/docs/tools/bing-search/ and TOOLS section where you can find more information about each tool.
```

## Why?

**Nutek Security Platform** is one stop solution for your reconessaince, information gathering and establishing attack surface on targets of your choice. Have fun exploring vast library of hacking challenges, games and hackathons being like a boss. Tired of non-stop clicking mouse instead of real input to the prompt? Then this is for you. You're in control. You choose what you'll do with this tools and you are responsible for the outcome.