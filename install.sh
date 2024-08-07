#!/usr/bin/env bash

install_path=$NUTEK_APPLE_ROOT


if [[ $(which nutek-apple ) ]]; then
    echo "OK. nutek-apple already installed"
    return
fi

# install Homebrew
if ! [[ $(which brew) ]] && [[ "$OSTYPE" == "darwin"* ]]; then
    echo "Installing Homebrew for macOS (and Linux)"
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
else
    echo "OK. Homebrew already installed"
fi



# Clone repository and build Go program
if [ -n "$install_path" ]; then
    git clone https://github.com/NutekSecurity/nutek-apple.git $install_path
    cd $install_path
else
    git clone https://github.com/NutekSecurity/nutek-apple.git $HOME/.nutek-apple
    cd $HOME/.nutek-apple
fi

go build
sudo ln -s nutek-apple $HOMEBREW_PREFIX/bin/nutek-apple # where to symlink on macOS silicon?

if [[ "$OSTYPE" == "darwin"* ]]; then # Homebrew on macOS
    if [ -n "$install_path" ]; then
        echo "export NUTEK_APPLE_ROOT=$NUTEK_APPLE_ROOT" >> $HOME/.zprofile
    else
        echo "export NUTEK_APPLE_ROOT=$HOME/.nutek-apple" >> $HOME/.zprofile
    fi
fi

source $HOME/.zprofile

nutek-apple