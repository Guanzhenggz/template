package queue

// An FIFQ queue.
type Queue []interface{}

// Pushes the element into the queue.
func (q *Queue) Push(v int) {
	*q = append(*q,v)
}

// Pops element from head.
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

//Returns if the queue is Empty or not
func (q *Queue) IsEmpty() bool{
	return len(*q) == 0
}



