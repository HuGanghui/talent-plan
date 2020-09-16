package CompBinTree

type Entry struct{
	value int64
	key int64
}

func NewEntry(key int64, value int64) *Entry {
	return &Entry{
		value: value,
		key:   key,
	}
}

func (entry *Entry) GetValue() int64 {
	return entry.value
}

func (entry *Entry) SetValue(value int64) {
	entry.value = value
}

func (entry *Entry) GetKey() int64 {
	return entry.key
}

func (entry *Entry) SetKey(key int64) {
	entry.key = key
}
