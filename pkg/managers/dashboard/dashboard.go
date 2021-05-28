package dashboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/parkedwards/go-dashboard-api/pkg/models"
)

type DashboardManager struct{}

func New() *DashboardManager {
	return &DashboardManager{}
}

// external HTTP GET helper
func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("http get failed! %v", err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, err
}

func (m *DashboardManager) GetAllClients() []models.Client {
	body, _ := httpGet("https://sandbox.future.fit/users")
	data := []models.Client{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func (m *DashboardManager) GetAllWorkoutsForClient(clientId string) []models.ClientWorkout {
	body, _ := httpGet(fmt.Sprintf("https://sandbox.future.fit/users/%v/workouts", clientId))
	data := []models.ClientWorkout{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func (m *DashboardManager) NormalizeWorkoutForDashboard(workout models.ClientWorkout) models.WorkoutResponse {

	// some computation needed here to flag as REST or SKIPPED
	workoutName := workout.Name
	if workoutName == "" && workout.Type == "rest" {
		workoutName = "REST"
	}

	workoutResponse := models.WorkoutResponse{
		Name: workoutName,
		// ScheduledAt
		// CompletedAt string `json:"completed_at"`
		// MissedAt    string `json:"missed_at"`
	}

	return workoutResponse
}
