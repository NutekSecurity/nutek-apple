![logo](logo.png)

# nutek-apple 🍎

Nutek Security Platform for macOS and Linux operating systems. Tools for hackers, bug hunters and hobbiests. You might like it, you might not. It's a matter of taste.

## What it do?

This - part of `Nutek Security Platform` - program for _macOS_ and _Linux_ operating systems installs all the neccessairly _tools_ for well prepared 
_hacker_, _bug hunter_ or
a seasoned _hobbiest_ willing to pursue a way of turning things inside out and
is not scared to use only command line tools with . Why? Almost all of the fancy
GUI apps started as a command line tool, or are
in many ways similar, or even less capable, because in command line - you're
on the top, you're left alone like _Alice in Wonderland_ deep inside your dreams of power, and you know what? You can have this power! Just follow along.

If you look for some other help in hacking, pentesting or else, there are
links to other parts of Internet, that when you instal Firefox, or have
been using Safari, will open in browser window...


The only exceptions for the CLI rule are text editors (NeoVim in CLI), terminals and Wireshark (try termshark in CLI).

## Installation

### Homebrew

1. Install Homebrew (dependency for most mmacOS packages and some Linux too)

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

or look it up on [brew.sh](https://brew.sh)

2. Install Go programming language

You may use your Homebrew

```bash
brew install go
```

or directly from upstream on the official Go language webpage

[https://go.dev/dl/](https://go.dev/dl/)

### nutek-apple installation using go or latest binaries

Use Go or download [latest builds](https://github.com/nuteksecurity/nutek-apple/releases/)

```bash
go install -v github.com/nuteksecurity/nutek-apple@latest
```

### nutek-apple first run

4. You're on the start line. Run this command:

```shell
bundle install
```

```shell
ruby nutek-apple.rb
```

or make this happen:

```shell
chmod 744 nutek-apple.rb
./nutek-apple.rb
```

## If you are stuck, and can't run the script

call the git pull command:

```bash
git pull origin main --rebase
```

This will let you incorporate any changes you've made.

## What you will get?

6GB worth of loot, including:

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
bettercap - network monitoring
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
firefox - web browser
zap - web app scanner
```

## TODO

* change README.md to reflect change to Go programming languagae
* ⚠️ add useful links with open by firefox option
* ⚠️ export git auto-update to Ruby gem
* ⚠️ add install and uninstall for one program
* ⚠️ make user able to read Homebrew information about a particular program; thin wrapper-like around Homebrew - not to deep, just basic install, uninstall and info, also very important, __update__
* ⚠️ keep user in my app, so no info is leaked to Homebrew and back.
* ⚠️ add programming languages (rust, ruby - (how? maybe with an installation script?), python, dotnet (C#?, powershell?), go...) to list of available programs
* ⚠️ add nutek-cipher and nutek-smoke
* ⚠️ make a proposal for katarina to gather around Zed text editor and collaborate
* ⚠️ write about each program in my book - _Hacking with a DREAM in the mind_
* ⚠️ add test that check if software is installed, and other cli commands
