package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port is not found")
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    port,
	}

	fmt.Printf("server is running on port %v\n", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
