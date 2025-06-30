/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import "strings"

// # Purpose
//
// - get the list of a Map:key from a map
//
// # Parameters
//
// - m: map[string]any
//
// # Return
//
// - []string: containing the keys of the map
//
// # Notes
//
// - This function accepts any kind of map
func GetMapKeys[V any](m map[string]V) []string {

	// defining size is more efficient
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key) // Add the current key to our slice.
	}

	return keys
}

// # Purpose
//
// - convert a []string to a String "item1 item2 item3 ..."
//
// # Parameters
//
// - ListKey: A []string slice containing the keys to be formatted.
//
// # Return
//
// - string:  A single string containing all the keys joined by the separator.
func GetStringfromSlice(ListString []string, separator string) string {
	return strings.Join(ListString, separator)
}

// # Purpose
//
// - convert a []string to a String "item1 item2 item3 ..."
//
// # Parameters
//
// - ListKey: A []string slice containing the keys to be formatted.
//
// # Return
//
// - string:  A single string containing all the keys separated by a space.
func GetStringfromSliceWithSpace(keys []string) string {
	return GetStringfromSlice(keys, " ")
}
