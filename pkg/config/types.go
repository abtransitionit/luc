/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package config

type UrlType string

const (
	UrlExe UrlType = "exe"
	UrlTgz UrlType = "tgz"
	UrlGit UrlType = "git"
	UrlGo  UrlType = "go"
	UrlXxx UrlType = "xxx"
	// etc.
)

type CLIConfig struct {
	Name    string
	Tag     string
	Url     string
	DocUrl  string
	GitUrl  string
	UrlType UrlType
}
