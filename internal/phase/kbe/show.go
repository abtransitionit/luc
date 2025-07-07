/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"fmt"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/jedib0t/go-pretty/table"
)

// Structure used to holds the cluster node informations.
type ClusterConfig struct {
	KbeListNode       string
	KbeListNodeWorker string
	KbeListNodeCplane string
}

// 1 instance of the structure for the cluster
var CurrentClusterConfig = ClusterConfig{
	KbeListNode:       config.KbeListNode,
	KbeListNodeWorker: config.KbeListNodeWorker,
	KbeListNodeCplane: config.KbeListNodeWorker,
}

const DisplayDescription = "display the desired KBE Cluster's configuration."

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
		{"All Node names", obj.KbeListNode},
		{"Worker Node names", obj.KbeListNodeWorker},
		{"Control Plane Node names", obj.KbeListNodeCplane},
	})

	return t.Render() // returns string output of the table
}
