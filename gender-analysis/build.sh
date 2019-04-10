#!/usr/bin/env bash

NODE_VERSION="8.10"

export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh"  # This loads nvm

nvm use ${NODE_VERSION}

pushd gender-analysis

npm install
npm run riffraff-artefact