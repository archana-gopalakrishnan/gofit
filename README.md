# GoFit - Workout Tracker

A command-line application for tracking walking, cycling, and running workouts.

## Running the Application

```bash
go run main.go
```

## Running Tests

```bash
go test -v
```

Or use the test script:
```bash
chmod +x run_tests.sh
./run_tests.sh
```

## Features

- Record workouts with date, time, duration, and distance
- List all workouts for a specific customer
- Data persisted in JSON file

## Code Smells for Clean Code Workshop

This application intentionally contains the following code smells:

### 1. **Magic Strings**
- Workout types ("walking", "cycling", "running") are hardcoded throughout
- Filename "workouts.json" is repeated multiple times
- Menu options and messages use magic strings

### 2. **Duplicate Code**
- Date/time validation logic is duplicated
- Workout type checking is repeated in multiple places
- File reading logic for "workouts.json" appears multiple times
- Display logic for workout types is duplicated

### 3. **Anemic Domain Model**
- The `Workout` struct has no behavior (no methods)
- All business logic is in procedural functions
- No encapsulation of workout-related operations

## Workshop Discussion Points

- How can we eliminate magic strings using constants or enums?
- Where can we extract common functionality to reduce duplication?
- How can we add behavior to the Workout model to make it richer?
- What validation logic belongs in the domain model?
- How can we improve separation of concerns?
