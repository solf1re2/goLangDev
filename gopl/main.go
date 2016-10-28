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
	// if GOOS == "windows" {
	gopathDir = os.Getenv("GOPATH")
	// fmt.Println(gopathDir)
	// } else {

	// }
	gopathDir = gopathDir + "\\src\\github.com\\solf1re2\\goLangDev\\gopl"
}

func main() {
	setHomeDir()
	chapter := selectChapter()
	program := selectProgramme(chapter)
	callFunc(program)
}

func selectChapter() string {
	//Output a list of "chapters" or "packages" for the user to choose from
	//TODO Hard code for now --> map in file or scan directories under gopl in future?
	fmt.Println("Select chapter (enter 1,2,etc)")
	var chapterList []string
	chapterList = readDirectoryContents(gopathDir, "dir")

	for i := 0; i < len(chapterList); i++ {
		chapter := chapterList[i]
		fmt.Printf("%v) %s\n", i+1, chapter)
	}
	//Wait for input from user
	chapterSelectionInput := returnInputFromUser()

	number, err := strconv.Atoi(chapterSelectionInput)
	if err != nil {
		// handle error
		log.Fatalln("strconv.Atoi error")
	}
	if number > len(chapterList) {
		log.Fatalf("Selection %v not recognised.", number)
	}
	// fmt.Println("user converted input", number)
	// chapterSelection := chapterList[(number - 1)]
	chapterSelection := chapterList[number-1]
	log.Printf("User has selected chapter %s\n", chapterSelection)

	return chapterSelection
}

func selectProgramme(dir string) string {
	var programList []string
	programList = readDirectoryContents(gopathDir+"\\"+dir, "file")
	for i := 0; i < len(programList); i++ {
		program := programList[i]
		fmt.Printf("%v) %s/%s\n", i+1, dir, program)
	}
	//Wait for input from user
	programSelectionInput := returnInputFromUser()
	// fmt.Println(programSelectionInput)
	// fmt.Printf("User input: %v\n", programSelectionInput)
	number, err := strconv.Atoi(programSelectionInput)
	if err != nil {
		// handle error
		log.Fatalln("strconv.Atoi error")
	}
	if number > len(programList) {
		log.Fatalf("Selection %v not recognised.", number)
	}

	programSelection := programList[number-1]
	// log.Printf("User has selected program %s\n", programSelection)
	fmt.Printf("\n%s/%s\n\n", dir, programSelection)
	// invoke("ch1.Echo1")
	return programSelection
}

func callFunc(programSelection string, a ...interface{}) {
	switch programSelection {
	case "echo1.go":
		ch1.Echo1()
	case "echo2.go":
		ch1.Echo2()
	case "helloworld.go":
		ch1.Helloworld()
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
	// var i = 0
	// var i string
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	// assign the input to i
	// 	i = scanner.Text()

	// 	// Check that there is input
	// 	if len(i) > 0 {
	// 		break
	// 	} else {
	// 	}
	// }
	// if err := scanner.Err(); err != nil {
	// 	log.Print(os.Stderr, "reading standard input:", err)
	// }
	// fmt.Printf("\n")

	// return i

	//---------------------------------
	// reader := bufio.NewReader(os.Stdin)
	// b, _ := reader.ReadByte()
	// var ba []byte
	// ba = nil
	// ba = append(ba, b)
	// _, i = binary.Varint(ba)
	// fmt.Println(i)
	// return i
	fmt.Print("User Selection: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	firstInput := string([]byte(input)[0])
	// log.Println(firstInput)
	// fmt.Printf("Input Char Is : %v", string([]byte(input)[0]))
	return firstInput
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
