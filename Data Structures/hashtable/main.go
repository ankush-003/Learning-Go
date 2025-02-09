package main

import (
	"fmt"
	"learninghashtable/hashtable"
)

func main() {
	ht := hashtable.NewHashTable[string, int](hashtable.WithSize[string, int](32))
	// fmt.Println(h)
	ht.Insert("Ankush", 100)
	ht.Insert("Ganya", 200)
	fmt.Println(ht.Get("Ankush"))
	fmt.Println(ht.Get("Ganya"))

	a := []int{1, 2, 3, 4, 5}
	a = append(a, 6)
	fmt.Println(a)
}