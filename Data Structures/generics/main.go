package main


import (
	"fmt"
)

type CustomMap[K comparable, V any] map[K]V

func (m CustomMap[K, V]) Print() {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	newMap := make(CustomMap[string, int])
	newMap["Ankush"] = 100
	newMap["Ganya"] = 200
	newMap.Print()	
}