#!/bin/bash

if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "macOS detected, installing Homebrew..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    echo "Homebrew installation completed."
else
    echo "Not macOS, skipping Homebrew installation."
fi