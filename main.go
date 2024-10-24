package main

import (
	"LoadBalancer/router"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	SpinServers()
	r := chi.NewRouter()
	// Use some middleware (like request logging)
	r.Use(middleware.Logger)

	// Define the /getData route
	r.Get("/testlb", router.GetLBHandler)
	// fmt.Println("Server starting at http://localhost", p)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}

	// Start a lb

}

func SpinServers() {

	go func() {
		r := chi.NewRouter()

		// Use some middleware (like request logging)
		r.Use(middleware.Logger)
		// Define the /getData route
		r.Get("/testlb", router.GetDataHandlerA)
		p := ":8081"
		// fmt.Println("Server starting at http://localhost", p)
		if err := http.ListenAndServe(p, r); err != nil {
			log.Fatal("Server failed to start:", err)
		}
	}()

	go func() {
		r := chi.NewRouter()

		// Use some middleware (like request logging)
		r.Use(middleware.Logger)
		// Define the /getData route
		r.Get("/testlb", router.GetDataHandlerB)
		p := ":8082"
		// fmt.Println("Server starting at http://localhost", p)
		if err := http.ListenAndServe(p, r); err != nil {
			log.Fatal("Server failed to start:", err)
		}
	}()

}
