#!/bin/sh
# scripts/completions.sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
    pwd
	go run ../main.go completion "$sh" >"completions/azioncli.$sh"
done
source completion/*