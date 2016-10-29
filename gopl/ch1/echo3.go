// Echo3 prints its command-line arguments
package ch1

import (
	"fmt"
	"os"
	"strings"
)

func Echo3() {
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
