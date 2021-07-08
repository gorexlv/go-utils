//go:build windows
// +build windows

package xsignal

import (
	"os"
	"syscall"
)

var shutdownSignals = []os.Signal{syscall.SIGQUIT, os.Interrupt}
