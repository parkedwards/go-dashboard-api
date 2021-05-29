package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	DashboardManager "github.com/parkedwards/go-dashboard-api/pkg/managers/dashboard"
	DashboardRouter "github.com/parkedwards/go-dashboard-api/pkg/routers/dashboard"
)

const (
	port = ":8000"
)

// @title Dashboard API
// @version 1.0
// @description Pull data on client workouts + statuses.
func Init() *chi.Mux {
	r := chi.NewRouter()

	manager := DashboardManager.New()

	router := DashboardRouter.New(manager)

	router.RegisterRoutes(r)

	return r
}

func Boot(router *chi.Mux) {
	fmt.Printf("bootin it up at %v", port)
	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
