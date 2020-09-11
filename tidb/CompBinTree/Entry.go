package CompBinTree

type Entry struct{
	value int64
	key int64
}

func (entry *Entry) getValue() int64 {
	return entry.value
}

func (entry *Entry) setValue(value int64) {
	entry.value = value
}

func (entry *Entry) getKey() int64 {
	return entry.key
}

func (entry *Entry) setKey(key int64) {
	entry.key = key
}
