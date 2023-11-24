package main

import (
  "sync"
  "log"
)

// critical section
func processData(wg *sync.WaitGroup, result *[]int, data int) {
  defer wg.Done()

  processedData := data * 2

  *result = append(*result, processedData)
}

func main() {
  var wg sync.WaitGroup

  input := []int{1, 2, 3, 4, 5}

  // shared resource
  result := []int{}

  for _, data := range input {
    wg.Add(1)
    go processData(&wg, &result, data)
  }
  wg.Wait()
  log.Println(result)
}
