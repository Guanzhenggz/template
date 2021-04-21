package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node TreeNode) Print() {
	fmt.Println(node.Value, " ")
}

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil , Returning ...")
		return
	}
	node.Value = value
}


func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}




