package main

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

func (comp comparison) compareByteArrays() []int {
	var s []int
	for i := 0; i < len(comp.textByteArray); i++ {
		for j := 0; j < len(comp.subtextByteArray); j++ {
			if comp.textByteArray[i+j] == comp.subtextByteArray[j] {
				if j == len(comp.subtextByteArray)-1 {
					s = append(s, i+1)
				}
			} else {
				break
			}
		}
	}
	return s
}
