/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"os"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

const DisplayDescription = "display KBE Cluster's informations."

func display(arg ...string) (string, error) {
	logx.L.Info(DisplayDescription)
	// display informayions on the cluster
	DisplayClusterConfig(CurrentClusterConfig)
	// on SUCCESS
	return "", nil
}

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

// display the cluster node informations from the structure ClusterConfig
func DisplayClusterConfig(config ClusterConfig) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.Style().Title.Align = text.AlignCenter

	t.AppendHeader(table.Row{"Key", "Value"})
	t.AppendRows([]table.Row{
		{"Node names", config.KbeListNode},
		{"Worker Node names", config.KbeListNodeWorker},
		{"Control Plane Node names", config.KbeListNodeCplane},
	})

	t.Render()
}
