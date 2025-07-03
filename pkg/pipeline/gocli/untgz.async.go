/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func unTgz(in <-chan PipelineData, out chan<- PipelineData, nbWorker int) {
	var wg sync.WaitGroup
	defer close(out)

	// Worker function
	worker := func() {
		defer wg.Done()

		for data := range in {
			if data.Err != nil {
				out <- data
				logx.L.Debugf("[%s] [%s] ❌ Previous error detected", data.Config.Name, data.HostName)
				continue
			}

			// use cases
			switch data.Config.UrlType {
			case config.UrlTgz:
				data = helperUnTgz(data)
			default:
				logx.L.Debugf("[%s] [%s] UrlType '%s' is not impacted by this stage", data.Config.Name, data.HostName, data.Config.UrlType)
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

func helperUnTgz(data PipelineData) PipelineData {

	// set instance property
	uniqPath := filepath.Join("/tmp", fmt.Sprintf("%s_%d", data.Config.Name, time.Now().UnixNano()))
	data.ArtPath2 = uniqPath

	// Log
	logx.L.Debugf("[%s] [%s] Untaring artifact", data.Config.Name, data.HostName)

	// play code
	cli := fmt.Sprintf("luc util utgz %s %s --local", data.ArtPath1, data.ArtPath2)
	_, err := util.RunCLIRemote2(cli, data.HostName)
	if err != nil {
		logx.L.Debugf("[%s][%s] ❌ Error detected 1", data.Config.Name, data.HostName)
		data.Err = err
		return data
	}

	// success
	logx.L.Debugf("[%s] [%s] untarred artifact", data.Config.Name, data.HostName)
	return data
}
