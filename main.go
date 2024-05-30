package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviroment.")
		// log.Fatal(http.ListenAndServe(":8000", nil))
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	// put the http methode in a path
	v1Router := chi.NewRouter()

	// v1Router.HandleFunc("/healthz", handlerReadiness)
	v1Router.Get("/healthz", handlerReadiness)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	//fmt.Println("Port:", portString)
	log.Printf("Server starting on port %v", portString) //date
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
