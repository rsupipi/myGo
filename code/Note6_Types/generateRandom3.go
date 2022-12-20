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
		generatedValues = append(generatedValues, RandomString(4))
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

func RandomString(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
