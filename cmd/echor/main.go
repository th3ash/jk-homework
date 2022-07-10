package main

import (
	"fmt"
	"os"

	"github.com/th3ash/jk-homework/stringutil"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(stringutil.Reverse(os.Args[1]))
	} else {
		fmt.Println() // print empty string
	}
}
