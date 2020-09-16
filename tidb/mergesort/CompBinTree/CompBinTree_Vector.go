package CompBinTree

type CompBinTree_Vector struct {
	T *[]BinTreePosition
}

func NewCompBinTree_Vector(entrys []*Entry) *CompBinTree_Vector {
	temp := make([]BinTreePosition, 0, len(entrys))
	var compBinTree = &CompBinTree_Vector{
		T: &temp,
	}
	if nil != entrys {
		for i := 0; i < len(entrys); i++ {
			compBinTree.AddLast(entrys[i])
		}
	}
	return compBinTree
}

func (tree *CompBinTree_Vector) GetRoot() BinTreePosition {
	return (*tree.T)[0]
}

func (tree *CompBinTree_Vector) IsEmpty() bool {
	var result bool
	if len(*tree.T) == 0 {
		result = true
	} else {
		result = false
	}
	return result
}

func (tree *CompBinTree_Vector) GetSize() int64 {
	return int64(len(*tree.T))
}

func (tree *CompBinTree_Vector) AddLast(entry *Entry) BinTreePosition {
	node := NewCompBinTreeNode(tree.T, entry)
	return node
}

func (tree *CompBinTree_Vector) DelLast() BinTreePosition {
	lastEntry := (*tree.T)[len(*tree.T)-1]
	*tree.T = (*tree.T)[:len(*tree.T)-1]
	return lastEntry
}

func (tree *CompBinTree_Vector) PosOfNode(index int64) BinTreePosition {
	return (*tree.T)[index]
}
