package CompBinTree

type CompBinTreeNode struct {
	array *[]BinTreePosition
	rank int64
	ele *Entry
}

func NewCompBinTreeNode(T *[]BinTreePosition, entry *Entry) *CompBinTreeNode {
	node := &CompBinTreeNode{
		array:    T,
		rank: int64(len(*T)),
		ele:  entry,
	}
	*T = append(*T, node)
	return node
}

func (node *CompBinTreeNode) HasParent() bool {
	return 0 != node.rank
}

func (node *CompBinTreeNode) GetParent() BinTreePosition {
	if node.HasParent() != false {
		return (*node.array)[(node.rank -1) / 2]
	} else {
		return nil
	}
}

func (node *CompBinTreeNode) HasLChild() bool {
	var result bool;
	if node.rank*2 + 1 < int64(len(*node.array)) {
		result = true
	} else {
		result = false
	}
	return result
}

func (node *CompBinTreeNode) HasRChild() bool {
	var result bool;
	if node.rank*2 + 2 < int64(len(*node.array)) {
		result = true
	} else {
		result = false
	}
	return result
}

func (node *CompBinTreeNode) GetLChild() BinTreePosition {
	var lChild BinTreePosition
	if node.HasLChild() {
		lChild = (*node.array)[node.rank*2 + 1]
	} else {
		lChild = nil
	}
	return lChild
}

func (node *CompBinTreeNode) GetRChild() BinTreePosition {
	var rChild BinTreePosition
	if node.HasRChild() {
		rChild = (*node.array)[node.rank*2 + 2]
	} else {
		rChild = nil
	}
	return rChild
}

func (node *CompBinTreeNode) GetElem() *Entry {
	return node.ele
}

func (node *CompBinTreeNode) SetElem(entry *Entry) {
	node.ele = entry
}
