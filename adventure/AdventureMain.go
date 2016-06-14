package main

import "fmt"
import "log"

func main() {
	fmt.Println("Hello")
	fmt.Println(promptAndGetInputString())
//	fmt.Println(promptAndGetInputInt())
}


func promptAndGetInputString() string {
	var i string
	_, err := fmt.Scanf("%s", &i)
	if err != nil {
		log.Print("  Input Error ", err)
		i = ""
	}
	return i
}

func promptAndGetInputInt() int {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		log.Print("  Input Error ", err)
		i = 0
	}
	return i
}