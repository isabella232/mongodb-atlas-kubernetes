#!/bin/sh

git config --local user.email "41898282+github-actions[bot]@users.noreply.github.com"
git config --local user.name "github-actions[bot]"

git pull origin pull
date > 1.txt
git add 1.txt
git status
git commit -m "Update configuration"
git push origin test/sign
# git tag "v${INPUT_VERSION}"
# git push origin --tags
