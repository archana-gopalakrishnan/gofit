# Code Smells in GoFit Application

This document lists all the intentional code smells in the application for the clean code workshop.

## 1. Magic Strings

### Workout Types
- **Location**: `main.go` - `recordWorkout()` function (lines 64-67)
- **Issue**: Hardcoded strings "walking", "cycling", "running" repeated multiple times
- **Also appears in**: `listWorkouts()` function (lines 133-139)

### Date/Time Format Strings
- **Location**: `main.go` - `recordWorkout()` function
- **Issue**: Format strings "2006-01-02" (line 75) and "15:04" (line 85) are hardcoded

### Filename
- **Location**: `repository.go` - package variable
- **Issue**: "workouts.json" appears as a magic string

### Menu Options and Messages
- **Location**: `main.go` - `main()` function
- **Issue**: Menu strings "1", "2", "3" are hardcoded (lines 38-40)

## 2. Duplicate Code

### Date/Time Validation Logic
- **Location**: `main.go` - `recordWorkout()` function
- **Issue**: Similar validation pattern repeated for both date (lines 73-77) and time (lines 83-87)
- **Pattern**: Parse, check error, print message, return

### File Reading Logic
- **Location**: `repository.go`
- **Issue**: File reading code duplicated in both `Save()` (lines 13-17) and `Fetch()` (lines 26-29) methods
- **Pattern**: ReadFile, unmarshal JSON

### Workout Type Display Logic
- **Location**: `main.go` - `listWorkouts()` function (lines 133-139)
- **Issue**: Repetitive if-else chain for displaying workout types
- **Could be**: A map or method on Workout struct

### Speed Calculation
- **Location**: `main.go` - `listWorkouts()` function (lines 146-149)
- **Issue**: Calculation logic embedded in display function
- **Could be**: A method on Workout struct

## 3. Anemic Domain Model

### Workout Struct
- **Location**: `main.go` (lines 13-20)
- **Issue**: The Workout struct is just a data container with no behavior
- **Missing methods**:
  - `IsValid()` - to validate workout data
  - `CalculateSpeed()` - to compute average speed
  - `Display()` - to format output
  - `GetFormattedType()` - to return capitalized type name

### No Encapsulation
- **Issue**: All business logic is in procedural functions rather than methods on domain objects
- **Examples**:
  - Validation logic in `recordWorkout()` should be in Workout
  - Display logic in `listWorkouts()` should be in Workout
  - Speed calculation should be in Workout

## 4. Other Code Smells

### Long Method
- **Location**: `main.go` - `recordWorkout()` function
- **Issue**: Function does too many things (input, validation, saving)
- **Could be split into**: Input gathering, validation, persistence

### Primitive Obsession
- **Issue**: Using strings for workout types instead of a proper type/enum
- **Better approach**: Define a WorkoutType type with constants

### Error Handling
- **Location**: Throughout the code
- **Issue**: Errors are ignored with `_` or not properly handled
- **Examples**: JSON unmarshal errors, file write errors

## Workshop Refactoring Goals

1. Extract constants for all magic strings
2. Create helper functions to eliminate duplicate validation logic
3. Add methods to Workout struct to make it a rich domain model
4. Consider creating a WorkoutType enum/type
5. Improve error handling throughout
6. Extract file operations to eliminate duplication in repository
