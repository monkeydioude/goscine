package main

import (
	"fmt"
	"goscine/pkg"
)

func step5(sliceA []int, sliceB []int, myMap map[string]int) {
	pkg.ModifyElemt(&sliceB, append(sliceA, 4, 5, 6))
	fmt.Println("- sliceA =", sliceA, "\n- sliceB =", sliceB, "\n- myMap =", myMap)
}

func main() {
	sliceA := []int{1, 2, 3}
	sliceB := sliceA
	myMap := map[string]int{"old_key": 1}

	fmt.Println("ex 5")
	step5(sliceA, sliceB, myMap)
	fmt.Println()
}
