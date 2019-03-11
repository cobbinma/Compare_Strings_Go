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

func TestCompareByteArrays(t *testing.T) {
	var text = "hsadkjahsdkjateghgfhgfhfsttestafasfasf"
	var subtext = "test"
	var expectedMatch = 27
	comparisonTest := setComparison(text, subtext, createByteArray(text), createByteArray(subtext))
	matchArray := comparisonTest.compareByteArrays()

	if matchArray[0] != expectedMatch {
		t.Fatalf("Expected match at %v but did there was not", expectedMatch)
	}
}
