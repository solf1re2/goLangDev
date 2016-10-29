// Echo1 prints its command-line arguments
package ch1

import (
	"fmt"
	"os"
	"strconv"
)

func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		index := strconv.Itoa(i)
		s += sep + index + ") " + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
