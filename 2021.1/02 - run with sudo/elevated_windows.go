package main

import "golang.org/x/sys/windows"

func isElevated() bool {
	return windows.GetCurrentProcessToken().IsElevated()
}
