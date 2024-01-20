#!/usr/bin/env sh

go build -ldflags "-X eslint-git-diff/cmd.stringVersion=$(git describe --abbrev=0 --tags)" -v -o ~/bin/eslint-git-diff

echo 'Save executable in ~/bin/eslint-git-diff'
