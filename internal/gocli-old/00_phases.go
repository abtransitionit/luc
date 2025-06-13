/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package gocli

import (
	"github.com/abtransitionit/luc/pkg/phase"
)

var Phases = []phase.Phase{
	phase.SetPhase("cli", cli, CliDescription),
	phase.SetPhase("curl", curl, CurlDescription),
	phase.SetPhase("dcp", dcp, DcpDescription),
	phase.SetPhase("info", info, InfoDescription),
	phase.SetPhase("mv", mv, MvDescription),
	phase.SetPhase("dst", dst, DstDescription),
	phase.SetPhase("path", pathx, PathDescription),
}
