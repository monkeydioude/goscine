package main

import (
	"fmt"
	"goscine/pkg"
)

func step2(sliceA []int, myMap map[string]int) {
	pkg.ModifyInt(&sliceA[1], 21)
	tmp := myMap["old_key"]
	pkg.ModifyInt(&tmp, 123)
	myMap["old_key"] = tmp
	fmt.Println("- sliceA =", sliceA, "\n- myMap =", myMap)
}

func main() {
	sliceA := []int{1, 2, 3}
	myMap := map[string]int{"old_key": 1}

	fmt.Println("ex 2")
	step2(sliceA, myMap)
	fmt.Println()
}
