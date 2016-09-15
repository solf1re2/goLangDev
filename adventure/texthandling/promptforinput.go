package texthandling

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
Function called with input string used to prompt user for a single line input.
@promptMessage message to prompt user for input.
@returns user input
*/
func PromptAndRetrunInputFromUser(promptMessage string) string {
	var i string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(promptMessage)
	for scanner.Scan() {
		// assign the input to i
		i = scanner.Text()

		// Check that there is input
		if len(i) > 0 {
			break
		} else {
			fmt.Println(promptMessage)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Print(os.Stderr, "reading standard input:", err)
	}

	return i
}

// PromptAndGetInputString - returns i - user input string
func PromptAndGetInputString() string {
	var i string
	_, err := fmt.Scanf("%s", &i)
	if err != nil {
		log.Print("  Input Error ", err)
		i = ""
	}
	return i
}

//PromptAndGetInputInt - returns i- user input int
func PromptAndGetInputInt() int {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		log.Print("  Input Error ", err)
		i = 0
	}
	return i
}
