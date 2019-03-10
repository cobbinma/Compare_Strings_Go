package main

import "fmt"

type comparison struct {
	text             string
	subtext          string
	textByteArray    []byte
	subtextByteArray []byte
}

func createByteArray(givenString string) []byte {
	byteArray := []byte(givenString)
	return byteArray
}

func setComparison(text string, subtext string, textByteArray []byte,
	subtextByteArray []byte) comparison {
	comp := comparison{text, subtext, textByteArray, subtextByteArray}
	return comp
}

func (comp comparison) printComparison() {
	fmt.Println(comp.text)
	fmt.Println(comp.subtext)
}
