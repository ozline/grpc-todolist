package utils

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ListenSignal(f func()) {
	signalCh := make(chan os.Signal, 5)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	log.Println("Listening signal...")
	select {
	case sig := <-signalCh:
		{
			log.Println("stopping service, because received signal:", sig)
			f()
			log.Println("service has stopped")
			os.Exit(0)
		}
	}
}
