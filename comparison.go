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

func areTwoByteArraysTheSame(first []byte, second []byte) bool {
	for i, s := range first {
		if s != second[i] {
			return false
		}
	}
	return true
}

func (comp comparison) compareByteArrays() {
	for i := 0; i < len(comp.textByteArray); i++ {
		for j := 0; j < len(comp.subtextByteArray); j++ {
			if comp.textByteArray[i+j] == comp.subtextByteArray[j] {
				if j == len(comp.subtextByteArray)-1 {
					fmt.Println("Math Found")
				}
			} else {
				break
			}
		}
	}
}
