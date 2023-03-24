package main

import (
	"fmt"
	"testproject"
)

func main() {
	testproject.StartServer()

	// arr1 := []string{"a", "b", "c", "e", "f"}
	// arr2 := []string{"b", "c", "d", "e"}

	// printUniqueValues(arr1, arr2)

}

func printUniqueValues(arr1 []string, arr2 []string) {
	uniqueValues := make(map[string]int)
	var allval = []string{}
	var allval2 = []string{}

	for _, value := range arr1 {
		uniqueValues[value]++
	}

	for _, value := range arr2 {
		uniqueValues[value]++
	}

	for key, count := range uniqueValues {
		allval = append(allval, key)
		if count == 1 {
			allval2 = append(allval2, key)
		}
		fmt.Println("i", key)
	}
	fmt.Println("allval1", allval)
	fmt.Println("allval2", allval2)
}
