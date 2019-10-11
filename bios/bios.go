package bios

import (
	"fmt"

	"github.com/iancoleman/orderedmap"
)

type Bios struct {
	*orderedmap.OrderedMap
}

type Item struct {
	Offset uint16
	Table  interface{}
}

func New() *Bios {
	om := orderedmap.New()
	return &Bios{
		OrderedMap: om,
	}
}

func (b Bios) Init() {
	b.OrderedMap = orderedmap.New()
}

func (b Bios) SetO(object interface{}, offset uint16) bool {
	return b.SetItem(Item{Offset: offset, Table: object})
}

func (b Bios) SetItem(item Item) (r bool) {
	key := fmt.Sprintf("%T", item.Table)
	b.Set(key, item)
	if val, ok := b.Get(key); ok {
		r = val.(Item).Offset == item.Offset
	}
	return
}

func (b Bios) GetItem(object interface{}) (Item, bool) {
	key := fmt.Sprintf("%T", object)
	val, ok := b.Get(key)
	return val.(Item), ok
}
