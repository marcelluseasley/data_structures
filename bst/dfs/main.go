package main


import (

	"fmt"
	"github.com/Workiva/go-datastructures/queue"
)

type Node struct {
	data string
	left *Node
	right *Node
}

func LevelOrder(root *Node) {
	if root == nil {
		return
	}
	Q := queue.New(10)
	Q.Put(root)
	for !Q.Empty() {
		c, _ := Q.Peek()
		current := c.(*Node)
		fmt.Println(current.data)
		if current.left != nil{
			Q.Put(current.left)
		}
		if current.right != nil {
			Q.Put(current.right)
		}
		Q.Get(1)
	}
}

func main() {

}