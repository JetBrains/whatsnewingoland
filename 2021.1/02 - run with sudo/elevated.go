//+build !windows

package main

import (
	"os"
)

func isElevated() bool {
	return os.Getuid() == 0
}
