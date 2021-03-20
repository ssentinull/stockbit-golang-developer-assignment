package main

import (
	"fmt"
)

func main() {
	alphabets := generateAlphabetsMap()

	var highestNumericVal, lowestNumericVal int
	anagramGroupsMap := make(map[int][]string)
	words := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	for _, word := range words {
		wordNumericVal := calculateWordNumericValue(word, alphabets)
		anagramGroup := anagramGroupsMap[wordNumericVal]
		anagramGroup = append(anagramGroup, word)
		anagramGroupsMap[wordNumericVal] = anagramGroup

		highestNumericVal, lowestNumericVal = getHighestLowestNumericValue(highestNumericVal,
			lowestNumericVal, wordNumericVal)
	}

	anagramGroupsSlices := anagramGroupMapToSlice(highestNumericVal, lowestNumericVal,
		anagramGroupsMap)

	fmt.Println(anagramGroupsSlices)
}

func generateAlphabetsMap() map[string]int {
	var char string
	chars := make(map[string]int)
	for i := 0; i < 26; i++ {
		char = string(rune('a' + i))
		chars[char] = i
	}

	return chars
}

func calculateWordNumericValue(word string, alph map[string]int) int {
	var charNumericVal, wordNumericVal int
	for _, char := range word {
		charNumericVal = alph[string(char)]
		wordNumericVal += charNumericVal
	}

	return wordNumericVal
}

func getHighestLowestNumericValue(high, low, current int) (int, int) {
	if high == 0 || current > high {
		high = current
	}

	if low == 0 || current < low {
		low = current
	}

	return high, low
}

func anagramGroupMapToSlice(high, low int, anagramMap map[int][]string) [][]string {
	anagramSlices := make([][]string, 0)
	for i := high; i >= low; i-- {
		if anagramGroup, ok := anagramMap[i]; ok {
			anagramSlices = append(anagramSlices, anagramGroup)
		}
	}

	return anagramSlices
}
