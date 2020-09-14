package CompBinTree

type CompBinTree interface {
	addLast(ele int64) BinTreePosition
	delLast()
	posOfNode(i int64) BinTreePosition
}
