package main

import . "fmt"

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

type pickUp interface {
	pickUp(i *item)
}

type give interface {
	give(i *item)
}

func (c character) String() string {
	if c.leftHand == nil {
		return Sprintf("%v ничего не держит в руке", c.name)
	}
	return Sprintf("%v взял %v", c.name, c.leftHand.name)
}

func (c *character) pickUp(i *item) {
	if c == nil || i == nil {
		return
	}
	Printf("%v поднял %v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(k *character) {
	if c == nil || k == nil {
		return
	}
	if c.leftHand == nil {
		Printf("%v руки пусты\n", c.name)
	}
	if k.leftHand != nil {
		Printf("%v руки заняты\n", k.name)
	}
	k.leftHand = c.leftHand
	c.leftHand = nil
	Printf("%v отдал %v %v\n", c.name, k.name, k.leftHand.name)
}

func makePerson() (*character, *character, *item) {
	artur := &character{name: "Артур"}
	knight := &character{name: "лыцаль"}
	sword := &item{name: "меч"}
	return artur, knight, sword
}

func move() {
	artur, knight, sword := makePerson()
	artur.pickUp(sword)
	artur.give(knight)
	Println(knight)
	Println(artur)
}

func main() {
	move()
}
