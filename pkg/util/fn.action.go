/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/jedib0t/go-pretty/table"
)

// type FnActionHandler func(...string) (string, error)
type FnActionHandler struct {
	Fn       func(...string) (string, error) // a function with that signature
	NbParams int                             // the number of parameters required by that function
}

var FnActionMap = map[string]FnActionHandler{
	"AddLineToFile":    {Fn: AddLineToFileFn, NbParams: 2},
	"IsFileExists":     {Fn: IsFileExistsFn, NbParams: 1},
	"MoveFile":         {Fn: MoveFileFn, NbParams: 4},
	"SaveStringToFile": {Fn: SaveStringToFileFn, NbParams: 3},
}

func IsFileExistsFn(fnParameters ...string) (string, error) {
	// check arg
	if _, err := logAndCheckParams(1, fnParameters...); err != nil {
		return "", err
	}

	// get input
	srcFilePath := fnParameters[0]

	return IsFileExists(srcFilePath)
}

func MoveFileFn(fnParameters ...string) (string, error) {
	// check arg
	if _, err := logAndCheckParams(4, fnParameters...); err != nil {
		return "", err
	}

	// get input
	srcFilePath := fnParameters[0]
	dstFilePath := fnParameters[1]
	filePermissionStr := fnParameters[2]
	isRootFile := strings.ToLower(fnParameters[3]) == "true"

	// convert input
	permUint64, err := strconv.ParseUint(filePermissionStr, 8, 32)
	if err != nil {
		return "", fmt.Errorf("❌ Error: invalid permission format: %w", err)
	}
	filePermission := os.FileMode(permUint64)

	return MvFile(srcFilePath, dstFilePath, filePermission, isRootFile)

}

func SaveStringToFileFn(fnParameters ...string) (string, error) {
	// check arg
	if _, err := logAndCheckParams(3, fnParameters...); err != nil {
		return "", err
	}

	// get input
	data := fnParameters[0]
	path := fnParameters[1]
	isRootPath := strings.ToLower(fnParameters[2]) == "true"

	return SaveStringToFile(data, path, isRootPath)
}

func AddLineToFileFn(fnParameters ...string) (string, error) {
	// check arg
	if _, err := logAndCheckParams(2, fnParameters...); err != nil {
		return "", err
	}

	// get input
	filepath := fnParameters[0]
	line := fnParameters[1]

	return AddLineToFile(filepath, line)
}

func PlayFnLocally(fnKey string, fnParameters ...string) (string, error) {

	// get the instance
	fnActionHandler, ok := FnActionMap[fnKey]
	if !ok {
		return "", fmt.Errorf("❌ unknown function key requested: %s", fnKey)
	}

	// log and check parameters for that function
	if _, err := logAndCheckParams(fnActionHandler.NbParams, fnParameters...); err != nil {
		return "", err
	}

	// execute the function locally
	return fnActionHandler.Fn(fnParameters...)
}
func PlayFnOnRemote(vm string, fnKey string, fnParameters ...string) (string, error) {

	// get the instance
	fnActionHandler, ok := FnActionMap[fnKey]
	if !ok {
		return "", fmt.Errorf("❌ unknown function key requested: %s", fnKey)
	}

	// log and check parameters for that function
	if _, err := logAndCheckParams(fnActionHandler.NbParams, fnParameters...); err != nil {
		return "", err
	}

	// create sequence of quoted paramaters
	listParams := ""
	for _, param := range fnParameters {
		// skip empty parameters - due to the way parameters are passed via "luc action ...""
		if strings.TrimSpace(param) == "" {
			continue
		}
		listParams = fmt.Sprintf("%s %q", listParams, param)
	}

	// define the cli
	cli := fmt.Sprintf(`luc action %s %s`, fnKey, listParams)

	// // log
	logx.L.Debugf("[%s] Running on remote CLI: %s", vm, cli)

	// execute the function remotely
	return RunCLIRemote(vm, cli)
}

func logAndCheckParams(nbRequired int, fnParameters ...string) (int, error) {
	nbParams := 0
	for _, p := range fnParameters {
		// skip empty parameters
		if strings.TrimSpace(p) != "" {
			// logx.L.Debugf("param: %s", p)
			nbParams++
		}
	}
	if nbParams < nbRequired {
		return nbParams, fmt.Errorf("❌ Error: %d parameters are required", nbRequired)
	}
	return nbParams, nil
}

// ShowFnActionMap displays the list of available remote functions.
func ShowFnActionMap() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Header
	t.AppendHeader(table.Row{"Available Functions"})

	// Collect and sort keys
	var functionKeys []string
	for name := range FnActionMap {
		functionKeys = append(functionKeys, name)
	}
	sort.Strings(functionKeys)

	// Add each action as a row
	for _, name := range functionKeys {
		t.AppendRow(table.Row{name})
	}

	t.Render()
}
