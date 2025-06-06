/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package packagex

import (
	"github.com/spf13/cobra"
)

// Description
var packageSDesc = "Manage packages."
var packageLDesc = packageSDesc + ` xxx.`

// root Command
var PackageCmd = &cobra.Command{
	Use:   "package",
	Short: packageSDesc,
	Long:  packageLDesc,
}

func init() {
	PackageCmd.AddCommand(addCmd)
	PackageCmd.AddCommand(updateCmd)
}
