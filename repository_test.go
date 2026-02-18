package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestWorkoutRepository_Save(t *testing.T) {
	testFile := "test_workouts.json"
	defer os.Remove(testFile)

	// Temporarily replace the filename
	originalFile := "workouts.json"
	replaceFilename(testFile)
	defer replaceFilename(originalFile)

	repo := &WorkoutRepository{}

	workout := Workout{
		CustomerID: "C001",
		Type:       "running",
		Date:       "2024-01-15",
		Time:       "08:30",
		Duration:   30,
		Distance:   5000,
	}

	err := repo.Save(workout)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	var workouts []Workout
	json.Unmarshal(data, &workouts)

	if len(workouts) != 1 {
		t.Errorf("Expected 1 workout, got %d", len(workouts))
	}

	if workouts[0].CustomerID != "C001" {
		t.Errorf("Expected CustomerID C001, got %s", workouts[0].CustomerID)
	}
}

func TestWorkoutRepository_SaveMultiple(t *testing.T) {
	testFile := "test_workouts_multiple.json"
	defer os.Remove(testFile)

	originalFile := "workouts.json"
	replaceFilename(testFile)
	defer replaceFilename(originalFile)

	repo := &WorkoutRepository{}

	workout1 := Workout{CustomerID: "C001", Type: "walking", Date: "2024-01-15", Time: "08:00", Duration: 20, Distance: 2000}
	workout2 := Workout{CustomerID: "C002", Type: "cycling", Date: "2024-01-16", Time: "09:00", Duration: 45, Distance: 15000}

	repo.Save(workout1)
	repo.Save(workout2)

	data, _ := os.ReadFile(testFile)
	var workouts []Workout
	json.Unmarshal(data, &workouts)

	if len(workouts) != 2 {
		t.Errorf("Expected 2 workouts, got %d", len(workouts))
	}
}

func TestWorkoutRepository_Fetch(t *testing.T) {
	testFile := "test_workouts_fetch.json"
	defer os.Remove(testFile)

	originalFile := "workouts.json"
	replaceFilename(testFile)
	defer replaceFilename(originalFile)

	workouts := []Workout{
		{CustomerID: "C001", Type: "running", Date: "2024-01-15", Time: "08:00", Duration: 30, Distance: 5000},
		{CustomerID: "C002", Type: "walking", Date: "2024-01-16", Time: "09:00", Duration: 20, Distance: 2000},
		{CustomerID: "C001", Type: "cycling", Date: "2024-01-17", Time: "10:00", Duration: 45, Distance: 15000},
	}

	data, _ := json.MarshalIndent(workouts, "", "  ")
	os.WriteFile(testFile, data, 0644)

	repo := &WorkoutRepository{}
	result := repo.Fetch("C001")

	if len(result) != 2 {
		t.Errorf("Expected 2 workouts for C001, got %d", len(result))
	}

	for _, w := range result {
		if w.CustomerID != "C001" {
			t.Errorf("Expected CustomerID C001, got %s", w.CustomerID)
		}
	}
}

func TestWorkoutRepository_FetchNoResults(t *testing.T) {
	testFile := "test_workouts_empty.json"
	defer os.Remove(testFile)

	originalFile := "workouts.json"
	replaceFilename(testFile)
	defer replaceFilename(originalFile)

	workouts := []Workout{
		{CustomerID: "C001", Type: "running", Date: "2024-01-15", Time: "08:00", Duration: 30, Distance: 5000},
	}

	data, _ := json.MarshalIndent(workouts, "", "  ")
	os.WriteFile(testFile, data, 0644)

	repo := &WorkoutRepository{}
	result := repo.Fetch("C999")

	if len(result) != 0 {
		t.Errorf("Expected 0 workouts for C999, got %d", len(result))
	}
}

func TestWorkoutRepository_FetchFileNotExists(t *testing.T) {
	testFile := "nonexistent.json"

	originalFile := "workouts.json"
	replaceFilename(testFile)
	defer replaceFilename(originalFile)

	repo := &WorkoutRepository{}
	result := repo.Fetch("C001")

	if len(result) != 0 {
		t.Errorf("Expected 0 workouts when file doesn't exist, got %d", len(result))
	}
}

// Helper function to replace filename in repository (for testing purposes)
func replaceFilename(newFile string) {
	filename = newFile
}
