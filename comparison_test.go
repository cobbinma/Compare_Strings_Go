package main

import "testing"

func TestComparisonSetComparison(t *testing.T) {
	var text = "testtesttest"
	var subtext = "text"
	comparisonTest := setComparison(text, subtext, []byte{}, []byte{})

	if comparisonTest.text != text {
		t.Fatalf("Expected %s but got %s", text, comparisonTest.text)
	}

	if comparisonTest.subtext != subtext {
		t.Fatalf("Expected %s but got %s", subtext, comparisonTest.subtext)
	}
}

func TestAreTwoByteArraysTheSame(t *testing.T) {
	isMatch := areTwoByteArraysTheSame([]byte{1, 2, 3, 4}, []byte{1, 2, 3, 4})

	if isMatch != true {
		t.Fatalf("Arrays do not match")
	}
}
