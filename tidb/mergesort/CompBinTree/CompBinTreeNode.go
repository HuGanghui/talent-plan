package CompBinTree

type CompBinTreeNode struct {
	T []BinTreePosition
	rank int64
	ele *Entry
}

func NewCompBinTreeNode(T []BinTreePosition, entry *Entry) BinTreePosition {
	node := &CompBinTreeNode{
		T:    T,
		rank: int64(len(T)),
		ele:  entry,
	}
	T = append(T, node)
	return node
}

func (node *CompBinTreeNode) hasParent() bool {
	return 0 != node.rank
}

func (node *CompBinTreeNode) getParent() BinTreePosition {
	if node.hasParent() != false {
		return node.T[(node.rank -1) / 2]
	} else {
		return nil
	}
}

func (node *CompBinTreeNode) hasLChild() bool {
	var result bool;
	if node.rank*2 + 1 < int64(len(node.T)) {
		result = true
	} else {
		result = false
	}
	return result
}

func (node *CompBinTreeNode) hasRChild() bool {
	var result bool;
	if node.rank*2 + 2 < int64(len(node.T)) {
		result = true
	} else {
		result = false
	}
	return result
}

func (node *CompBinTreeNode) getLChild() BinTreePosition {
	var lChild BinTreePosition
	if node.hasLChild() {
		lChild = node.T[node.rank*2 + 1]
	} else {
		lChild = nil
	}
	return lChild
}

func (node *CompBinTreeNode) getRChild() BinTreePosition {
	var rChild BinTreePosition
	if node.hasRChild() {
		rChild = node.T[node.rank*2 + 2]
	} else {
		rChild = nil
	}
	return rChild
}

func (node *CompBinTreeNode) getElem() *Entry {
	return node.ele
}

func (node *CompBinTreeNode) setElem(entry *Entry) {
	node.ele = entry
}
