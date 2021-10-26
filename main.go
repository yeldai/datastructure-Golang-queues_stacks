package main

import "fmt"

type Queue struct{
	items []int
}


//Enqueue

func (q *Queue) Enqueue(i int){
	q.items= append(q.items, i)
}


//Dequeue

func (q *Queue) Dequeue()int{
	toRemove:= q.items[0]
	q.items = q.items[1:]
	return toRemove
}
type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)

}

func (s *Stack) Pop() int{
	l:= len(s.items)-1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {

	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(299)
	myStack.Push(1299)
	myStack.Push(3)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

	myQueue:=Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(11)
	myQueue.Enqueue(23)
	myQueue.Enqueue(3)
	myQueue.Enqueue(33)
	fmt.Printf("%T", myQueue)
	fmt.Println(myQueue)

	myDequeue:=myQueue
	fmt.Println(myDequeue)
	myDequeue.Dequeue()
	fmt.Printf("%T", myDequeue)
	fmt.Println(myDequeue)


}