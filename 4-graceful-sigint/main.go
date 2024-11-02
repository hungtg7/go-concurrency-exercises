//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a process
	proc := MockProcess{}

	os_sig_ch := make(chan os.Signal)
	signal.Notify(os_sig_ch, syscall.SIGINT)

	// Run the process (blocking)
	go proc.Run()

	<-os_sig_ch
	fmt.Println("Attemp graceful shutdown")

	go proc.Stop()

	select {
	case <-os_sig_ch:
		fmt.Println("Force shutdown")
	case <-time.After(5 * time.Second):
		fmt.Println("gracefully shutdown success")
	}
}
