package main

import "fmt"

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Printf("m1 = %+v\n", m1)

	m2 := map[string]int{}
	// m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Printf("m2 = %+v\n", m2)
}
