/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package gocliold

import "github.com/abtransitionit/luc/pkg/phase"

var Phases = []phase.Phase{
	phase.SetPhase("install", RunPipeline, EpDescription),
	// phase.SetPhase("install2", ep2, EpDescription),
}

// each phase launch a pipeline
