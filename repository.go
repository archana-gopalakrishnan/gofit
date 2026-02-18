package main

import (
	"encoding/json"
	"os"
)

var filename = "workouts.json"

type WorkoutRepository struct{}

func (r *WorkoutRepository) Save(workout Workout) error {
	file, err := os.ReadFile(filename)
	var workouts []Workout
	if err == nil {
		json.Unmarshal(file, &workouts)
	}

	workouts = append(workouts, workout)

	data, _ := json.MarshalIndent(workouts, "", "  ")
	os.WriteFile(filename, data, 0644)

	return nil
}

func (r *WorkoutRepository) Fetch(customerId string) []Workout {
	file, err := os.ReadFile(filename)
	if err != nil {
		return []Workout{}
	}

	var workouts []Workout
	json.Unmarshal(file, &workouts)

	var customerWorkouts []Workout
	for _, w := range workouts {
		if w.CustomerID == customerId {
			customerWorkouts = append(customerWorkouts, w)
		}
	}

	return customerWorkouts
}
