/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cpluc

import (
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - Std function that builds LUC
func buildLuc(datapip PipelineData) (string, error) {

	logx.L.Debug("building LUC locally")

	// define cli
	cli := fmt.Sprintf(`
	cd /var/tmp/luc 							&& 
	rm -rf /tmp/luc* &> /dev/null &&  
	go build -o %s 					&& 
	sudo mv %s /usr/local/bin/luc && 
	GOOS=linux GOARCH=amd64 go build -o %s && cd -
	`, datapip.localOutput, datapip.localExePath, datapip.localOutXptf)

	// play cli
	_, err := util.RunCLILocal(cli)
	if err != nil {
		return "", err
	}

	logx.L.Debug("builded LUC locally and deployed it locally")
	return "", nil
}
