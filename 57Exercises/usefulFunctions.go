package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
)
/*
Function called with input string used to prompt user for a single line input.
@promptMessage message to prompt user for input.
@returns user input
 */
func prompt_and_return_input_from_user(promptMessage string) string {
	var i string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(promptMessage)
	for scanner.Scan() {
		// assign the input to i
		i = scanner.Text()

		// Check that there is input
		if len(i) > 0 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Print(os.Stderr, "reading standard input:", err)
	}

	return i
}

