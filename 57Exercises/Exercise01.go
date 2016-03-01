package main

import (
	"fmt"
	"log"
)

func main() {
	var i = ""
	fmt.Print("Hello, what is your name?")
	i = getInput()

	fmt.Println("Hello" , i , "nice to meet you")
}


func getInput() string {
	var i string
	_, err := fmt.Scanf("%s", &i)

	if err != nil {
		log.Print("  Input Error ", err)
		i = ""
	}
	if len(i) == 0 {
		fmt.Println("You must enter a name.")
		return getInput()
	}
	return i
}