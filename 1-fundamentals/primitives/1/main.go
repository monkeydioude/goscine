package main

import (
	"fmt"
	"goscine/pkg"
)

func step1(sliceA []int) {
	fmt.Println("- step1 -> sliceA avant AppendSlicePtr")
	fmt.Printf("- step1 -> sliceA addr: %p\n", &sliceA)
	for i := range sliceA {
		fmt.Printf("- step1 -> sliceA[%d] addr: %p\n", i, &sliceA[i])
	}
	pkg.AppendSlicePtr(&sliceA, 4, 5, 6)
	fmt.Println("- step1 -> sliceA après AppendSlicePtr")
	fmt.Printf("- step1 -> sliceA addr: %p\n", &sliceA)
	for i := range sliceA {
		fmt.Printf("- step1 -> sliceA[%d] addr: %p\n", i, &sliceA[i])
	}
	pkg.ModifySlice(sliceA)
	fmt.Println("- step1 -> sliceA =", sliceA)
}

func main() {
	sliceA := []int{1, 2, 3}
	fmt.Println("ex 1")
	fmt.Printf("- main -> sliceA addr: %p\n", &sliceA)
	step1(sliceA)
	fmt.Println("- main -> sliceA =", sliceA)
	fmt.Println()
}
