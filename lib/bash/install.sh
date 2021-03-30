#!/bin/bash

#install brew
export NONINTERACTIVE=1
bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

#install casks
brew install --cask google-chrome
brew install --cask slack
brew install --cask visual-studio-code
brew install --cask 1password
brew install --cask docker
#install dev-tools
brew install azure-cli helm istioctl kubectl zsh jq go kubectx git go node terraform


bash <(curl -s https://raw.githubusercontent.com/pritunl/pritunl-client-electron/master/tools/install_macos.sh)


