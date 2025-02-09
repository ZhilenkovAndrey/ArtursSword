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

func (c *character) String() string {
	if c.leftHand != nil {
		Printf("%V держит в руке %V", c.name, c.leftHand.name)
	}
	if c.leftHand == nil {
		Printf("%V ничего не держит в руке", c.name)
	}
}

func (c *character) pickUp(i *item) {
	if c == nil || i == nil {
		return
	}

}
