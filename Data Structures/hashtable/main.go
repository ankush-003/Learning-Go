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
}