package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	port = ":8000"
)

// @title Dashboard API
// @version 1.0
// @description Pull data on client workouts + statuses.
func Init() *chi.Mux {
	router := chi.NewRouter()
	return router
}

func Boot(router *chi.Mux) {
	fmt.Printf("bootin it up at %v", port)
	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
