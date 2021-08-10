package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dexterlb/dumbloader/server"
)

func main() {
	srv := server.New()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        srv.Handler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

	fmt.Printf("hello")
}
