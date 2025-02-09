package main

import(
	"fmt"

	"learningsq/stack"
	"learningsq/queue"
)

func main() {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for {
		item, ok := s.Pop()
		if !ok {
			break
		}
		fmt.Println(item)
	}

	q := queue.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	for {
		item, ok := q.Dequeue()
		if !ok {
			break
		}
		fmt.Println(item)
	}
}