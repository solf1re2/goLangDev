package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/solf1re2/goLangDev/gopl/ch1"
)

func main() {
	// prepareChapterList()
	chapter := selectChapter()
	program := selectProgramme(chapter)
	callFunc(program)
	// ch1.Echo1()
	// ch1.Echo2()
	// ch1.Helloworld()
	// init()
}

func prepareChapterList() {

	var chapterList []string
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

func selectChapter() string {
	//Output a list of "chapters" or "packages" for the user to choose from
	//TODO Hard code for now --> map in file or scan directories under gopl in future?
	fmt.Println("Select chapter (enter 1,2,etc)")
	// for _, chapter := range chapterList {
	// 	fmt.Println(chapter)
	// }

	var chapterList []string
	chapterList = readDirectoryContents(".", "dir")

	for i := 0; i < len(chapterList); i++ {
		chapter := chapterList[i]
		fmt.Printf("%v) %s\n", i+1, chapter)
	}
	//Wait for input from user
	chapterSelectionInput := returnInputFromUser()
	chapterSelection := chapterList[chapterSelectionInput-1]
	log.Printf("User has selected chapter %s\n", chapterSelection)

	return chapterSelection
}

func selectProgramme(dir string) string {
	var programList []string
	programList = readDirectoryContents(dir, "file")
	for i := 0; i < len(programList); i++ {
		program := programList[i]
		fmt.Printf("%v) %s\n", i+1, program)
	}
	//Wait for input from user
	programSelectionInput := returnInputFromUser()
	fmt.Println(programSelectionInput)
	fmt.Printf("User input: %v\n", programSelectionInput)
	programSelection := programList[(programSelectionInput - 1)]
	log.Printf("User has selected program %s\n", programSelection)
	fmt.Println(dir + "/" + programSelection)
	// invoke("ch1.Echo1")
	return programSelection
}

// func invoke(fn interface{}, args ...string) {
// 	v := reflect.ValueOf(fn)
// 	rargs := make([]reflect.Value, len(args))
// 	for i, a := range args {
// 		rargs[i] = reflect.ValueOf(a)
// 	}
// 	v.Call(rargs)
// }

func callFunc(programSelection string) {
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
	var i = 0
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

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// fmt.Printf("Input Char Is : %v", string([]byte(input)[0]))
	return input
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
