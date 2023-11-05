package structs

import "errors"

type Node_List struct {
	Data string
	Next *Node
}

type One_List struct {
	Head *Node
}

func (one_List *One_List) Opush(val string) error {
	if val == "" {
		return errors.New("-->unknown command")
	}
	node := &Node{Data: val}
	if one_List.Head == nil {
		one_List.Head = node
	} else {
		node.Next = one_List.Head
		one_List.Head = node
	}
	return nil
}

func (one_List *One_List) Opop() (string, error) {
	if one_List.Head == nil {
		return "", errors.New("--> one_List is empty")
	} else {
		val := one_List.Head.Data
		one_List.Head = one_List.Head.Next
		return val, nil
	}
}

func (one_List *One_List) Odel(val string) (string, error) {
	if one_List.Head == nil {
		return "", errors.New("--> one_List is empty")
	} else {
		curtedNode := one_List.Head
		for curtedNode != nil {
			if curtedNode.Next.Data == val {
				if curtedNode.Next != nil {
					curtedNode.Next = curtedNode.Next.Next
				} else {
					curtedNode.Next = nil
				}
				return val, nil
			}
			curtedNode = curtedNode.Next
		}
		return "", errors.New("--> element not found")
	}
}
