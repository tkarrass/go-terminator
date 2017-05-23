package terminator // udico.de/terminator

import (
  "os"
	"os/signal"
	"syscall"
)

var Terminator = make(chan struct{})

func init() {
  sigs := make(chan os.Signal)
  signal.Notify(sigs, os.Interrupt, os.Kill)
	signal.Notify(sigs, syscall.SIGTERM)

	// When notified about SIGINT, SIGTERM, or SIGKILL close the Terminator channel
	// in order to notify interested parties to quit.
	go func(c <-chan os.Signal) {
		<-c
		close(Terminator)
	}(sigs)
}
