package utils

import (
	"fmt"
	"jgttech/ypm/tui"
	"time"
)

// The timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
//
// defer timer("main")()
func Timer(name string) func() {
	start := time.Now()

	return func() {
		s := fmt.Sprintf("%s took %v", name, time.Since(start))
		fmt.Printf(tui.DarkGray(s) + "\n")
	}
}
