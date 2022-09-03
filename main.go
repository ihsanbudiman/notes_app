package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/ihsanbudiman/notes_app/sqlcpg"
	user_handler "github.com/ihsanbudiman/notes_app/user/delivery/http"
	user_repo_pg "github.com/ihsanbudiman/notes_app/user/repository/postgres"
	user_ucase "github.com/ihsanbudiman/notes_app/user/usecase"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if os.Getenv("ENV") == "" {
		// load env
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// make connection to postgres
	var conn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDBNAME"))

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("failed to connect to database")
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to connect to database")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// cors origin

	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/testing", func(w http.ResponseWriter, r *http.Request) {
		type test struct {
			Success int `json:"success"`
			File    struct {
				URL string `json:"url"`
			} `json:"file"`
		}

		var t test
		t.Success = 1
		t.File.URL = "https://picsum.photos/200"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		ra, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(ra)

	})
	// close connection to postgres
	defer db.Close()

	sqlc := sqlcpg.New(db)

	userRepo := user_repo_pg.NewPostgresUserRepo(sqlc)
	userUseCase := user_ucase.NewUserUseCase(userRepo)
	user_handler.NewUserHandler(r, userUseCase)

	http.ListenAndServe(":3000", r)

}
