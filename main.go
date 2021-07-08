package main

import (
	"net/http"
	"pb3/db"
	"pb3/rest"

	"github.com/go-chi/chi"
)

func main() {
	db.InitDatabase("postgres://postgres:mysecretpassword@localhost:5432/problem3?sslmode=disable", 1)

	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/vehicle", rest.GetVehicle)
		r.Post("/vehicle", rest.PostVehicle)
		r.Delete("/vehicle", rest.DeleteVehicle)
		r.Get("/vehicle/candrive", rest.GetCanDrive)
	})
	http.ListenAndServe(":8080", router)
}
