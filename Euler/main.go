// Main go package, will handle asking user which Project Euler Problem they wish to run and call that function from the problem package
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	index "github.com/solf1re2/goLangDev/Euler/problemindex"
)

func main() {
	p := promptUserForProblem("problem to run")
	index.RunProgram(p)
}

func promptUserForProblem(promptString string) int {

	fmt.Printf("Select %s (enter 1,2,etc)\n\n", promptString)

	//Wait for input from user
	selection := returnInputFromUser()

	number, err := strconv.Atoi(selection)
	if err != nil {
		// handle error
		log.Fatalln("strconv.Atoi error")
	}
	// if number > len(list) {
	// 	log.Fatalf("Selection %v not recognised.", number)
	// }

	log.Printf("User has selected %s: %s\n", promptString, selection)

	problem := number

	return problem
}

/*
Function called with input string used to prompt user for a single line input.
@promptMessage message to prompt user for input.
@returns user input
*/
func returnInputFromUser() string {
	fmt.Print("User Selection: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	firstInput := string([]byte(input)[0])
	// log.Println(firstInput)
	// fmt.Printf("Input Char Is : %v", string([]byte(input)[0]))
	return firstInput
}
