# Learning Go

a repository for learning Go

## Functions

A function can take zero or more arguments.

In this example, `add` takes two parameters of type `int`.

Notice that the type comes *after* the variable name.

```go
func add(x int, y int) int {
	return x + y
}
```

```go
// function can return multiple results
func swap(x, y string) (string, string) {
	return y, x
}
```

Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.

Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.

Go's basic types are

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

```go
fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
```

Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the `for` statement and the braces `{ }` are always required.

```go
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
// while loop
for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
```

```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		pic[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			pic[i][j] = uint8((j^i))
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
```

The `make` function returns a map of the given type, initialized and ready for use.

```go
m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

// map literals
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```

Insert or update an element in map `m`:

```
m[key] = elem
```

Retrieve an element:

```
elem = m[key]
```

Delete an element:

```
delete(m, key)
```

Test that a key is present with a two-value assignment:

```
elem, ok = m[key]
```

function values

```go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

[Tutorial: Developing a RESTful API with Go and Gin - The Go Programming Language](https://go.dev/doc/tutorial/web-service-gin)

## Making an API with GO

[Go API Tutorial - Make An API With Go](https://youtu.be/bj77B59nkTQ?si=Toh90Th2FOAM89Dl)

[Quickstart](https://gin-gonic.com/docs/quickstart/)

```bash
go get -u github.com/gin-gonic/gin
```

## Concurrency

Goroutines are lightweight, much lighter than a thread. Goroutines run in the same address space, so access to shared memory must be synchronized; This can be done by sync package, but it is recommended to use channels to synchronize goroutines.

A *goroutine* is a lightweight thread managed by the Go runtime.

```
go f(x, y, z)
```

```go
func randomFunc(x int) {
        fmt.Printf("Hello %v \n", x)
}

func main() {
        go randomFunc(1)
        go randomFunc(2)

        time.Sleep(2 * time.Second)

        fmt.Println("Bye!")
}
```

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

```
ch := make(chan int)
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

```go
func main() {
        myChannel := make(chan string)

        go func() {
                myChannel <- "data"
        }()
        
        // blocking
        msg := <- myChannel
        fmt.Println(msg)
}
```

The `select` statement lets a goroutine wait on multiple communication operations.

A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

```go
func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		chan1 <- "learning GO"
	} ()

	go func() {
		chan2 <- "learning concurrency in GO"
	} ()

	// if multiple at sametime then selects at random
	select {
		case msgFromChan1 := <- chan1:
			fmt.Println(msgFromChan1)
		case msgFromChan2 := <- chan2:
			fmt.Println(msgFromChan2)
	}
}
```

## References

- [A Tour of Go](https://go.dev/tour/list)

- [Concurrency Patterns in Go](https://youtu.be/qyM8Pi1KiiM?si=TgkcbIXxNpAWEnOV)

- [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)