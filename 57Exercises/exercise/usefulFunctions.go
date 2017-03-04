package exercise

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
func PromptAndReturnInputFromUser(promptMessage string) string {
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

//func open_file_return_contents_string(fileLocation string) string {
//	//file,_ := os.Open(fileLocation)
//	//var s string;
//	//file.WriteString("hellow")
//	//return s
//}
