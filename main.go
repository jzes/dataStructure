package main

import (
	"fmt"

	"github.com/jzes/dataStructure/linkedlist"
)

func main() {
	root := linkedlist.Node{
		Value: "root",
	}
	values := [5]string{
		"first",
		"second",
		"third",
		"forth",
	}
	for _, v := range values {
		root.Push(&linkedlist.Node{
			Value: v,
		})
	}

	fmt.Println(root)
	found := root.GetByValue("forth")
	fmt.Println(found)
}
