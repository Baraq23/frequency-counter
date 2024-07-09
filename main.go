//  Frequency Counter - Write a function that takes a string and returns a map with the frequency of each character.

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <string>")
		os.Exit(0)
	}

	str := os.Args[1]

	freqMap := MapFrequency(str)
	fmt.Println(freqMap)
}

// function MapFrequency() takes a string and returns a map with the frequency of each character.
func MapFrequency(s string) map[string]int {
	freqMap := make(map[string]int)

	for _, c := range s {
		freqMap[string(c)] += 1
	}

	return freqMap
}
