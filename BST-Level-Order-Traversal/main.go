package main

import "fmt"

type Node struct {
	left *Node
	right *Node
	data int
}

type Bst struct {
	root *Node
}

func (root *Node) insert(newNode *Node) {
	if root.data == newNode.data {
		return
	}

	if root.data <= newNode.data {
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.insert(newNode)
		}
	} else {
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.insert(newNode)
		}
	}
}

func (bst *Bst) insert(data int) {
	node := &Node{data: data}

	if bst.root == nil {
		bst.root = node
	} else {
		bst.root.insert(node)
	}
}

func (bst *Bst) get_height(node *Node) int {
	if node == nil {
		return 0
	}

	max_left_depth := bst.depth_search(node.left)
	max_right_depth := bst.depth_search(node.right)
	
	var max_depth int
	if max_left_depth > max_right_depth {
		max_depth = max_left_depth
	}else {
		max_depth = max_right_depth
	}

	return max_depth
}

func (bst *Bst) depth_search(node *Node) int {
	if node == nil {
		return 0
	}

	left_count := bst.depth_search(node.left) + 1
	right_count := bst.depth_search(node.right) + 1

	var depth_count int
	if left_count > right_count {
		depth_count = left_count
	}else {
		depth_count = right_count
	}

	return depth_count
}

func (bst *Bst) levelOrder(node *Node) {
	height := bst.get_height(node)
	for i := 0; i <= height; i++ {
		bst.print_levelOrder(node, i)
	}
}

func (bst *Bst) print_levelOrder(node *Node, level int) {
	if node == nil {
		return
	}

	if level == 0 {
		fmt.Print(node.data," ")
	} else if level > 0 {
		bst.print_levelOrder(node.left, level-1)
		bst.print_levelOrder(node.right, level-1)
	}
}

func main() {
	myTree := Bst{}

	var num, data int
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		fmt.Scan(&data)
		myTree.insert(data)
	}
	myTree.levelOrder(myTree.root)
}