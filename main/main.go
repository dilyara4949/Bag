package main

import (
	"github.com/dilyara4949/goItems/item"
	"goBag/Bag"
	"fmt"
)

// bag with list of items, own characteristics

func main() {
	i1:=new(item.Item)
	i1.SetDefault()
	i1.GetInfo()
	fmt.Print("\n")

	b1 := new(Bag.Bag)
	b1.Init("shopper", "balck", 5000)
	b1.GetInfo()

	i2:=new(item.Item)
	i2.Init("phone", "desk", 2000)
	b1.AddItem(*i2)
	b1.GetInfo()

	i3:=new(item.Item)
	i3.Init("watremelon", "fridge", 4000)
	b1.AddItem(*i3)
	b1.GetInfo()
	
}

