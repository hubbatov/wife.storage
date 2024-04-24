package router

import (
	"fmt"
	"log"
	"net/http"
	"wife/internal/user"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func NewUserRouter(e *user.Endpoint) *Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", e.RegisterUser).Methods("POST")
	r.HandleFunc("/login", e.Login).Methods("POST")

	r.HandleFunc("/users", e.GetUsers).Methods("GET")

	return &Router{
		router: r,
	}
}

func (r *Router) Run(port int) {
	log.Print("Router started on port: ", port)
	http.ListenAndServe(":"+fmt.Sprint(port), r.router)
}
