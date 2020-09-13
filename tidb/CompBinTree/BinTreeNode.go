package CompBinTree

type BinTreeNode struct {
	elem *Entry
	parent BinTreePosition
	lChild BinTreePosition
	rChild BinTreePosition
}

func NewBinTreeNode(elem *Entry, parent BinTreePosition,
	lChild BinTreePosition, rChild BinTreePosition) BinTreePosition {
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

func (node *BinTreeNode) getParent() BinTreePosition {
	return node.parent
}

func (node *BinTreeNode) getLChild() BinTreePosition {
	return node.lChild
}

func (node *BinTreeNode) getRChild() BinTreePosition {
	return node.rChild
}

func (node *BinTreeNode) getElem() *Entry {
	return node.elem
}

func (node *BinTreeNode) setElem(entry *Entry) {
	node.elem = entry
}
