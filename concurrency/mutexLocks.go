package main

import (
  "sync"
  "log"
  "time"
)

var lock sync.Mutex
var m = sync.RWMutex{}
var wg2 sync.WaitGroup

var dbData = []int{1, 2, 3, 4, 5}

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

func dbCall(i int) {
  delay := 2000
  defer wg2.Done()
  time.Sleep(time.Duration(delay) * time.Millisecond)
  save(dbData[i])
  logData()
}

func save(data int) {
  m.Lock()
  defer m.Unlock()
  dbData = append(dbData, data)
}

func logData() {
  m.RLock()
  defer m.RUnlock()
  log.Println(dbData)
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

  // Mutex Locks
  for i, _ := range dbData {
    wg2.Add(1)
    go dbCall(i)
  }
  wg2.Wait()
  
}

