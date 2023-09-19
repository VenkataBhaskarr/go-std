package tree

import "math"

type Node struct {
	Data   int
	height int
	Left   *Node
	Right  *Node
}

type Tree struct {
	Root *Node
	size int
}

func Newbst() *Tree {
	mainRoot := &Tree{
		Root: &Node{
			height: 1,
		},
		size: 0,
	}
	return mainRoot
}

func (this *Tree) Add(data int) {
	if this.size == 0 {
		this.Root.Data = data
		this.size += 1
		return
	}
	// logic to handle Data
	this.Root = this.handlingInsertion(this.Root, data)
	this.size += 1
}

func (this *Tree) handlingInsertion(temp *Node, data int) *Node {
	if temp == nil {
		return &Node{
			Data:   data,
			height: 1,
		}
	}
	if data > temp.Data {
		temp.Right = this.handlingInsertion(temp.Right, data)
	} else if data < temp.Data {
		temp.Left = this.handlingInsertion(temp.Left, data)
	} else {
		panic("Data points matched could not construct a tree")
	}
	temp.height = 1 + int(math.Max(float64(getHeight(temp.Left)), float64(getHeight(temp.Right))))
	balanceFactor := getHeight(temp.Left) - getHeight(temp.Right)
	if balanceFactor > 1 && data < temp.Left.Data {
		// Left Left
		return rotateRight(temp)
	}
	if balanceFactor > 1 && data > temp.Left.Data {
		// Left Right
		temp.Left = rotateLeft(temp.Left)
		return rotateRight(temp)

	}
	if balanceFactor < -1 && data > temp.Right.Data {
		// Right Right
		return rotateLeft(temp)
	}
	if balanceFactor < -1 && data < temp.Right.Data {
		// Right Left
		temp.Right = rotateRight(temp.Right)
		return rotateLeft(temp)
	}
	return temp
}
func rotateRight(node *Node) *Node {
	temp := node.Left
	tempRight := temp.Right
	temp.Right = node
	node.Left = tempRight
	temp.height = 1 + int(math.Max(float64(getHeight(temp.Right)), float64(getHeight(temp.Left))))
	node.height = 1 + int(math.Max(float64(getHeight(node.Right)), float64(getHeight(node.Left))))
	return temp
}
func rotateLeft(node *Node) *Node {
	temp := node.Right
	tempLeft := temp.Left
	temp.Left = node
	node.Right = tempLeft
	temp.height = 1 + int(math.Max(float64(getHeight(temp.Right)), float64(getHeight(temp.Left))))
	node.height = 1 + int(math.Max(float64(getHeight(node.Right)), float64(getHeight(node.Left))))
	return temp
}
func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (this *Tree) Remove(data int) {
	deletingTree(this.Root, data)
}
func deletingTree(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if data > node.Data {
		node.Right = deletingTree(node.Right, data)
	} else if data < node.Data {
		node.Left = deletingTree(node.Left, data)
	} else {
		// found the node to be deleted
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		// Finding the node successor

		temp := node
		succ := node.Right
		for succ.Left != nil {
			temp = succ
			succ = succ.Left
		}
		node.Data = succ.Data
		if temp == node {
			temp.Right = succ.Right
		} else {
			temp.Left = succ.Right
		}
	}
	return node
}
func (this *Tree) Getroot() *Node {
	return this.Root
}

func (this *Tree) Contains(data int) bool {
	return checkContains(this.Root, data)
}
func checkContains(node *Node, data int) bool {
	if node == nil {
		return false
	}
	if node.Data == data {
		return true
	}
	if data > node.Data {
		return checkContains(node.Right, data)
	} else {
		return checkContains(node.Left, data)
	}
}
