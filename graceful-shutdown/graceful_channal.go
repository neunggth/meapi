package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("server started")
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	log.Println("wait of signal")
	<-stop
	log.Println("Server Stopped")

}
