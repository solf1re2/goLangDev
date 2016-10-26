package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var chapterList []string

func main() {
	prepareChapterList()
	selectChapter()
	// ch1.Echo1()
	// ch1.Echo2()
	// ch1.Helloworld()
	// init()
}

func prepareChapterList() {
	// prepare list options
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// fmt.Println(file.Name())
		f, err := os.Open(file.Name())
		if err != nil {
			fmt.Println(err)
			return
		}
		fi, err := f.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			// this is a chapter/package, add to the list.
			// fmt.Println("directory")
			chapterList = append(chapterList, file.Name())
		case mode.IsRegular():
			// do file stuff - ignore
			// fmt.Println("file")
		}
	}
	// fmt.Println(chapterList)
	// fmt.Println("preparing chapter list")
	log.Println("preparing chapter list")
}

func selectChapter() {
	//Output a list of "chapters" or "packages" for the user to choose from
	//TODO Hard code for now --> map in file or scan directories under gopl in future?
	fmt.Println("Select chapter (enter 1,2,etc)")
	// for _, chapter := range chapterList {
	// 	fmt.Println(chapter)
	// }
	for i := 0; i < len(chapterList); i++ {
		chapter := chapterList[i]
		fmt.Printf("%v) %s\n", i+1, chapter)
	}
	//Wait for input from user
	chapterSelection := returnInputFromUser()
	log.Printf("User has selected chapter %s\n", chapterSelection)
}

func selectProgramme() {

}

/*
Function called with input string used to prompt user for a single line input.
@promptMessage message to prompt user for input.
@returns user input
*/
func returnInputFromUser() string {
	var i string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// assign the input to i
		i = scanner.Text()

		// Check that there is input
		if len(i) > 0 {
			break
		} else {
		}
	}
	if err := scanner.Err(); err != nil {
		log.Print(os.Stderr, "reading standard input:", err)
	}
	fmt.Printf("\n")

	return i
}

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
	fmt.Printf("\n")

	return i
}
