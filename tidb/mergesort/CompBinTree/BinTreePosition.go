package CompBinTree

type BinTreePosition interface {
	hasParent() bool
	getParent() BinTreePosition
	getLChild() BinTreePosition
	getRChild() BinTreePosition
	getElem() *Entry
	setElem(*Entry)
}
