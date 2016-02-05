package main

import "fmt"
import "log"

func main() {
	fmt.Println("Hello")
	fmt.Println(promptAndGetInput())
}


func promptAndGetInput() string {
	var i string
	_, err := fmt.Scanf("%s", &i)
	if err != nil {
		log.Print("  Input Error ", err)
		i = ""
	}
	return i
}

