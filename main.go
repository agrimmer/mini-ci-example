package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type myHandler struct {
}

func (myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "Good day, it is %02d:%02d", t.Hour(), getMinute(t.Minute(), t.Second()))
}

func getMinute(minute int, second int) int {
	return minute + second/30
}

func main() {

	srv := &http.Server{
		Addr:         ":8888",
		Handler:      &myHandler{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			// extra handling here
			cancel()
		}()

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
