package tree

func (node *Node) Traverse() {
	if node == nil {
		return
	}

}

func (node *Node) TravelseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
