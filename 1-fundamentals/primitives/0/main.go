package main

import (
	"fmt"
	"goscine/pkg"
)

func step0() {
	a := 10
	var b int
	b = 20
	str := "salut"
	c := 12
	pkg.ModifyInt(&a, 21)
	pkg.ModifyInt(&b, 100)
	pkg.ModifyElemt(&str, "cabane123")
	pkg.ModifyElemt(&c, 1<<4)
	fmt.Println("- step0 -> a =", a, "\n- step0 -> b =", b, "\n- step0 -> str =", str, "\n- step0 -> c =", c)
}

func main() {
	fmt.Println("ex 0")
	step0()
	fmt.Println()
}
