package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Workout struct {
	CustomerID string `json:"customerId"`
	Type       string `json:"type"`
	Date       string `json:"date"`
	Time       string `json:"time"`
	Duration   int    `json:"duration"`
	Distance   int    `json:"distance"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	repo := &WorkoutRepository{}

	for {
		fmt.Println("\n=== Workout Tracker ===")
		fmt.Println("1. Record Workout")
		fmt.Println("2. List Workouts")
		fmt.Println("3. Exit")
		fmt.Print("Choose option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "1" {
			recordWorkout(reader, repo)
		} else if choice == "2" {
			listWorkouts(reader, repo)
		} else if choice == "3" {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Invalid option")
		}
	}
}

func recordWorkout(reader *bufio.Reader, repo *WorkoutRepository) {
	var workout Workout

	fmt.Print("Enter Customer ID: ")
	customerId, _ := reader.ReadString('\n')
	workout.CustomerID = strings.TrimSpace(customerId)

	fmt.Print("Enter workout type (walking/cycling/running): ")
	workoutType, _ := reader.ReadString('\n')
	workout.Type = strings.TrimSpace(workoutType)

	if workout.Type != "walking" && workout.Type != "cycling" && workout.Type != "running" {
		fmt.Println("Invalid workout type! Must be walking, cycling or running")
		return
	}

	fmt.Print("Enter date (YYYY-MM-DD): ")
	date, _ := reader.ReadString('\n')
	workout.Date = strings.TrimSpace(date)

	_, err := time.Parse("2006-01-02", workout.Date)
	if err != nil {
		fmt.Println("Invalid date format! Use YYYY-MM-DD")
		return
	}

	fmt.Print("Enter time (HH:MM): ")
	timeStr, _ := reader.ReadString('\n')
	workout.Time = strings.TrimSpace(timeStr)

	_, err = time.Parse("15:04", workout.Time)
	if err != nil {
		fmt.Println("Invalid time format! Use HH:MM")
		return
	}

	fmt.Print("Enter duration (minutes): ")
	durationStr, _ := reader.ReadString('\n')
	duration, err := strconv.Atoi(strings.TrimSpace(durationStr))
	if err != nil {
		fmt.Println("Invalid duration!")
		return
	}
	workout.Duration = duration

	fmt.Print("Enter distance (metres): ")
	distanceStr, _ := reader.ReadString('\n')
	distance, err := strconv.Atoi(strings.TrimSpace(distanceStr))
	if err != nil {
		fmt.Println("Invalid distance!")
		return
	}
	workout.Distance = distance

	repo.Save(workout)
	fmt.Println("Workout recorded successfully!")
}

func listWorkouts(reader *bufio.Reader, repo *WorkoutRepository) {
	fmt.Print("Enter Customer ID: ")
	customerId, _ := reader.ReadString('\n')
	customerId = strings.TrimSpace(customerId)

	workouts := repo.Fetch(customerId)

	if len(workouts) == 0 {
		fmt.Println("No workouts found for this customer!")
		return
	}

	fmt.Println("\n=== Your Workouts ===")
	for _, w := range workouts {
		if w.Type == "walking" {
			fmt.Printf("\nType: Walking\n")
		} else if w.Type == "cycling" {
			fmt.Printf("\nType: Cycling\n")
		} else if w.Type == "running" {
			fmt.Printf("\nType: Running\n")
		}
		fmt.Printf("Date: %s\n", w.Date)
		fmt.Printf("Time: %s\n", w.Time)
		fmt.Printf("Duration: %d minutes\n", w.Duration)
		fmt.Printf("Distance: %d metres\n", w.Distance)

		if w.Duration > 0 {
			speed := float64(w.Distance) / float64(w.Duration)
			fmt.Printf("Average Speed: %.2f metres/minute\n", speed)
		}
	}
}
