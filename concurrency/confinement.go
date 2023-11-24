package main

import (
  "sync"
  "log"
  "time"
)

var lock sync.Mutex

func process(data int) int {
  time.Sleep(time.Second * 2)
  return data * 2
}

// critical section
func processData(wg *sync.WaitGroup, result *int, data int) {
  defer wg.Done()
  processedData := process(data)
  *result = processedData
}

func main() {
  start := time.Now()
  var wg sync.WaitGroup

  input := []int{1, 2, 3, 4, 5}

  // shared resource
  result := make([]int, len(input))

  for i, data := range input {
    wg.Add(1)
    go processData(&wg, &result[i], data)
  }
  wg.Wait()

  log.Println(time.Since(start))
  log.Println(result)
}
