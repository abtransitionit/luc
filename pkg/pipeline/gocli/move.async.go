/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"
	"path"
	"sync"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

// Move File or folder to final destination
func Move(in <-chan PipelineData, out chan<- PipelineData, nbWorker int) {
	var wg sync.WaitGroup
	defer close(out)

	// Worker function
	worker := func() {
		defer wg.Done()

		for data := range in {

			if data.Err != nil {
				out <- data
				logx.L.Debugf("❌ Previous error detected")
				continue
			}

			// use cases
			switch data.Config.UrlType {
			case config.UrlTgz:
				data = helperMvTgz(data)
			case config.UrlExe:
				data = helperMvExe(data)
			default:
				logx.L.Debugf("[%s] UrlType '%s' is not supported or not yet managed", data.Config.Url, data.Config.UrlType)
			}

			out <- data
		} // for
	} // worker
	// Start N workers
	for i := 0; i < nbWorker; i++ {
		wg.Add(1)
		go worker()
	}
	wg.Wait()
}

func helperMvExe(data PipelineData) PipelineData {
	// define var
	dstFolder := data.DstFolder
	dstPath := path.Join(dstFolder, data.Config.Name)

	// log
	logx.L.Debugf("[%s] Moving '%s' to '%s'", data.Config.Name, data.ArtPath1, dstPath)

	// play code
	cli := fmt.Sprintf("luc do MoveFile %s %s %o %t", data.ArtPath1, dstPath, 0755, true)
	out, err := util.RunCLIRemote(data.HostName, cli)
	if err != nil {
		data.Err = fmt.Errorf("❌ Error: %v, %s", err, out)
		logx.L.Debugf("[%s][%s] ❌ Error detected 1", data.Config.Name, data.HostName)
		return data
	}

	// success
	logx.L.Infof("[%s] [%s] File moved to '%s'", data.Config.Name, data.HostName, dstPath)
	return data
}

func helperMvTgz(data PipelineData) PipelineData {
	// define var
	dstFolder := data.DstFolder
	dstPath := path.Join(dstFolder, data.Config.Name)

	// log
	logx.L.Debugf("[%s] Moving '%s' to '%s'", data.Config.Name, data.ArtPath1, dstPath)

	// play code
	cli := fmt.Sprintf("luc util mvdir %s %s %o %t %t --local", data.ArtPath2, dstPath, 0755, true, true)
	out, err := util.RunCLIRemote(data.HostName, cli)
	if err != nil {
		data.Err = fmt.Errorf("❌ Error: %v, %s", err, out)
		logx.L.Debugf("[%s][%s] ❌ Error detected 2", data.Config.Name, data.HostName)
		return data
	}

	// success
	logx.L.Infof("[%s] [%s] File moved to '%s'", data.Config.Name, data.HostName, dstPath)
	return data
}

// func helperMvTgz(data PipelineData) PipelineData {

// 	// define var
// 	dstFolder := data.DstFolder
// 	dstPath := path.Join(dstFolder, data.Config.Name)

// 	// log
// 	logx.L.Debugf("Moving '%s' to '%s'", data.ArtPath2, dstPath)

// 	// play code
// 	success, err := util.MvFolder(data.ArtPath2, dstPath, 0755, true, true)
// 	if err != nil {
// 		logx.L.Debugf("❌ Error detected 2")
// 		data.Err = err
// 		return data
// 	}
// 	if !success {
// 		logx.L.Debugf("❌ Error detected 20")
// 		data.Err = fmt.Errorf("[%s] Folder not moved to '%s'", data.Config.Name, dstPath)
// 		return data
// 	}

// 	// success
// 	logx.L.Infof("✅ [%s] Folder moved successfully to '%s'", data.Config.Name, dstPath)
// 	return data
// }
