package CompBinTree

type Entry struct{
	value *SortedSlice
	key   int64
}

func NewEntry(key int64, value *SortedSlice) *Entry {
	return &Entry{
		value: value,
		key:   key,
	}
}

func (entry *Entry) GetValue() *SortedSlice {
	return entry.value
}

func (entry *Entry) SetValue(value *SortedSlice) {
	entry.value = value
}

func (entry *Entry) GetKey() int64 {
	return entry.key
}

func (entry *Entry) SetKey(key int64) {
	entry.key = key
}
