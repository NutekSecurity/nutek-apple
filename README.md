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

## Installation

Clone this repository `git clone https://github.com/nutek-terminal/nutek-apple.git` to some safe location on your Apple computer. Make sure you have Ruby installed (preinstalled with every on Mac), and also Homebrew (that's from where we'll be fetching apps) - you can get Homebrew from [brew.sh](https://brew.sh).

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

```text
Usage: ruby nutek-apple.rb [options]
Automated installation of hacking command line programs on macOS - Nutek Security Platform. Requires Homebrew.
Curated by Nutek Security Solutions
	and Szymon Błaszczyński.
Download the latest version from GitHub:
https://github.com/NutekSecurity/nutek-apple
This version: 0.1.7 GitHub version: 0.1.7

Options:
  -h, --help				Show this help message and exit
  -i, --install				Install programs. Choose programs to install with --gui, --code, --knowledge, --utility or --all
  -u, --uninstall			Uninstall programs. Choose programs to uninstall with --gui, --code, --knowledge, --utility or --all
  --web, --safari			Open Safari and lookup nuteksecurity.com
  --license				Show license information
  --all					Install or uninstall all programs
  --cli				Install or uninstall cli set of programs
  --gui				Install or uninstall GUI programs
  --list				List all programs
  --list-cli				List cli set of programs
  --list-gui				List gui programs
  --unattended				Unattended mode. Install selected programs without asking for confirmation
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