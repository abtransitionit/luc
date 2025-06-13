/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package gocli

import "github.com/abtransitionit/luc/pkg/deploy"

var Phases = []deploy.Phase{
	deploy.SetPhase("install", ep, EpDescription),
}
