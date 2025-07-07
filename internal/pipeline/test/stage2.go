/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import "time"

func Stage2(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for item := range in {
			out <- "[stage2] " + item
			time.Sleep(2 * time.Second)
		}
	}()
	return out
}
