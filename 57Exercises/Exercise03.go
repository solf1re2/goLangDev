package main


import (
	"fmt"
	"bufio"
	"os"
	"log"
)
func main() {
	//var quoteString = ""
	//fmt.Print("Hello, what is your name? ")
	promptForQuote := "What is the Quote? "
	promptForAuthor := "Who said it? "
	quoteString := promptForInputFromUser(promptForQuote)
	authorString := promptForInputFromUser(promptForAuthor)


	fmt.Println(authorString," says,  \"" , quoteString , "\"")
}

func promptForInputFromUser(promptMessage string) string {
	var i string
	//reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(promptMessage)
	//i, _ := reader.ReadString('\n')
	count := 0
	for scanner.Scan() {
		count++
		//fmt.Println(scanner.Text())
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
	//if err != nil {
	//	log.Print("  Input Error ", err)
	//	i = ""
	//}
	//if len(i) == 0 {
	//	fmt.Println("You must enter a value.")
	//	i = promptForInputFromUser(promptMessage)
	//}
	return i
}
