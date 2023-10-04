package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/leonyork/templ-csp-example/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.ServeHTTP)
	
	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	
	fmt.Printf("Server starting on http://%s", server.Addr)
	server.ListenAndServe()
}