package xsignal

import (
	"os"
	"syscall"
	"testing"
)

func kill(sig os.Signal) {
	pro, _ := os.FindProcess(os.Getpid())
	err := pro.Signal(sig)
	if err != nil {
		return
	}
}
func TestShutdownSIGQUIT(t *testing.T) {
	quit := make(chan struct{})
	t.Run("test shutdown signal by SIGQUIT", func(t *testing.T) {
		fn := func() {
			close(quit)
		}
		Shutdown(fn)
		kill(syscall.SIGQUIT)
		<-quit
	})
}
