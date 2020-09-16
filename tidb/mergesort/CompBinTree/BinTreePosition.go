package CompBinTree

type BinTreePosition interface {
	HasParent() bool
	GetParent() BinTreePosition
	HasLChild() bool
	GetLChild() BinTreePosition
	HasRChild() bool
	GetRChild() BinTreePosition
	GetElem() *Entry
	SetElem(*Entry)
}
