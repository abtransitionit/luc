/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"fmt"
	"path/filepath"

	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// locally create an empty file as current user
func touchFileLocal(folderPath string, fileName string) error {
	filePath := filepath.Join(folderPath, fileName)

	// touch file
	cli := fmt.Sprintf(`touch %s`, filePath)
	_, err := util.RunCLILocal(cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return fmt.Errorf("failed to touch file locally: %w", err)
	}
	logx.L.Infof("✅ touched local file without garantee: %s", filePath)
	return nil
}

func deleteFileLocal(folderPath string, fileName string) error {
	filePath := filepath.Join(folderPath, fileName)

	// delete file
	cli := fmt.Sprintf(`rm %s`, filePath)
	_, err := util.RunCLILocal(cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return fmt.Errorf("failed to delete file locally: %w", err)
	}
	logx.L.Infof("✅ deleted local file without garantee: %s", filePath)
	return nil
}

func touchFileOnRemote(vm string, folderPath string, fileName string) error {
	filePath := filepath.Join(folderPath, fileName)

	// touch file
	cli := fmt.Sprintf(`touch %s`, filePath)
	_, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return fmt.Errorf("failed to touch file remotely: %w", err)
	}
	logx.L.Infof("✅ touched remote file without garantee: %s", filePath)
	return nil
}

func deleteFileOnRemote(vm string, folderPath string, fileName string) error {
	filePath := filepath.Join(folderPath, fileName)

	// delete
	cli := fmt.Sprintf(`rm %s`, filePath)
	_, err := util.RunCLIRemote(vm, cli)
	if err != nil {
		logx.L.Debugf("%s", err)
		return fmt.Errorf("failed to delete file remotely: %w", err)
	}
	logx.L.Infof("✅ deleted remote file without garantee: %s", filePath)
	return nil
}
