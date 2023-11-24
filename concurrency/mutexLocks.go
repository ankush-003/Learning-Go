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
func processData(wg *sync.WaitGroup, result *[]int, data int) {
  // lock.Lock()
  defer wg.Done()
  processedData := process(data)
  lock.Lock()
  *result = append(*result, processedData)
  lock.Unlock()
}

func main() {
  start := time.Now()
  var wg sync.WaitGroup

  input := []int{1, 2, 3, 4, 5}

  // shared resource
  result := []int{}

  for _, data := range input {
    wg.Add(1)
    go processData(&wg, &result, data)
  }
  wg.Wait()

  log.Println(time.Since(start))
  log.Println(result)
}

