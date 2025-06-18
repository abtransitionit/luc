/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package util

import (
	"errors"
)

// computes the Cartesian product of two string slices.
//
// Returns:
//
//   - []string: containing the Cartesian product
//   - error:    if either input slice is empty
//
// Time complexity: O(n*m) where n = len(a), m = len(b)
// Space complexity: O(n*m)
func CartesianProduct(a, b []string) ([]string, error) {
	if len(a) == 0 || len(b) == 0 {
		return nil, errors.New("empty input slice")
	}

	var product []string
	for _, x := range a {
		for _, y := range b {
			product = append(product, x+y)
		}
	}
	return product, nil
}
