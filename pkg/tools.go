package pkg

import "fmt"

func AppendSlicePtr(s *[]int, sliceToAdd ...int) {
	fmt.Printf("AppendSlicePtr avant append sliceToAdd: %v, addr: %p\n", sliceToAdd, &sliceToAdd)
	*s = append(*s, sliceToAdd...)
	fmt.Printf("AppendSlicePtr après append -> s: %v, addr: %p\n", *s, s)
	for i := range *s {
		fmt.Printf("AppendSlicePtr après append *s[%d]: %v, addr: %p\n", i, (*s)[i], &(*s)[i])
	}
}

func AppendSlice(s []int) {
	s = append(s, 4, 5, 6)
}

func ModifySlice(s []int) {
	s[0] = 999
}

func ModifyMap(m map[string]int) {
	m["new_key"] = 42
}

func ModifyInt(dst *int, value int) {
	*dst = value
}

func ModifyElemt[T any](dst *T, value T) {
	*dst = value
}
