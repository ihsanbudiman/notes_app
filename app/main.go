package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ihsanbudiman/notes_app/sqlcpg"
	user_handler "github.com/ihsanbudiman/notes_app/user/delivery/http"
	user_repo_pg "github.com/ihsanbudiman/notes_app/user/repository/postgres"
	user_ucase "github.com/ihsanbudiman/notes_app/user/usecase"
	_ "github.com/lib/pq"
)

const (
	host     = "103.171.85.68"
	port     = 5433
	user     = "postgres"
	password = "postgress"
	dbname   = "notes_app"
)

func main() {
	// make connection to postgres
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
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

	data, err := sqlc.FindUserByEmail(context.Background(), sql.NullString{
		String: "ninoidgt@gmail.coms",
		Valid:  true,
	})
	fmt.Println(data, err)

	userRepo := user_repo_pg.NewPostgresUserRepo(sqlc)
	userUseCase := user_ucase.NewUserUseCase(userRepo)
	user_handler.NewUserHandler(r, userUseCase)

	http.ListenAndServe(":3000", r)

}
