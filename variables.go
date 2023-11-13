package main

import "fmt"

func main() {

	// strings
	var nameOne string = "mario"
	fmt.Println(nameOne)

	var nameTwo = "luigi"
	fmt.Println(nameTwo)

	// type inference (cannot be used outside fn)
	nameThree := "Test"
	fmt.Printf("Type of variable: %T\n", nameThree)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory: int8, int16, int32, int64
	var numOne int8 = 127
	fmt.Println(numOne)
	
	// unsigned int
	// var numTwo uint = 25

	// default float64
	var scoreOne float32 = 69.69
	scoreTwo := 25.25
	fmt.Printf("Type of variable: %T\n", scoreTwo)
	fmt.Println(scoreOne) 

	// fmt package
	fmt.Print("Hello, ")
	fmt.Println("World!")

	fmt.Println("my age is", ageOne, "my name is", nameOne)

	// Printf (formatted strings) , %_ = format specifier -> %v for variables, %q for strings
	fmt.Printf("my age is %v and my name is %v \n", ageOne, nameOne)
	fmt.Printf("my age is %q \n", ageOne)
	fmt.Printf("ageOne is of type: %T\n", ageOne)

	// float -> %f
	fmt.Printf("you scored %0.1f points \n", 225.55)
	
	// Sprintf (save formatted strings)
	str := fmt.Sprintf("My age is %v \n", ageOne)
	fmt.Println(str)
}
