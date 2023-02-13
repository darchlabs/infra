package util

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// :istenInterrupt method used to listen SIGTERM OS Signal
func ListenInterrupt(quit chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-c
		fmt.Println("Signal received", s.String())
		quit <- struct{}{}
	}()
}

// GracefullShutdown method used to close all synchronizer processes
func GracefullShutdown() {
	log.Println("Gracefully shutdown")
}
