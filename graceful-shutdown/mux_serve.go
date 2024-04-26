package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Thirawat Hello Sawadee Kub")
	mux := http.NewServeMux()

	mux.HandleFunc("/muxlai", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`mux hang lai`))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`hello`))
	})
	srv := http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	//goroutine
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	fmt.Println("Server Start at port name: 80")

	//buffer
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	log.Println("shutting down . . .")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Println("shutdown err:", err)
	}

	fmt.Println("Server Stop!")
}
