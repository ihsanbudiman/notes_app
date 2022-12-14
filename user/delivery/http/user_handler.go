package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ihsanbudiman/notes_app/domain"
	"github.com/ihsanbudiman/notes_app/helpers"
	"github.com/ihsanbudiman/notes_app/user/delivery/http/middleware"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(r *chi.Mux, u domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: u,
	}

	// make group v1
	r.Route("/user", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/register", helpers.RecoverWrap(handler.Register))
			r.Post("/login", helpers.RecoverWrap(handler.Login))
			r.With(middleware.MyMiddleware).Get("/", helpers.RecoverWrap(handler.FindUser))
		})
	})

}

func (u UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// get request body
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call usecase
	user, err = u.UserUsecase.Register(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helpers.HttpResponse{
		Message: "user created",
		Data: map[string]interface{}{
			"user": user,
		},
	}

	// return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (u UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// get request form body json
	req := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call usecase
	loginData, err := u.UserUsecase.Login(r.Context(), req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if user is empty
	if loginData.User.ID == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	response := helpers.HttpResponse{
		Message: "login success",
		Data: map[string]interface{}{
			"user":  loginData.User,
			"token": loginData.Token,
		},
	}

	// return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func (u UserHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	// get request form query params
	query := r.URL.Query()

	id := query.Get("id")
	if id == "" {
		http.Error(w, "id cannot be empty", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call usecase
	user, err := u.UserUsecase.FindUser(r.Context(), idInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := helpers.HttpResponse{
		Message: "user found",
		Data: map[string]interface{}{
			"user": user,
		},
	}

	// return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
