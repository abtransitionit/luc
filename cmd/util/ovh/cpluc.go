/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

import (
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/phase"
	"github.com/spf13/cobra"
)

// Description
var cplucSDesc = "copy LUC CLI from local to remote VM(s)."
var cplucLDesc = cplucSDesc + ` xxx.`

// init Command
var CplucCmd = &cobra.Command{
	Use:   "cpluc",
	Short: cplucSDesc,
	Long:  cplucLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.L.Infof("%s", cplucSDesc)
		cmd.Help()
	},
}

func init() {
	phase.CmdInit(CplucCmd)
}
