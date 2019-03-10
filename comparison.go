package main

import "fmt"

type comparison struct {
	text    string
	subtext string
}

func setComparison(text string, subtext string) comparison {
	comp := comparison{text, subtext}
	return comp
}

func (comp comparison) printComparison() {
	fmt.Println(comp.text)
	fmt.Println(comp.subtext)
}
