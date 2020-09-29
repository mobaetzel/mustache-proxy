package src

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
)

func RunService(host *string, port *string, configFile *string, debugMode *bool) {
	allowedTargets := readAllowedTargets(configFile)

	if *debugMode {
		fmt.Print("Found the following allowed targets:\n")
		for _, target := range allowedTargets {
			fmt.Printf("-> %s\n", target)
		}
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", createRequestHandler(allowedTargets))

	addr := fmt.Sprintf("%s:%s", *host, *port)

	if *debugMode {
		fmt.Printf("Listening on %s\n", addr)
	}

	http.ListenAndServe(addr, router)
}


