package main

import "fmt"
import (
	"log"
)
func main() {
	var i = ""
	fmt.Print("Hello, what is your name? ")
	i = getInput2()

	fmt.Println(i,"has" , len(i) , "characters")
}

func getInput2() string {
	var i string
	_, err := fmt.Scanf("%s", &i)

	if err != nil {
		log.Print("  Input Error ", err)
		i = ""
	}
	if len(i) == 0 {
		fmt.Println("You must enter a name.")
		return getInput2()
	}
	return i
}