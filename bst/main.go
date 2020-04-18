package main

import (
	"fmt"
)

type BstNode struct {
	data int
	left *BstNode
	right *BstNode
}

func GetNewNode(data int) *BstNode {
	newNode := &BstNode{data:data, left: nil, right: nil}
	return newNode
}

func Insert(root *BstNode, data int) *BstNode {
	if root == nil {
		root = GetNewNode(data)

	} else if data <= root.data {
		root.left = Insert(root.left, data)
	} else {
		root.right= Insert(root.right, data)
	}
	return root
}

func Delete(root *BstNode, data int) *BstNode {
	if root == nil {
		return root
	} else if data < root.data {
		root.left = Delete(root.left, data)
	} else if data > root.data {
		root.right = Delete(root.right, data)
	} else {
		//case 1: no child
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left == nil { // case 2: one child
			root = root.right
		} else if root.right == nil { // case 2: one child
			root = root.left
		} else{// case 3: two children
			temp := FindMinNode(root.right)
			root.data = temp.data
			root.right = Delete(root.right, temp.data)
		}



	}
	return root
}

func Search(root *BstNode, data int) bool {
	if root == nil {
		return false
	}
	if data == root.data {
		return true
	} else if data < root.data {
		return Search(root.left, data)
	} else {
		return Search(root.right, data)
	}
}

func FindMin(root *BstNode) int {
	if root == nil {
		return -1
	}
	if root.left == nil {
		return root.data
	}
	return FindMin(root.left)
}

func FindMax(root *BstNode) int {
	if root == nil {
		return -1
	}
	if root.right == nil {
		return root.data
	}
	return FindMax(root.right)
}

func FindMinNode(root *BstNode) *BstNode {
	if root == nil {
		return nil
	}
	if root.left == nil {
		return root
	}
	return FindMinNode(root.left)
}

func FindMaxNode(root *BstNode) *BstNode {
	if root == nil {
		return nil
	}
	if root.right == nil {
		return root
	}
	return FindMaxNode(root.right)
}

func FindHeight(root *BstNode) int {
	if root == nil {
		return -1
	}
	leftHeight := FindHeight(root.left)
	rightHeight := FindHeight(root.right)
	return max(leftHeight, rightHeight + 1)

}

func max(a, b int) int {
	if a>b {
		return a
	} else {
		return b
	}
}

func Inorder(root *BstNode) {
	if root == nil {
		return
	}
	Inorder(root.left)
	fmt.Println(root.data)
	Inorder(root.right)
}

func main() {
	var root *BstNode = nil
	root = Insert(root, 15)
	root = Insert(root, 10)
	root = Insert(root, 20)
	root = Insert(root, 25)
	root = Insert(root, 8)
	root = Insert(root, 12)

	root = Insert(root, 17)
	root = Insert(root, 6)
	root = Insert(root, 11)
	root = Insert(root, 16)
	root = Insert(root, 27)

	Inorder(root)

}