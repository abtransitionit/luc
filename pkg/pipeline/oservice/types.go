/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package oservice

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/luc/pkg/config"
	"github.com/jedib0t/go-pretty/table"
)

// Map of services that shouldn't be created for each OS family
var excludedServices = map[string][]string{
	"rhel":   {"apparmor"},
	"debian": {"test1", "test2"},
}

func ServiceIsExcluded(key, item string) (bool, error) {
	//  input
	normalizedKey := strings.ToLower(strings.TrimSpace(key))
	normalizedItem := strings.ToLower(strings.TrimSpace(item))

	// error
	items, exists := excludedServices[normalizedKey]
	if !exists {
		return false, fmt.Errorf("key '%s' not found in excluded services", key)
	}

	// exist
	for _, excludedItem := range items {
		normalizedExcluded := strings.ToLower(strings.TrimSpace(excludedItem))
		if normalizedExcluded == normalizedItem {
			return true, nil
		}
	}

	// not exists
	return false, nil
}

type PipelineData struct {
	Config       config.OsServiceConfig
	HostName     string
	OsFamily     string // Rhel, Debian, fedora
	ServiceInfos string
	Err          error
}

// # Pupose
//
// Pretty print the Pipelined Data (usually for debugging)
func (obj PipelineData) String() string {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRows([]table.Row{
		{"Host name", obj.HostName},
		{"OS   Family", obj.OsFamily},
		{"Service name", obj.Config.Name},
		{"Service infos", obj.ServiceInfos},
		{"Name used", obj.Config.SName},
		{"Error", func() string {
			if obj.Err != nil {
				return obj.Err.Error()
			}
			return "-"
		}()},
	})

	return t.Render()
}
