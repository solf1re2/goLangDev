package main

import "fmt"

func exerciseOne() {
	var i = ""
	prompt := "Hello, what is your name? "
	// i = getInput()
	i = promptAndReturnInputFromUser(prompt)
	fmt.Println("Hello", i, "nice to meet you")
}

// func getInput() string {
// 	var i string
// 	_, err := fmt.Scanf("%s", &i)

// 	if err != nil {
// 		log.Print("  Input Error ", err)
// 		i = ""
// 	}
// 	if len(i) == 0 {
// 		fmt.Println("You must enter a name.")
// 		return getInput()
// 	}
// 	return i
// }
