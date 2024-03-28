package loot

import "fmt"

var verbose bool = false

func IsVerbose() bool {
	return verbose
}

func SetVerbose(v bool) {
	if v == true {
		fmt.Println("I'll be more verbose...")
	}
	verbose = v
}

func doIfVerbose(f func()) {
	if IsVerbose() {
		f()
	}
}
