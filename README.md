# nutek-apple 🍎

Nutek Security Platform for macOS and Linux operating systems. Tools for hackers, bug hunters and hobbiests. You might like it, you might not. It's a matter of taste.

## What it does?

This - `nutek-apple.rb` - script for macOS and Linux operating systems installs
all the neccessairly tools for well prepared hacker, bug hunter or
a seasoned hobbiest willing to pursue a way of turning things inside out and
is not scared to use only command line tools with . Why? Almost all of the fancy
GUI (don't there are some of them) apps started as a command line tool, or are
in many ways similar, or even less capable, because in command line - you're
on top, you're left alone like Alice in Wonderland deep inside your dreams of power, and you know what? You can have this power! Just follow along.


The only exceptions for this are text editors, terminals and Wireshark (try termshark in CLI).

## Installation

1. Install Homebrew

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

2. Clone this repository

```bash
git clone https://github.com/NutekSecurity/nutek-apple.git
```

to some safe location on your Apple/PC computer. Make sure you have Ruby installed, and also Homebrew (that's from where we'll be fetching apps) - you can get Homebrew from [brew.sh](https://brew.sh).

3. You should already have Ruby installed on your computer, but if you don't, you can install it with Homebrew:

```bash
brew install ruby
```

on Linux, you can install Ruby with your package manager, for example on Ubuntu:

```bash
sudo apt install ruby
```
on Fedora:

```bash
sudo dnf install ruby
```

4. You're on start line. Run this command:

```shell
bundle install --without development,test
```

```shell
ruby nutek-apple.rb
```

or make this happen:

```shell
chmod 744 nutek-apple.rb
./nutek-apple.rb
```

This should present you with help message, and there you will find everything you need, but in case of clarity, I provide the transcript.

```text
Usage: ruby nutek-apple.rb [options]
Automated installation of hacking command line programs on macOS - Nutek Security Platform. Requires Homebrew.
Curated by Nutek Security
Downloading the latest version from GitHub:
https://github.com/NutekSecurity/nutek-apple

Options:
  -h, --help				Show this help message and exit
  -i, --install				Install programs. Choose programs to install with --gui, --cli or --all
  -u, --uninstall			Uninstall programs. Choose programs to uninstall with --gui, --cli or --all
  --license				Show license information
  --all					Install or uninstall all programs
  --cli				Install or uninstall cli set of programs
  --gui				Install or uninstall GUI programs
  --list				List all programs
  --list-cli				List cli set of programs
  --list-gui				List gui programs
  --unattended				Unattended mode. Install selected programs without asking for confirmation (on Linux run with sudo)
  --dry-run				Dry run. Show what would be installed without actually installing anything

Examples:
  ruby nutek-apple.rb --install --cli
  ruby nutek-apple.rb --install --gui
  ruby nutek-apple.rb --install --all
  ruby nutek-apple.rb --uninstall --gui
  ruby nutek-apple.rb --uninstall --all
  ruby nutek-apple.rb --list
  ruby nutek-apple.rb --list-gui
  ruby nutek-apple.rb --unattended --install --gui
  ruby nutek-apple.rb --unattended --install --all
  ruby nutek-apple.rb --unattended --uninstall --gui
  ruby nutek-apple.rb --unattended --uninstall --all
```

## What you will get?

```text
ollama - run LLM locally
podman - docker alternative
neovim - text editor
openvpn - vpn client
irssi - irc client
dos2unix - convert dos to unix
ipcalc - ip calculator
whatmask - ip calculator
expect - automate interactive applications
fd - find alternative
tmux - terminal multiplexer
lsd - ls alternative
bat - cat alternative
ripgrep-all - grep alternative
sd - sed alternative
termshark - wireshark alternative
httpie - curl alternative
smap - map network
nmap - network scanner
p0f - passive os fingerprinting
masscan - port scanner
feroxbuster - directory bruteforcer
ffuf - directory bruteforcer
nuclei - vulnerability scanner
mitmproxy - transparent proxy
metasploit - red team toolbox
httpx - http scanner
amass - subdomain scanner
jq - json parser
htmlq - html parser
httrack - website copier
monolith - website copier
mdcat - markdown reader
ouch - archive extractor
exploitdb - exploit database
asciinema - terminal recorder
agg - ascii art generator
hashcat - password cracker
john-jumbo - password cracker
mdbook - book generator
```

and

```text
podman-desktop - docker alternative
imhex - hex editor
warp - terminal
alacritty - terminal
kitty - terminal
wireshark - network analyzer
font-hack-nerd-font - font
zed - text editor
vscodium - text editor
```
