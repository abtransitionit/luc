/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"fmt"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/jedib0t/go-pretty/table"
)

const DisplayDescription = "display the desired KIND Cluster's configuration."

// Structure used to holds the informations.
type ClusterConfig struct {
	VmName string
}

// 1 instance of the structure
var CurrentClusterConfig = ClusterConfig{
	VmName: config.KindVm,
}

func show(arg ...string) (string, error) {
	logx.L.Info(DisplayDescription)

	fmt.Println(CurrentClusterConfig)
	return "", nil
}

// pretty display
func (obj ClusterConfig) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetTitle("Cluster config")
	t.AppendHeader(table.Row{"Type", "Node(s)"})
	t.AppendRows([]table.Row{
		{"VM name", obj.VmName},
	})

	return t.Render() // returns string output of the table
}
