#!/bin/bash
## Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com

set -e # make the script stop at the first failing test:

# define var
TestFolderPath="$(dirname $(go env GOMOD))/test" # Test folder for this GO project


# get normalized input
UserInput=$(echo "$1" | tr '[:upper:]' '[:lower:]')

# Determine mode
case "$UserInput" in
    rem|remote|r) 
        mode="remote" ;;
    l|local|locale|loc) 
        mode="local" ;;
    *) 
        echo "❌ Error: Invalid input '$1'. Use remote/rem/r or local/l/loc/locale" >&2
        exit 1 ;;
esac

# define var from input
SubTestFolder="$mode"

# go to test folder
cd $TestFolderPath

# Run all [function]test in folder 
echo "playing test for [function]test in folder $SubTestFolder"
go test -v ./${SubTestFolder}
# go test -v -run TestRemote
