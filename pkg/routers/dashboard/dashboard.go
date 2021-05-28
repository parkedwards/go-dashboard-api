package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/parkedwards/go-dashboard-api/pkg/models"
)

type DashboardManager interface {
	GetAllClients() []models.Client
	GetAllWorkoutsForClient(clientId string) []models.ClientWorkout
}

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
	dashboardPayload := []models.ClientResponse{}

	clients := r.manager.GetAllClients()

	// ideally, handle the first clients GET failure here + exit

	for _, c := range clients {
		clientWorkouts := r.manager.GetAllWorkoutsForClient(c.Id)

		// ideally, handle any subsequent workout GET here + exit

		payload := models.ClientResponse{
			Id:            c.Id,
			FirstName:     c.FirstName,
			LastName:      c.LastName,
			Gender:        c.Gender,
			InitialWeight: c.InitialWeight,
			Workouts:      clientWorkouts,
		}

		dashboardPayload = append(dashboardPayload, payload)
	}

	result, err := json.Marshal(dashboardPayload)

	fmt.Println(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
