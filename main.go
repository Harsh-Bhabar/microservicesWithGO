package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Harsh-Bhabar/products-api/handlers"
)

func main() {
	fmt.Println("Starting")

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)
	byeHandler := handlers.NewBye(logger)

	serverMux := http.NewServeMux()
	serverMux.Handle("/hello", helloHandler)
	serverMux.Handle("/bye", byeHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// graceful shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Reacreiced Terminate, graceful shutdown ", sig)

	server.ListenAndServe()

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
