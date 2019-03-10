package main

import "os"

func main() {
	args := os.Args[1:]
	textArray := createByteArray(args[0])
	subTextArray := createByteArray(args[1])
	comp1 := comparison{args[0], args[1], textArray, subTextArray}
	comp1.compareByteArrays()
}
