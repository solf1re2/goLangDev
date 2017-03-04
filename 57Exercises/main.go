package main

import (
	"fmt"
	"strconv"
	"github.com/solf1re2/goLangDev/57Exercises/exercise"
)

// main - initial function, will prompt user for exercise to run.
func main() {
	prompt := "Enter Exercise number to run (1-4):"
	exerciseToRunString := exercise.PromptAndReturnInputFromUser(prompt)
	// Check exerciseToRun is an int within the selectable range
	exerciseToRunInt, _ := strconv.ParseInt(exerciseToRunString, 10, 64)
	runExercise(exerciseToRunInt, prompt)
}

func runExercise(e int64, errmsg string) {
	switch e {
	case 1:
		exercise.ExerciseOne()
	case 2:
		exercise.ExerciseTwo()
	case 3:
		exercise.ExerciseThree()
	case 4:
		exercise.ExerciseFour()
	default:
		fmt.Println(errmsg)
	}
}
