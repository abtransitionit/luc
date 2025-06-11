/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package gocli

import "github.com/abtransitionit/luc/pkg/deploy"

var Phases = []deploy.Phase{
	deploy.SetPhase("cli", cli, CliDescription),
	deploy.SetPhase("curl", curl, CurlDescription),
	deploy.SetPhase("dcp", dcp, DcpDescription),
	deploy.SetPhase("info", info, InfoDescription),
	deploy.SetPhase("mv", mv, MvDescription),
	deploy.SetPhase("dst", dst, DstDescription),
	deploy.SetPhase("path", pathx, PathDescription),
}
