package structs

import "errors"

type Node struct {
	Data string
	Next *Node
}

type Stack struct {
	Head *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

// Stack
func (stack *Stack) Spush(val string) error {
	if val == "" {
		return errors.New("-->unknown command")
	}
	node := &Node{Data: val}
	if stack.Head == nil {
		stack.Head = node
	} else {
		node.Next = stack.Head
		stack.Head = node
	}
	return nil
}

func (stack *Stack) Spop() (string, error) {
	if stack.Head == nil {
		return "", errors.New("--> stack is empty")
	} else {
		val := stack.Head.Data
		stack.Head = stack.Head.Next
		return val, nil
	}
}

// Queue
func (queue *Queue) Qpush(val string) error {
	if val == "" {
		return errors.New("-->unknown command")
	}
	node := &Node{Data: val}
	if queue.Head == nil {
		queue.Head = node
		queue.Tail = node
	} else {
		queue.Tail.Next = node
		queue.Tail = node
	}
	return nil
}

func (queue *Queue) Qpop() (string, error) {
	if queue.Head == nil {
		return "", errors.New("--> queue is empty")
	} else {
		val := queue.Head.Data
		queue.Head = queue.Head.Next
		return val, nil
	}
}
