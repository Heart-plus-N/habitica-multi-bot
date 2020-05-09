package main

import (
	"log"
	"net/http"

	"os"

	"github.com/go-chi/chi"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port, nil
}

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// r.Post("/chat")
	// r.Post("/quest")

	host, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening at: ", host)
	err = http.ListenAndServe(host, r)
	log.Fatal(err)
}
