package structs

import "errors"

// Set
type Set struct {
	Table []*Doubly_Connected_Set
	Size  int
}

type Node_Set struct {
	Data     string
	Next     *Node_Set
	Previous *Node_Set
}

type Doubly_Connected_Set struct {
	Lenght int
	Head   *Node_Set
	Tail   *Node_Set
}

// Set
func NewSet(size int) Set {
	set := make([]*Doubly_Connected_Set, size)
	for i := range set {
		set[i] = &Doubly_Connected_Set{}
	}
	return Set{
		Table: set,
		Size:  size,
	}
}

func (ht *Set) Hash_Set(key string) int {
	key_int := 0
	for _, symbol := range key {
		key_int += int(symbol)
	}
	return key_int % ht.Size
}

func (ht *Set) Sadd(key string) error {
	hash := ht.Hash_Set(key)
	if ht.Table[hash].Lenght < 20 {
		return ht.Table[hash].dspush(key)
	} else {
		oldSize := ht.Size
		newHT := NewSet(oldSize * 2)
		for i := 0; i < oldSize; i++ {
			currentNode := ht.Table[i].Head
			for currentNode != nil {
				new_hash := newHT.Hash_Set(currentNode.Data)
				newHT.Table[new_hash].dspush(currentNode.Data)
				currentNode = currentNode.Next
			}
		}
		*ht = newHT
		new_hash := ht.Hash_Set(key)
		return ht.Table[new_hash].dspush(key)
	}
}

func (ht *Set) Srem(key string) error {
	hash := ht.Hash_Set(key)
	return ht.Table[hash].dsdel(key)
}

func (ht *Set) Sismember(key string) error {
	hash := ht.Hash_Set(key)
	currentNode := ht.Table[hash].Head
	for currentNode != nil {
		if currentNode.Data == key {
			return nil
		}
		currentNode = currentNode.Next
	}
	return errors.New("--> False")
}

func (doubly_connected *Doubly_Connected_Set) dspush(val string) error {
	node_list := &Node_Set{Data: val}
	if doubly_connected.Head == nil {
		doubly_connected.Head = node_list
		doubly_connected.Tail = node_list
	} else {
		currentNode := doubly_connected.Head
		for currentNode != nil {
			if currentNode.Data == val {
				return errors.New("--> key already exists")
			}
			currentNode = currentNode.Next
		}
		doubly_connected.Tail.Next = node_list
		node_list.Previous = doubly_connected.Tail
		doubly_connected.Tail = node_list
	}
	doubly_connected.Lenght++
	return nil
}

func (doubly_connected *Doubly_Connected_Set) dsdel(val string) error {
	currentNode := doubly_connected.Head
	for currentNode != nil {
		if currentNode.Data == val {
			if currentNode == doubly_connected.Head {
				doubly_connected.Head = currentNode.Next
				if doubly_connected.Head != nil {
					doubly_connected.Head.Previous = nil
				}
			} else if currentNode == doubly_connected.Tail {
				doubly_connected.Tail = currentNode.Previous
				if doubly_connected.Tail != nil {
					doubly_connected.Tail.Next = nil
				}
			} else {
				currentNode.Previous.Next = currentNode.Next
				currentNode.Next.Previous = currentNode.Previous
			}
			doubly_connected.Lenght--
			return nil
		}

		currentNode = currentNode.Next
	}
	return errors.New("--> key not founde")
}
