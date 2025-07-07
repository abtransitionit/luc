/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/jedib0t/go-pretty/table"
)

type ActionHandler func(...string) (string, error)

var ActionMap = map[string]ActionHandler{
	"AddLineToFile":    AddLineToFileAction,
	"SaveStringToFile": SaveStringToFileAction,
}

func SaveStringToFileAction(params ...string) (string, error) {
	if len(params) < 3 {
		return "", fmt.Errorf("need 3 parameters: data, path, pathIsRoot")
	}

	// get input
	data := params[0]
	path := params[1]
	pathIsRoot := strings.ToLower(params[2]) == "true"

	return SaveStringToFile(data, path, pathIsRoot)
}

func AddLineToFileAction(params ...string) (string, error) {

	// count nb of non empty params
	nbParams := 0
	for _, p := range params {
		if p != "" {
			logx.L.Debugf("param: %s", p)
			nbParams++
		}
	}

	// check nb of params
	if len(params) < 2 {
		return "", fmt.Errorf("need 2 parameters: filePath and line")
	}

	// get input
	filepath := params[0]
	line := params[1]

	return AddLineToFile(filepath, line)
}

func PlayActionLocal(action string, params ...string) (string, error) {
	fn, ok := ActionMap[action]
	if !ok {
		return "", fmt.Errorf("❌ unknown action requested: %s", action)
	}
	return fn(params...)
}
func PlayActionRemote(vm string, action string, params ...string) (string, error) {
	logx.L.Debugf("Remote action : %s", action)

	// create seq of quoted params
	listParams := ""
	for _, param := range params {
		// skip empty params
		if strings.TrimSpace(param) == "" {
			continue
		}
		listParams = fmt.Sprintf("%s %q", listParams, param)
	}
	cli := fmt.Sprintf(`luc action %s %s`, action, listParams)
	logx.L.Debugf("Running remotely CLI: %s", cli)

	return RunCLIRemote(vm, cli)
	// return "", nil
}

// func GetActionRemote(vm string, action string, params ...string) (string, error) {
// 	cli := fmt.Sprintf(`luc remote %s`, action)
// 	if len(params) > 0 {
// 		for i, param := range params {
// 			params[i] = param
// 		}
// 		actionParams := strings.Join(params, " ")
// 		// actionParams:=params
// 		cli = fmt.Sprintf(`luc remote %s '%s'`, action, actionParams)
// 	}
// 	logx.L.Debugf("Running remotely CLI: %s", cli)
// 	// return RunCLIRemote(vm, cli)
// 	return "", nil
// }

// ShowActionMap displays the list of available remote functions.
func ShowActionMap() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Header
	t.AppendHeader(table.Row{"Available Remote Actions"})

	// Collect and sort keys
	var actionNames []string
	for name := range ActionMap {
		actionNames = append(actionNames, name)
	}
	sort.Strings(actionNames)

	// Add each action as a row
	for _, name := range actionNames {
		t.AppendRow(table.Row{name})
	}

	t.Render()
}
