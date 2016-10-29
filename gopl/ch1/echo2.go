// Echo2 prints its command-line arguments
package ch1

import (
	"fmt"
	"os"
	"strconv"
)

func Echo2() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		index := strconv.Itoa(i)
		s += sep + index + ") " + arg
		sep = " \n"
	}
	fmt.Println(s)
}
