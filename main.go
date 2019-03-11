package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	isError, errors := errorChecks(args)
	if isError {
		fmt.Println(errors)
	} else {
		comp1 := comparison{args[0], args[1],
			createByteArray(args[0]), createByteArray(args[1])}
		matchArray := comp1.compareByteArrays()
		if len(matchArray) > 0 {
			fmt.Println(matchArray)
		} else {
			fmt.Println("No matches found")
		}
	}
}
