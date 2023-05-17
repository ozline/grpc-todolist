package utils

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ListenSignal(f func()) {
	signalCh := make(chan os.Signal, 5)

	signal.Notify(signalCh, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	log.Println("listening signal...")

	for {
		sig := <-signalCh

		if sig == syscall.SIGQUIT || sig == syscall.SIGTERM || sig == syscall.SIGINT {
			log.Println("stopping service")

			f()
			log.Println("service stopped")
			os.Exit(0)
		}
	}
}
