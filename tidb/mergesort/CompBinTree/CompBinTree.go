package CompBinTree

type CompBinTree interface {
	AddLast(entry *Entry) BinTreePosition
	DelLast() BinTreePosition
	PosOfNode(i int64) BinTreePosition
}
