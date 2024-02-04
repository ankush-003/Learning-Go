package main

import (
  "time"
  "context"
  "log"
  "fmt"
)

func main() {
  start := time.Now()
  ctx := context.WithValue(context.Background(), "userID", 10)
  userID := 10
  val, err := fetchUserData(ctx, userID)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("result: ", val)
  fmt.Println("time taken: ", time.Since(start))
}

type Response struct {
  value int
  err error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
  ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 200)
  defer cancel()
  
  respch := make(chan Response)

  // schedule the call
  go func() {
    val, err := fetchThirdPartyStuffWhichCanBeSlow(ctx)
    respch <- Response{
      value: val,
      err: err, 
    }
  } ()
    for {
      select {
      case <-ctx.Done():
        return 0, fmt.Errorf("fetching data from third party took too long, userID: %d", userID)
      case resp := <-respch:
        return resp.value, resp.err
      }
    }
}

func fetchThirdPartyStuffWhichCanBeSlow(ctx context.Context) (int, error) {
  time.Sleep(time.Millisecond * 500)
  val := ctx.Value("userID").(int) // type assertion in Go is done using the .(type) syntax
  return val, nil
}
