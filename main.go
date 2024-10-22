package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/damirpavlik/meal-planning/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("port not found")
	}

	dbURL := os.Getenv("DB_URL")
	if port == "" {
		log.Fatal("db url not found")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't con to db: ", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
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
	v1router := chi.NewRouter()
	v1router.Get("/healtz", handlerHealthz)
	v1router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("srv started on: %v", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
