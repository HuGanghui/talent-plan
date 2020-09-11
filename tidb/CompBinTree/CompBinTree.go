package CompBinTree

type CompBinTree interface {
	addLast(entry Entry) BinTreePositon
	delLast()
	posOfNode(i int64) BinTreePositon
}

type CompBinTreeNode struct {
	T [] *CompBinTreeNode
	rank int64
	ele Entry
}

//hasParent() bool
//getParent() *BinTreePositon
//setParent(p BinTreePositon)
//getLChild() *BinTreePositon
//getRChild() *BinTreePositon
//getElem() *Entry
//setElem(*Entry)

func NewCompBinTreeNode(T []*CompBinTreeNode, entry Entry) *CompBinTreeNode{
	var node *CompBinTreeNode = &CompBinTreeNode{
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

func (node *CompBinTreeNode) getParent() *CompBinTreeNode {
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

func (node *CompBinTreeNode) getLChild() *CompBinTreeNode {
	var lChild *CompBinTreeNode
	if node.hasParent() {
		lChild = node.T[node.rank*2 + 1]
	} else {
		lChild = nil
	}
	return lChild
}

func (node *CompBinTreeNode) getRChild() *CompBinTreeNode {
	var rChild *CompBinTreeNode
	if node.hasParent() {
		rChild = node.T[node.rank*2 + 2]
	} else {
		rChild = nil
	}
	return rChild
}



