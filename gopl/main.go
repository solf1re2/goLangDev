package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"runtime"

	"github.com/solf1re2/goLangDev/gopl/ch1"
)

const (
	GOOS string = runtime.GOOS
)

var gopathDir string

func setHomeDir() {
	gopathDir = os.Getenv("GOPATH")
	// fmt.Println(gopathDir)
	if GOOS == "windows" {
		gopathDir = gopathDir + "\\src\\github.com\\solf1re2\\goLangDev\\gopl\\"
	} else if GOOS == "linux" {
		gopathDir = gopathDir + "/src/github.com/solf1re2/goLangDev/gopl/"
	}
}

func main() {
	setHomeDir()
	// chapter := selectChapter()
	chapter := selectFromList("dir", "chapter", "")
	// program := selectProgramme(chapter)
	program := selectFromList("file", "program", chapter)
	callFunc(program)
}

func selectFromList(dirOrFile string, promptString string, extraFilePath string) string {
	//Output a list of "chapters" or "packages" for the user to choose from
	//TODO Hard code for now --> map in file or scan directories under gopl in future?
	fmt.Printf("Select %s (enter 1,2,etc)\n\n", promptString)

	var list []string
	list = readDirectoryContents(gopathDir+extraFilePath, dirOrFile)
	// log.Println(list)
	for i := 0; i < len(list); i++ {
		listItem := list[i]
		fmt.Printf("%v) %s\n", i+1, listItem)
	}

	//Wait for input from user
	listItemSelection := returnInputFromUser()

	number, err := strconv.Atoi(listItemSelection)
	if err != nil {
		// handle error
		log.Fatalln("strconv.Atoi error")
	}
	if number > len(list) {
		log.Fatalf("Selection %v not recognised.", number)
	}
	// fmt.Println("user converted input", number)
	// chapterSelection := chapterList[(number - 1)]
	selectedListItem := list[number-1]
	log.Printf("User has selected %s %s\n", promptString, selectedListItem)

	return selectedListItem
}

func callFunc(programSelection string, a ...interface{}) {
	switch programSelection {
	// ch1 programs
	case "helloworld.go":
		ch1.Helloworld()
	case "echo1.go":
		ch1.Echo1()
	case "echo2.go":
		ch1.Echo2()
	case "echo3.go":
		ch1.Echo3()
	case "dup1.go":
		ch1.Dup1()
	case "dup2.go":
		ch1.Dup2()
		// ch2 programs

	default:
		fmt.Println("Not Recognised")
	}
}

/*
Function called with input string used to prompt user for a single line input.
@promptMessage message to prompt user for input.
@returns user input
*/
func returnInputFromUser() string {
	fmt.Print("User Selection: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	firstInput := string([]byte(input)[0])
	// log.Println(firstInput)
	// fmt.Printf("Input Char Is : %v", string([]byte(input)[0]))
	return firstInput
}

func readDirectoryContents(directory string, fileOrDir string) []string {

	var list []string
	// prepare list options
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// fmt.Println(file.Name())
		f, err := os.Open(directory + "/" + file.Name())
		if err != nil {
			fmt.Println("1", err)
			return list
		}
		fi, err := f.Stat()
		if err != nil {
			fmt.Println(err)
			return list
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			// this is a chapter/package, add to the list.
			// fmt.Println("directory")
			if fileOrDir == "dir" {
				list = append(list, file.Name())
			}
		case mode.IsRegular():
			// do file stuff - ignore
			// fmt.Println("file")
			if fileOrDir == "file" {
				list = append(list, file.Name())
			}
		default:
			fmt.Println("switch case default")
		}
	}
	return list
}
