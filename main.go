package main

import (
	"fmt"

	"github.com/jzes/dataStructure/linkedlist"
)

func main() {
	jose := linkedlist.Node{
		Value: "jose",
	}
	vanessa := linkedlist.Node{
		Value: "vanessa",
	}
	davi := linkedlist.Node{
		Value: "davi",
	}
	isis := linkedlist.Node{
		Value: "isis",
	}
	jose.Push(&vanessa)
	jose.Push(&davi)
	jose.Push(&isis)
	jose.Push(&linkedlist.Node{
		Value: "isabela",
	})

	fmt.Println(davi)

}
