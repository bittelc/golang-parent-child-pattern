package cat

import "fmt"

type Noisy interface {
	MakesThisNoise() string
}

type Cat struct {
	Noisy
}

func (cat Cat) SoundsLikeThis() {
	fmt.Println(cat.MakesThisNoise())
}

type Leapard struct {
	Cat
}

type HouseCat struct {
	Cat
}

func (Leapard) MakesThisNoise() string {
	return "ROOAAARRR"
}

func (HouseCat) MakesThisNoise() string {
	return "meow"
}
