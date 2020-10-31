package main

import (
	"net/http"

	"github.com/vSterlin/jwt-auth/cors"
	"github.com/vSterlin/jwt-auth/handlers"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.PostHandler).Methods(http.MethodPost)
	r.HandleFunc("/signin", handlers.LoginHandler).Methods(http.MethodPost)

	http.ListenAndServe("localhost:8080", cors.SetCORS(r))
}
