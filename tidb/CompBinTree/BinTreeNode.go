package CompBinTree

type BinTreeNode struct {
	elem *Entry
	parent *BinTreePositon
	lChild *BinTreePositon
	rChild *BinTreePositon
}

func NewBinTreeNode(elem *Entry, parent *BinTreePositon,
	lChild *BinTreePositon, rChild *BinTreePositon) *BinTreeNode {
	return &BinTreeNode{
		elem:   elem,
		parent: parent,
		lChild: lChild,
		rChild: rChild,
	}
}

func (node *BinTreeNode) hasParent() bool {
	return node.parent != nil
}

func (node *BinTreeNode) getParent() *BinTreePositon {
	return node.parent
}

func (node *BinTreeNode) getLchild() *BinTreePositon {
	return node.lChild
}

func (node *BinTreeNode) getRchild() *BinTreePositon {
	return node.rChild
}

func (node *BinTreeNode) getElem() *Entry {
	return node.elem
}

func (node *BinTreeNode) setElem(entry *Entry) {
	node.elem = entry
}
