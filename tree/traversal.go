package tree

import "fmt"

func (node *TreeNode) traverse(){
	node.TraverseFunc(func(node *TreeNode){
		node.Print()
	})
	fmt.Println()
}

func (node *TreeNode) TraverseFunc (f func(*TreeNode)){

		node.Left.TraverseFunc(f)
		f(node)
		node.Right.TraverseFunc(f)
}


func (node *TreeNode) TraverseWithChannel() chan *TreeNode {
	out := make(chan *TreeNode)
	go func() {
		node.TraverseFunc(func(node *TreeNode) {
			out <- node
		})
		close(out)
	}()
	return out
}
