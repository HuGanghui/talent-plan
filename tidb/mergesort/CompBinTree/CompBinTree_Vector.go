package CompBinTree

type CompBinTree_Vector struct {
	T []BinTreePosition
}

func NewCompBinTree_Vector(s []int64) CompBinTree {
	var compBinTree = &CompBinTree_Vector{
		T: make([]BinTreePosition, 0),
	}
	if nil != s {
		for i := 0; i < len(s); i++ {
			compBinTree.addLast(s[i])
		}
	}
	return compBinTree
}

func (tree *CompBinTree_Vector) getRoot() BinTreePosition {
	return tree.T[0]
}

func (tree *CompBinTree_Vector) isEmpty() bool {
	var result bool
	if len(tree.T) == 0 {
		result = true
	} else {
		result = false
	}
	return result
}

func (tree *CompBinTree_Vector) getSize() int64 {
	return int64(len(tree.T))
}

func (tree *CompBinTree_Vector) addLast(ele int64) BinTreePosition {
	entry := &Entry{
		key:   ele,
		value: ele,
	}
	node := NewCompBinTreeNode(tree.T, entry)
	return node
}

func (tree *CompBinTree_Vector) delLast() {
	tree.T = tree.T[:len(tree.T)-1]
}

func (tree *CompBinTree_Vector) posOfNode(index int64) BinTreePosition {
	return tree.T[index]
}
