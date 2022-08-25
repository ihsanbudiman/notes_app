package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	// close connection to postgres
	defer db.Close()

	sqlc := sqlcpg.New(db)

	userRepo := user_repo_pg.NewPostgresUserRepo(sqlc)
	userUseCase := user_ucase.NewUserUseCase(userRepo)
	user_handler.NewUserHandler(r, userUseCase)

	http.ListenAndServe(":3000", r)

}
