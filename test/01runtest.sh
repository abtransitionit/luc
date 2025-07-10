#!/bin/bash
## Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com

set -e # make the script stop at the first failing test:

# define var
TestFolderPath="$(dirname $(go env GOMOD))/test" # Test folder for this GO project

# go to test folder
cd $TestFolderPath

# Run all [function]test starting with "TestLocal" in folder "local"
go test -v ./local -run TestLocal
# go test -v -run TestRemote
