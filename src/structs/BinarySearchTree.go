package structs

import (
	"errors"
	"strings"
)

type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Tins(value string) {
	newNode := &TreeNode{Value: value}
	if bst.Root == nil {
		bst.Root = newNode
		return
	}
	bst.insertNode(bst.Root, newNode)
}

func (bst *BinarySearchTree) insertNode(node, newNode *TreeNode) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			bst.insertNode(node.Left, newNode)
		}
	} else if newNode.Value > node.Value {
		if node.Right == nil {
			node.Right = newNode
		} else {
			bst.insertNode(node.Right, newNode)
		}
	}
}

func (bst *BinarySearchTree) Tcon(value string) error {
	return bst.containsNode(bst.Root, value)
}

func (bst *BinarySearchTree) containsNode(node *TreeNode, value string) error {
	if node == nil {
		return errors.New("--element not found")
	}

	if value < node.Value {
		return bst.containsNode(node.Left, value)
	} else if value > node.Value {
		return bst.containsNode(node.Right, value)
	} else {
		return nil
	}
}

func (bst *BinarySearchTree) Tdel(value string) error {
	err := bst.Tcon(value)
	if err != nil {
		return err
	}
	if bst.Root.Left == nil && bst.Root.Right == nil {
		return errors.New("--> only the top remains in the tree")
	}
	bst.Root = bst.deleteNode(bst.Root, value)
	return nil
}

func (bst *BinarySearchTree) deleteNode(node *TreeNode, value string) *TreeNode {
	if node == nil {
		return node
	}

	if value < node.Value {
		node.Left = bst.deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.deleteNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		node.Value = Tmin(node.Right)
		node.Right = bst.deleteNode(node.Right, node.Value)
	}
	return node
}

func Tmin(node *TreeNode) string {
	if node == nil {
		return "" // Обработка пустого дерева
	}
	for node.Left != nil {
		node = node.Left
	}
	return node.Value
}

func Tmax(node *TreeNode) string {
	if node == nil {
		return "" // Обработка пустого дерева
	}
	for node.Right != nil {
		node = node.Right
	}
	return node.Value
}

func (bst *BinarySearchTree) InOrderTraversal(node *TreeNode, values *[]string, isRootNode bool) {
	if node != nil {
		bst.InOrderTraversal(node.Left, values, false)
		if !isRootNode {
			*values = append(*values, strings.TrimSpace(node.Value))
		}
		bst.InOrderTraversal(node.Right, values, false)
	}
}
