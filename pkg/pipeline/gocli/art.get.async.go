/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gocli

import (
	"fmt"
	"sync"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/abtransitionit/luc/pkg/logx"
	"github.com/abtransitionit/luc/pkg/util"
)

func getArtifact(in <-chan PipelineData, out chan<- PipelineData, nbWorker int) {
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
			case config.UrlExe, config.UrlTgz:
				data = helperExeTgz(data)
			case config.UrlGit:
				// TODO
			default:
				data.Err = fmt.Errorf("[%s] ❌ Unsupported or not yet managed UrlType :  '%s'", data.Config.Name, data.Config.UrlType)
				logx.L.Debugf("[%s] ❌ Error detected 2", data.Config.Name)
			}

			out <- data
		} // for
	} // worker
	// Start N workers
	wg.Add(nbWorker)
	for i := 0; i < nbWorker; i++ {
		go worker()
	}

	wg.Wait()
}

func helperExeTgz(data PipelineData) PipelineData {

	// log
	logx.L.Debugf("[%s] downloading artifact for UrlType '%s'", data.Config.Name, data.Config.UrlType)

	// play code
	cli := fmt.Sprintf("luc util url get %s %s --local", data.HostUrl, data.ArtPath1)
	_, err := util.RunCLIRemote(data.HostName, cli)

	// error
	if err != nil {
		logx.L.Debugf("[%s] [%s] ❌ Error detected during download", data.Config.Name, data.HostName)
		data.Err = err
		return data
	}

	// success
	logx.L.Debugf("[%s] [%s] donwloaded artifact", data.Config.Name, data.HostName)
	return data
}
