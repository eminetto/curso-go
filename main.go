package main

import (
	"encoding/json"
	"net/http"

	"github.com/eminetto/curso-go/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", calculaChurrasco)
	http.ListenAndServe(":3000", r)
}

func calculaChurrasco(w http.ResponseWriter, r *http.Request) {
	var param domain.Parametros
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	ch, err := domain.CalculaChurrasco(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(ch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
