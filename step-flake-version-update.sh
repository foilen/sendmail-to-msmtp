#!/bin/bash

set -e

RUN_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $RUN_PATH

echo ----[ Update flake.nix version ]----
sed -i "s/version = \".*\";/version = \"$VERSION\";/" flake.nix

if ! git diff --quiet -- flake.nix; then
  git add flake.nix
  git commit -m "Release $VERSION"
fi
