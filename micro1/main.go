package main

import (
	"context"
	"log"
	"micro1/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "ns", log.LstdFlags)

	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)
	//serve mux registers handler and returns info stored in the
	sm := http.NewServeMux()
	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gh)
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// since listen and sever are blocking, we put it in a go func

	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	// following will broadcast message on channel when os recieves killed or interrupted
	// command
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("recieved terminate, gracefull shutdown", sig)

	// graceful shutdown ==> wait till all the requests in the queue has finished and
	// will stop taking further requests

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)

}
