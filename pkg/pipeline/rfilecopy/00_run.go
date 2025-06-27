/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package rfilecopy

const RunPipelineDescription = "copy a local file concurently to remote VMs"

// The map here could be VMName => destination folder
func RunPipeline(vmList string, srcFile string, dstFile string) error {
	// logx.L.Debug(RunPipelineDescription)
	// logx.L.Debugf("Copying '%s' to %d VM(s)", srcFile, len(vms))

	// Declare all pipeline channels
	chOutSource := make(chan PipelineData)
	chOutSshCheck := make(chan PipelineData)
	chOutCopy := make(chan PipelineData)
	chOutLast := chOutCopy

	// Start each pipeline stage
	go source(chOutSource, vmList, srcFile, dstFile)
	go checkVM(chOutSource, chOutSshCheck)
	go concurrentlyCopyFile(chOutSshCheck, chOutCopy, len(vmList))

	// Last step
	err := lastStep(chOutLast)
	if err != nil {
		return err
	}
	return nil
}
