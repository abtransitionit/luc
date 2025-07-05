/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// # Purpose
//
// - Build and deploy luc for the current platform
//
// # Parameters
//
// - srcProjectFolderPath: the absolute folder path to the local GIT project
// - dstBinaryFolderPath:  the absolute folder path to the output file
const LucBadDescription = "Build Locally (for the current platform) the Go CLI:LUC from a local GIT project folder. Then, deploy it locally."

func LucBad(srcProjectFolderPath string, dstBinaryFolderPath string, dstBinaryFileName string) (dstFilePath string, err error) {
	logx.L.Debug(LucBadDescription)

	// check argument
	if srcProjectFolderPath == "" {
		return "", fmt.Errorf("❌ Error: folderPath not provided for srcProjectFolderPath")
	}
	if dstBinaryFolderPath == "" {
		return "", fmt.Errorf("❌ Error: folderPath not provided for dstBinaryFolderPath")
	}
	if dstBinaryFileName == "" {
		return "", fmt.Errorf("❌ Error: folderPath not provided for dstBinaryFileName")
	}

	// check git local project path is absolute
	_, err = filepath.Abs(srcProjectFolderPath)
	if err != nil {
		return "", fmt.Errorf("path is not an absolute path: %v", err)
	}

	// check dst binary file folder is absolute
	_, err = filepath.Abs(dstBinaryFolderPath)
	if err != nil {
		return "", fmt.Errorf("path is not an absolute path: %v", err)
	}

	// check folder exists
	if _, err := os.Stat(srcProjectFolderPath); os.IsNotExist(err) {
		return "", fmt.Errorf("path does not exist: %v", err)
	}

	// check folder exists
	if _, err := os.Stat(dstBinaryFolderPath); os.IsNotExist(err) {
		return "", fmt.Errorf("path does not exist: %v", err)
	}

	// TODO:IMPROVE:MAKE IT GENERIC: remove old stuff
	cli := `rm -rf /tmp/luc* &> /dev/null`
	_, err = util.RunCLILocal(cli)
	// error
	if err != nil {
		return "", err
	}

	// define var
	dstBinaryFilePath := filepath.Join(dstBinaryFolderPath, dstBinaryFileName)

	// build the GO CLI:LUC
	cli = fmt.Sprintf(`
		cd %s && 
		go build -o %s && 
		cd -
		`,
		srcProjectFolderPath, // eg. /var/tmp/luc
		dstBinaryFilePath,
	)
	_, err = util.RunCLILocal(cli)

	// error
	if err != nil {
		return "", err
	}

	// success
	return "", nil
}

// // // define cli
// cli := fmt.Sprintf(`
// cd /var/tmp/luc 							&&
// rm -rf /tmp/luc* &> /dev/null &&
// go build -o %s 					&&
// sudo mv %s /usr/local/bin/luc &&
// GOOS=linux GOARCH=amd64 go build -o %s && cd -
// `, datapip.localOutput, datapip.localExePath, datapip.localOutXptf)

// // play cli
// _, err := util.RunCLILocal(cli)
// if err != nil {
// 	return "", err
// }

// 	sudo mv %s /usr/local/bin/luc &&
// GOOS=linux GOARCH=amd64 go build -o %s
