package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	numberCount := 1000
	var generatedValues []string

	/* Add random values to a slice */
	for x := 0; x < numberCount; x++ {
		randomChar1 := 'A' + rune(rand.Intn(26))
		randomChar2 := 'A' + rune(rand.Intn(26))
		randomChar3 := 'A' + rune(rand.Intn(26))
		randomChar4 := 'A' + rune(rand.Intn(26))
		randomString := string(randomChar1) + string(randomChar2) + string(randomChar3) + string(randomChar4)
		generatedValues = append(generatedValues, randomString)
	}

	fmt.Println(generatedValues)

	/* search duplicates */
	uniqueValues := make(map[string]bool)
	duplicateValues := make(map[string]int)

	for _, item := range generatedValues {
		isNotUnique := uniqueValues[item]
		if isNotUnique {
			_, isPresent := duplicateValues[item]
			if isPresent {
				duplicateValues[item] = duplicateValues[item] + 1
			} else {
				duplicateValues[item] = 1
			}
		} else {
			uniqueValues[item] = true
		}
	}
	//fmt.Println(uniqueValues)
	fmt.Println(duplicateValues)
	fmt.Println(len(duplicateValues))
}
