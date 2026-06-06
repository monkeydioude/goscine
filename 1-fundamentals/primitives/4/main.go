package main

import (
	"fmt"
)

func step4(sliceA []int, sliceB []int, myMap map[string]int) {
	sliceB = sliceA
	fmt.Println("- sliceA =", sliceA, "\n- sliceB =", sliceB, "\n- myMap =", myMap)
}

func main() {
	sliceA := []int{1, 2, 3}
	sliceB := sliceA
	myMap := map[string]int{"old_key": 1}

	fmt.Println("ex 4")
	step4(sliceA, sliceB, myMap)
	fmt.Println()
}
