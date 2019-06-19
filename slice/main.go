package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明1 slice
	var s1 []int
	fmt.Printf("s1 = %+v\n", s1)
	fmt.Printf("s1 = %+v, s1.type = %T, s1.addr = %p, s1=%+v\n", s1, s1, s1, reflect.TypeOf(s1).String())

	a := [10]int{}
	fmt.Printf("a = %+v\n", a)

	s2 := a[9]
	fmt.Printf("s2 = %+v\n", s2)

	b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s3 := b[5:10]
	fmt.Printf("s3 = %+v\n", s3)

	s4 := b[5:len(b)]
	fmt.Printf("s4 = %+v\n", s4)

	s5 := b[5:]
	fmt.Printf("s5 = %+v\n", s5)

	// 另一种 slice 声明方式
	s6 := make([]int, 3, 10)
	fmt.Printf("s6 = %+v,len=%+v,cap=%+v\n", s6, len(s6), cap(s6))

	s7 := []int{1, 2, 3, 4, 5}
	s8 := s7[:]
	fmt.Printf("s8 = %+v\n", s8)
}
