package main

import "gin/tree"

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = tree.CreateNode(5)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Right.Right.SetValue(5)
	root.Traverse()
	myTree := myTreeNode{&root}
	myTree.postOrder()
}
