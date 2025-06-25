package repl

import "compiler/value"

type SliceValue struct {
	InternalValue []value.IVOR
	CurrentIndex  int
	ItemType      string
	FullType      string
	SizeValue     *value.IntValue
	IsEmpty       *value.BoolValue
}

// todas las clases que hemos realizado hasta ahora tienen el Copy(), el Type()
func (s SliceValue) Value() interface{} {
	return s
}
