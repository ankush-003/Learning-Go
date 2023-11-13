package main

import "fmt"

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
}
