package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gustsirt/BackendGo/handlers"
	"github.com/gustsirt/BackendGo/middleware"
	"github.com/gustsirt/BackendGo/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {

	api := r.PathPrefix("/api/v1").Subrouter()

	api.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/api/v1/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)

	api.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/posts/{postId}", handlers.DeletePostByIdHandler(s)).Methods(http.MethodDelete)
	api.HandleFunc("/posts/{postId}", handlers.UpdatePostByIdHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/posts/{postId}", handlers.GetPostByIDHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/ws", s.Hub().HandleWebSocket)
}

//en carpeta database : docker build . -t DbBackendGo
//                    : docker run -p 54321:5432 DbBackendGo

//en carpeta base     : docker build . -t BackendGo
//                    : docker run -p 5050:5050 BackendGo
