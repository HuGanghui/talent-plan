package CompBinTree

type CompBinTree interface {
	addLast(entry Entry) BinTreePosition
	delLast()
	posOfNode(i int64) BinTreePosition
}
