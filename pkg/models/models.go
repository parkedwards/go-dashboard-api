package models

type Client struct {
	Id            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Gender        string `json:"gender"`
	InitialWeight int64  `json:"initial_weight"`
}

type ClientWorkout struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Type         string `json:"type"`          // workout, rest
	ActivityType string `json:"activity_type"` // traditional_strength_training, running, or ""
	TrainerNotes string `json:"trainer_notes"`

	Duration       int64 `json:"duration"`
	ActualDuration int64 `json:"actual_duration"`

	Difficulty float64 `json:"difficulty,omitempty"`
	Notes      string  `json:"notes"`

	AverageHeartRate   int64   `json:"average_heart_rate"`
	MaxHeartRate       int64   `json:"max_heart_rate"`
	ActiveEnergyBurned int64   `json:"active_energy_burned"`
	Distance           float64 `json:"distance"`

	ScheduledAt string `json:"scheduled_at"`
	CompletedAt string `json:"completed_at"`
	MissedAt    string `json:"missed_at"`

	CompletionState string `json:"completion_state"` // full, partial, none (can be skipped OR rest)
}

// if you have time: do some sanitization on the workout payloads, and drop them in here
type WorkoutResponse struct {
	Name string `json:"name"`

	ScheduledAt string `json:"scheduled_at"`
	CompletedAt string `json:"completed_at"`
	MissedAt    string `json:"missed_at"`
	// add addt'l, normalized keys here (or use embedded struct)
}

type ClientResponse struct {
	Id            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Gender        string `json:"gender"`
	InitialWeight int64  `json:"initial_weight"`

	Workouts []ClientWorkout
}
