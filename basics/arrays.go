package main

import "fmt"

func RemoveAtIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	// var ages [3]int = [3]int{20, 25, 30} // fixed lenth
	var ages = [3]int {20, 25, 30} // Go infers the type automatically
	
	fmt.Println(ages, len(ages))

	names := [4]string{"yoshi", "mario", "peach", "bowSER"}
	fmt.Println("Array Names: ",names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60} // bo size specified
	scores[2] = 69
	fmt.Println("Scores:",scores)
	scores = append(scores, 85) // returns a new slice
	fmt.Println("Scores:", scores)

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println("name ranges:", rangeOne, rangeTwo, rangeThree)
	rangeOne = append(rangeOne, "test")
	fmt.Println(rangeOne)

	// removing without changing order
	fmt.Println(RemoveAtIndex(scores, 1))
}
