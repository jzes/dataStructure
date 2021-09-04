package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	Value string
	Next  *Node
}

func (n *Node) Push(nn *Node) {
	if n.Next == nil {
		n.Next = nn
		return
	}
	n.Next.Push(nn)
}

func (n *Node) GetByValue(value string) *Node {
	if n.Value == value {
		return n
	}
	if n.Next == nil {
		return nil
	}
	return n.Next.GetByValue(value)
}

func (n Node) String() string {
	var sb strings.Builder
	for n.Next != nil {
		fmt.Fprintf(&sb, "%v, ", n.Value)
		n = *n.Next
	}
	fmt.Fprintf(&sb, "%v ", n.Value)
	return sb.String()
}
