package main

import (
	"github.com/parent_and_subclasses/cat"
	"github.com/parent_and_subclasses/dog"
)

type Animal interface {
	SoundsLikeThis()
}

func MakeSound(a Animal) {
	a.SoundsLikeThis()
}

func main() {
	leapard := &cat.Cat{&cat.Leapard{}}
	kitty := &cat.Cat{&cat.HouseCat{}}
	dog := &dog.Dog{}

	MakeSound(leapard)
	MakeSound(kitty)
	MakeSound(dog)
}
