package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.Println("Request started")
		defer log.Println("Request ended")

		select {
		case <-time.After(5 * time.Second):
			log.Println("Request processada com sucesso")
			w.Write([]byte("Request processada com sucesso"))
		case <-ctx.Done():
			log.Println("Request cancelada pelo cliente")
		}
	})
	http.ListenAndServe(":8080", nil)
}
