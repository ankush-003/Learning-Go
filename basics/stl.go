package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	greeting := "hello there friends!"

	// all functions return new strings
	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("Original string:", greeting)

	ages := []int{45, 20, 54, 30, 75, 60, 50}
	
	// alters the original
	sort.Ints(ages)
	fmt.Println("ages after sorting",ages)
	index := sort.SearchInts(ages, 69)
	fmt.Println("Position of 69 is", index)
	// similarly sort.Strings, sort.SearchStrings
}
