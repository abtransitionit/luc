/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/internal/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

const CheckSshDescription = "check all nodes of the KBE clusters are SSH reachable."

func checkSsh(arg ...string) (string, error) {
	logx.L.Info(CheckSshDescription)

	// convert the list of nodes to a go slice
	SliceNodes := strings.Fields(config.KbeListNode)

	// check nodes are SSH configured
	logx.L.Info("check nodes are SSH configured")
	for _, vm := range SliceNodes {
		ok, err := util.IsVmSshConfigured(vm)
		if err != nil {
			return "", fmt.Errorf("vm: %s: %v", vm, err)
		}
		fmt.Printf("Node   %-5s ssh configured: %v\n", vm, ok)
	}

	// check nodes are SSH reachable
	logx.L.Info("check nodes are SSH reachable")
	for _, vm := range strings.Fields(config.KbeListNode) {
		ok, err := util.IsSshConfiguredVmSshReachable(vm)
		if err != nil {
			fmt.Printf("Error for %s: %v\n", vm, err)
			continue
		}
		fmt.Printf("Node   %-5s ssh reachable: %v\n", vm, ok)
	}
	// check results

	// on SUCCESS
	return "", nil
}

// for 1 VM
type CheckResult struct {
	Node       string
	Configured bool
	Reachable  bool
}

// for a set of VMs
var CurrentClusterResult = map[string]CheckResult{}
