package main

import (
	"fmt"
)

type Queue struct {
	slice []int
}

func (q *Queue) Enqueue(i int) {
	q.slice  = append(q.slice, i)

}

func (q *Queue) Dequeue() int {

	ret := q.slice[0]
	q.slice = q.slice[1:]
	return ret
}

func (q *Queue) String() string {
	return fmt.Sprintln(q.slice)
}

func main() {
	var q *Queue = new(Queue)
	q.Enqueue(123)
	q.Enqueue(75)
	q.Enqueue(3)
	q.Enqueue(44)
	fmt.Println(q.Dequeue())
	fmt.Println(q)

}
