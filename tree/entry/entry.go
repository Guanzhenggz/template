package main

import (
	"awesomeProject/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.TreeNode
}

func (myNode *myTreeNode) postOrder() {

	if myNode == nil && myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myTreeNode{myNode.node.Right}.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.TreeNode

	root =  tree.TreeNode{Value:3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5,nil,nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6,nil,&root},
	}

	root.Left.SetValue(3)
	root.Left.Print()

	fmt.Println(nodes)
	root.Print()
	fmt.Println(root)

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(100)
	pRoot.Print()

	var sNode *tree.TreeNode
	sNode.SetValue(800)
	fmt.Println()
	rootNode := myTreeNode{&root}
	rootNode.postOrder()
	fmt.Println()

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value",maxNode)
}
