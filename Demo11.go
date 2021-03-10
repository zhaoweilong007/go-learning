package main

import (
	"fmt"
)

func main() {

	//空接口
	var val Any
	val = 5
	fmt.Println("val:", val)
	person := new(Person)
	person.name = "张三"
	person.age = 12
	val = person
	fmt.Println("person:", person)

	switch t := val.(type) {
	case int:
		fmt.Println("type int:", t)
	case string:
		fmt.Println("type string:", t)
	case *Person:
		fmt.Println("type person:", t)
	default:
		fmt.Println("unException:", t)
	}

	root := NewNode(nil, nil)
	root.setData("root node")

	a := NewNode(nil, nil)
	b := NewNode(nil, nil)

	a.setData("left node")
	b.setData("right node")

	root.left = a
	root.right = b

	fmt.Printf("%v\n", root)

}

type Any interface {
}

type Person struct {
	name string
	age  int
}

type Node struct {
	left  *Node
	data  interface{}
	right *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left: left, right: right, data: nil}
}

func (node *Node) setData(data interface{}) {
	node.data = data
}
