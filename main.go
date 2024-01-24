package main

import "fmt"

func Delet[T any](s []T, idx int) []T {
	s = CoreV2[T](s, idx)
	return shrink(s)
}
func CoreV1[T any](s []T, idx int) []T {
	len := len(s)
	if idx >= len || idx < 0 {
		panic("index out of slice range")
	}
	s1 := s[:idx]
	for i := idx + 1; i < len; i++ {
		s1 = append(s1, (s)[i])
	}
	return s1
}
func CoreV2[T any](s []T, idx int) []T {
	len := len(s)
	if idx >= len || idx < 0 {
		panic("index out of slice range")
	}
	return append(s[:idx], s[idx+1:]...)
}

// 容量是长度的1.5倍则进行缩容， 缩容缩到长度的1.2倍
func shrink[T any](s []T) []T {
	c := cap(s)
	l := len(s)
	if c >= 2*l {
		s1 := make([]T, l, l+l/5)
		copy(s1, s)
		println("缩容")
		return s1
	}
	return s
}
func main() {
	s := []int{1}
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	fmt.Printf(" %d, %d\n", len(s), cap(s))
	for i := 0; i < 50; i++ {
		s = Delet(s, i+2)
	}
	fmt.Printf(" %d, %d\n", len(s), cap(s))
}
