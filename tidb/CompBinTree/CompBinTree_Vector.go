package CompBinTree

type CompBinTree_Vector struct {
	T []BinTreePosition
}

func NewCompBinTree_Vector(s []int64) {
	var compBinTree = &CompBinTree_Vector{
		T: make([]BinTreePosition, 0),
	}
	if nil != s {
		for i := 0; i < len(s); i++ {
			compBinTree.addLast(s[i])
		}
	}
}

func (tree *CompBinTree_Vector) addLast(ele int64) BinTreePosition {
	entry := &Entry{
		key:   ele,
		value: ele,
	}
	node := NewCompBinTreeNode(tree.T, entry)
	return node
}
