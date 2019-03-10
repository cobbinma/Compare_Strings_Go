package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	comp1 := comparison{args[0], args[1]}
	comp1.printComparison()
}
