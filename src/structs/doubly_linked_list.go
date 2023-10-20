package structs

import "errors"

type Node_Dl struct {
	Data     string
	Next     *Node_Dl
	Previous *Node_Dl
}

type Doubly_Linked_List struct {
	Lenght int
	Head   *Node_Dl
	Tail   *Node_Dl
}

func (dll *Doubly_Linked_List) Dlpush_end(val string) error {
	node_dl := &Node_Dl{Data: val}
	if dll.Head == nil {
		dll.Head = node_dl
		dll.Tail = node_dl
	} else {
		dll.Tail.Next = node_dl
		node_dl.Previous = dll.Tail
		dll.Tail = node_dl
	}
	dll.Lenght++
	return nil
}

func (dll *Doubly_Linked_List) Dlpush_begin(val string) error {
	node_dl := &Node_Dl{Data: val}
	if dll.Head == nil {
		dll.Head = node_dl
		dll.Tail = node_dl
	} else {
		dll.Head.Previous = node_dl
		node_dl.Next = dll.Head
		dll.Head = node_dl
	}
	dll.Lenght++
	return nil
}

func (dll *Doubly_Linked_List) Dldel_end() error {
	if dll.Tail == nil {
		return errors.New("--> list is void")
	}
	if dll.Head == dll.Tail {
		dll.Head = nil
		dll.Tail = nil
		dll.Lenght = 0
		return nil
	}
	dll.Tail = dll.Tail.Previous
	if dll.Tail != nil {
		dll.Tail.Next = nil
	}
	dll.Lenght--
	return nil
}

func (dll *Doubly_Linked_List) Dldel_begin() error {
	if dll.Head == nil {
		return errors.New("--> list is void")
	}
	dll.Head = dll.Head.Next
	if dll.Head != nil {
		dll.Head.Previous = nil
	}
	dll.Lenght--
	return nil
}

func (dll *Doubly_Linked_List) Dldel(val string) error {
	node_dl := dll.Head
	for node_dl != nil {
		if node_dl.Data == val {
			if node_dl == dll.Head {
				dll.Dldel_begin()
			} else if node_dl == dll.Tail {
				dll.Dldel_end()
			} else {
				node_dl.Previous.Next = node_dl.Next
				node_dl.Next.Previous = node_dl.Previous
			}
			dll.Lenght--
			return nil
		}
		node_dl = node_dl.Next
	}
	return errors.New("--> element not founde")
}

func (dll *Doubly_Linked_List) Dcout(val string) error {
	node_dl := dll.Head
	for node_dl != nil {
		if node_dl.Data == val {
			return nil
		}
		node_dl = node_dl.Next
	}
	return errors.New("--> element not founde")
}

func (dll *Doubly_Linked_List) Dlen() int {
	return dll.Lenght
}
