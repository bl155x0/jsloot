package loot

var verbose bool = false

func IsVerbose() bool {
	return verbose
}

func SetVerbose(v bool) {
	verbose = v
}

func doIfVerbose(f func()) {
	if IsVerbose() {
		f()
	}
}
