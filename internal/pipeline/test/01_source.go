/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

// Source stage of the pipeline that define the data to be pipelined
func Source(out chan<- string, msgs []string) {
	go func() {
		defer close(out)
		for _, msg := range msgs {
			out <- msg
		}
	}()
}
