package main

import (
	"fmt"
	"goscine/pkg"
)

func step3(sliceA []int, myMap map[string]int) {
	pkg.AppendSlice(sliceA)
	pkg.ModifyElemt(&sliceA[2], 42)
	fmt.Println("sliceA:", sliceA, "\nmyMap:", myMap)
}

func main() {
	sliceA := []int{1, 2, 3}
	myMap := map[string]int{"old_key": 1}

	fmt.Println("ex 3")
	step3(sliceA, myMap)
	fmt.Println()
}
