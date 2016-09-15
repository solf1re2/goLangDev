package main

import (
	"fmt"
	"strconv"
)

// main - initial function, will prompt user for exercise to run.
func main() {
	prompt := "Enter Exercise number to run (1-4):"
	exerciseToRunString := promptAndReturnInputFromUser(prompt)
	// Check exerciseToRun is an int within the selectable range
	exerciseToRunInt, _ := strconv.ParseInt(exerciseToRunString, 10, 64)
	runExercise(exerciseToRunInt, prompt)
}

func runExercise(e int64, errmsg string) {
	switch e {
	case 1:
		exerciseOne()
	case 2:
		exerciseTwo()
	case 3:
		exerciseThree()
	case 4:
		exerciseFour()
	default:
		fmt.Println(errmsg)
	}
}
