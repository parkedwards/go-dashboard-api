package dashboard

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type DashboardManager interface{}

type DashboardRouter struct {
	manager DashboardManager
}

func New(dashboardManager DashboardManager) *DashboardRouter {
	return &DashboardRouter{
		manager: dashboardManager,
	}
}

func (r *DashboardRouter) RegisterRoutes(c chi.Router) {
	c.Get("/dashboard", r.getAllWorkoutsForDashboard)
}

func (r *DashboardRouter) getAllWorkoutsForDashboard(w http.ResponseWriter, req *http.Request) {

	fmt.Println("HEYOOO")
	// manager: get all users

	// manager: get all workouts for user
}
