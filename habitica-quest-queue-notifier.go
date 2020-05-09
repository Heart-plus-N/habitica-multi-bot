package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

const host string = ":3333"

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// r.Post("/chat")
	// r.Post("/quest")

	log.Println("Listening at: ", host)
	err := http.ListenAndServe(host, r)
	log.Fatal(err)
}
