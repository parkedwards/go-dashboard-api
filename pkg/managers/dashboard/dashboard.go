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
func httpGet(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		// handle failed fetch
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func (m *DashboardManager) GetAllClients() []models.Client {
	body := httpGet("https://sandbox.future.fit/users")
	data := []models.Client{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func (m *DashboardManager) GetAllWorkoutsForClient(clientId string) []models.ClientWorkout {
	body := httpGet(fmt.Sprintf("https://sandbox.future.fit/users/%v/workouts", clientId))
	data := []models.ClientWorkout{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
