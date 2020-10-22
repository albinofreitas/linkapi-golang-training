package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/albinofreitas/linkapi-golang/database"
	"github.com/albinofreitas/linkapi-golang/internal/orders"
	"github.com/gorilla/mux"
)

// ResponseMessage sisi
type ResponseMessage struct {
	Message  string `json:"message"`
	Password string `json:"-"`
}

func main() {
	database.Connect()

	r := mux.NewRouter()
	r.Use(setHeaders)
	r.HandleFunc("/", orders.Store).Methods("POST")
	r.HandleFunc("/{id}", orders.Show).Methods("GET")

	srv := &http.Server{
		Addr:         "0.0.0.0:4040",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		srv.ListenAndServe()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func setHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
