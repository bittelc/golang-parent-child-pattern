package dog

import "fmt"

type Dog struct{}

func (d Dog) SoundsLikeThis() {
	fmt.Println("ruff")
}
