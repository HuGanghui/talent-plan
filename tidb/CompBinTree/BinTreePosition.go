package CompBinTree

type BinTreePositon interface {
	hasParent() bool
	getParent() *BinTreePositon
	getLChild() *BinTreePositon
	getRChild() *BinTreePositon
	getElem() *Entry
	setElem(*Entry)
}
