#!/bin/bash
## Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com

set -e # make the script stop at the first failing test:

# define var
TestFolderName="test"
TestFolderPath="$(dirname $(go env GOMOD))/${TestFolderName}" # Test folder for this GO project


# get normalized input
Input01=$(echo "$1" | tr '[:upper:]' '[:lower:]')

# choose folder based on input
case "$Input01" in
    rem|remote|r) 
        SubTestFolder="remote" ;;
    l|local|locale|loc) 
        SubTestFolder="local" ;;
    *) 
        echo "❌ Error: Invalid input '$1'. Use remote/rem/r or local/l/loc/locale" >&2
        exit 1 ;;
esac

# choose tests based on input
[ -n "$2" ] && TestToPlay="-run Test${2}" || TestToPlay=""


# go to test folder
cd $TestFolderPath

# Run all [function]test in folder 
# echo "playing test for [function]test in folder $SubTestFolder"
echo "✅ playing cli > go test -v ./${SubTestFolder} ${TestToPlay} "
go test -v ./${SubTestFolder} ${TestToPlay}
# go test -v -run TestRemote
