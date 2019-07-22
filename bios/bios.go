package bios

import (
	"fmt"

	"github.com/iancoleman/orderedmap"
)

type Bios struct {
	*orderedmap.OrderedMap
	Items *orderedmap.OrderedMap
}

type Item struct {
	Offset uint16
	Table  interface{}
}

func New() *Bios {
	om := orderedmap.New()
	return &Bios{
		OrderedMap: om,
		Items:      om,
	}
}

func (b *Bios) Make() {
	(*b).OrderedMap = orderedmap.New()
}

func (b *Bios) SetO(object interface{}, offset uint16) bool {
	return b.SetItem(Item{Offset: offset, Table: object})
}

func (b *Bios) SetItem(item Item) bool {
	key := fmt.Sprintf("%T", item.Table)
	(*b).Set(key, item)
	if val, ok := (*b).Get(key); ok {
		return val.(Item).Offset == item.Offset
	}
	return false
}

func (b Bios) SetItem2(item Item) bool {
	key := fmt.Sprintf("%T", item.Table)
	b.Items.Set(key, item)
	if val, ok := b.Items.Get(key); ok {
		return val.(Item).Offset == item.Offset
	}
	return false
}

func (b *Bios) GetItem(object interface{}) (Item, bool) {
	key := fmt.Sprintf("%T", object)
	val, ok := (*b).Get(key)
	return val.(Item), ok
}
