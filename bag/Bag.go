package Bag

import (
	// "fmt"
	"github.com/dilyara4949/goItems/item"
)

// type Bag struct {
// 	name string
// 	color string
// 	capacity float64  //gram
// 	items []item.Item
// }

// func (b *Bag) Init(name string, color string, capacity float64) {
// 	b.name = name
// 	b.color = color
// 	b.capacity = capacity
// }

// func (b *Bag) GetInfo() {
// 	fmt.Printf("%+v\n", b)
// }

// func (b *Bag) GetName() string {
// 	return b.name
// }
// func (b *Bag) SetName(name string) {
// 	b.name = name
// }
// func (b *Bag) GetColor() string {
// 	return b.color
// }
// func (b *Bag) SetColor(color string) {
// 	b.color = color
// }
// func (b *Bag) GetCapacity() float64{
// 	return b.capacity
// }
// func (b *Bag) SetCapacity(capacity float64) {
// 	b.capacity = capacity
// }

// func (b *Bag) GetItems() []item.Item {
// 	return b.items
// }

// func HaveSpace(b *Bag, newItemWeight int) bool {
// 	sum:=0.0
// 	sum += float64(newItemWeight)
// 	for i:=0; i< len(b.items); i++ {
// 		sum += float64(b.items[i].GetWeight())
// 		if (sum > b.capacity){
// 			return false
// 		}
// 	}
// 	return true
// }

// func (b *Bag) AddItem(i item.Item) {
// 	if (HaveSpace(b, i.GetWeight())) {
// 		b.items = append(b.items, i)
// 	} else {
// 		fmt.Println(b.name, "has no space")
// 	}
// }


