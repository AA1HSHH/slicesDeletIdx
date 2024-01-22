package main

import "fmt"

func sliceDeletIdx[T any](s *[]T, idx int) {
	sliceDeletIdxV2(s, idx)
}
func sliceDeletIdxV1[T any](s *[]T, idx int) {
	len := len(*s)
	if idx >= len || idx < 0 {
		panic("index out of slice range")
	}
	s1 := (*s)[:idx]
	for i := idx + 1; i < len; i++ {
		s1 = append(s1, (*s)[i])
	}
	*s = s1
}
func sliceDeletIdxV2[T any](s *[]T, idx int) {
	len := len(*s)
	if idx >= len || idx < 0 {
		panic("index out of slice range")
	}
	*s = append((*s)[:idx], (*s)[idx+1:]...)
}
func main() {
	s := []int{1, 2, 3}
	sliceDeletIdx[int](&s, 2)
	fmt.Println(s)
}
