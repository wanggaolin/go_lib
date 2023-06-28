package system

import "os"

func ExitError() {
	os.Exit(1)
}

func ExitSuccess() {
	os.Exit(0)
}
