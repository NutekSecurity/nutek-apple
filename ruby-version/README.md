![logo](logo.png)

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

### Homebrew

1. Install Homebrew

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

or look it up on [brew.sh](https://brew.sh)

### nutek-apple repository

2. Clone this repository

```bash
git clone https://github.com/NutekSecurity/nutek-apple.git
```

to some safe location on your Apple/PC computer. Make sure you have Ruby installed, and also Homebrew (that's from where we'll be fetching apps) - you can get Homebrew from [brew.sh](https://brew.sh).

### Ruby programming lanuguage

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

#### Ruby version manager rbenv

* 3.1. You can also use ruby version manager like `rbenv` _(I use it)_ to seamlessly change between versions of ruby with
`rbenv` command, or `.ruby-version` file

macOS, Linux (with Homebrew installed):

```shell
brew install rbenv
```

Debian based Linux:

```shell
sudo apt install rbenv
```

Red Hat based Linux:

```shell
sudo dnef install rbenv
```

and then use something along this lines, one by one

```shell
rbenv init
rbenv install --list
rbenv install 3.3.4
rbenv rehash
rbenv local 3.3.4
```

It first initialize rbenv, shows ruby versions to intall, install the long term support one, initialize it and set local
for this app.

After installing Ruby from rbenv, _restart your terminal_ and proceed to my nutek-apple specifics.

To check what ruby version is used:

```shell
ruby --version
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

* ⚠️ export git auto-update to Ruby gem
* ⚠️ add install and uninstall for one program
* ⚠️ make user able to read Homebrew information about a particular program; thin wrapper-like around Homebrew - not to deep, just basic install, uninstall and info, also very important, __update__
* ⚠️ keep user in my app, so no info is leaked to Homebrew and back.
* ⚠️ add programming languages (rust, ruby - (how? maybe with an installation script?), python, dotnet (C#?, powershell?), go...) to list of available programs
* ⚠️ add nutek-cipher and nutek-smoke
* ⚠️ make a proposal for katarina to gather around Zed text editor and collaborate
* ⚠️ write about each program in my book - _Hacking with a DREAM in the mind_
* ⚠️ add test that check if software is installed, and other cli commands
