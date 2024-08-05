#!/usr/bin/env bash

if [[ $(which nutek-apple ) ]]; then
    echo "OK. nutek-apple already installed"
    exit 0
fi

# install Homebrew
if ! [[ $(which brew) ]] ; then
    echo "Installing Homebrew for macOS (and Linux)"
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
else
    echo "OK. Homebrew already installed"
fi

# Check operating system and install Go accordingly
if [[ "$OSTYPE" == "darwin"* ]]; then # Homebrew on macOS
    echo "OK. curl should already be installed"
elif ! which dnf &>/dev/null; then
    if ! [[ $(which curl) ]]; then
        echo "Installing curl"
        sudo dnf install -y curl
        echo "OK. curl is installed"
    else
        echo "OK. curl is already installed"
    fi
else
    if ! [[ $(which curl) ]]; then
        echo "Installing curl"
        sudo apt update -y
        sudo apt install -y curl
    else
        echo "OK. curl is already installed"
    fi 
fi

"aarch64"

# Clone dummy repository and build Go program
mkdir ~/.nutek-apple
cd ~/.nutek-apple
if [[ $(which curl) ]]; then
    if [[ "$OSTYPE" == "darwin"* ]] && [[ "$(uname -m)" == "arm644"* ]]; then 
        echo "Downloading nutek-apple for macOS arm64"
        curl -fsSL https://github.com/nutek-apple/release/nutek-apple-apple-arm64
        sha256sum $(curl -fsSL sh256)
    elif [[ "$OSTYPE" == "darwin"* ]] && [[ "$(uname -m)" == "x86_64"* ]]; then 
        echo "Downloading nutek-apple for macOS x86_64"
        curl -fsSL https://github.com/nutek-apple/release/nutek-apple-apple-x86_64
        sha256sum $(curl -fsSL sh256)
    elif [[ "$(uname -s)" == "Linux"* ]]; [[ "$(uname -m)" == "aarch64"* ]]; then 
        echo "Downloading nutek-apple for Linux aarch64"
        curl -fsSL blahablasd-aarch64
        sha256sum $(curl -fsSL sh256)
    elif [[ "$(uname -s)" == "Linux"* ]]; [[ "$(uname -m)" == "arm644"* ]]; then
        echo "Downloading nutek-apple for Linux x86_64"
        curl -fsSL blahablasd-x86_64
        sha256sum $(curl -fsSL sh256)
    fi
else
    echo "Error: You don't seem to have curl installed"
    exit 1
fi

# Add executable to PATH and load
if [[ $(echo "$SHELL" | grep -o "zsh") ]]; then
    echo 'export PATH=$PATH:$HOME/.nutek-apple' >> ~/.zshrc
    source ~/.zshrc
elif [[ $(echo "$SHELL" | grep -o "bash") ]]; then
    echo 'export PATH=$PATH:$HOME/.nutek-apple' >> ~/.bashrc
    source ~/.bashrc
fi
