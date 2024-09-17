package main

import "fmt"

func mapExample() {
	// map 자료구조 이용 예시
	var m = make(map[string]int)
	m["dog"] = 4
	m["bird"] = 2
	m["man"] = 3
	m["circle"] = 0
	fmt.Println(len(m))

	for key, value := range m {
		fmt.Println(key, value)
	}

	value, ok := m["circle"]
	fmt.Println(value, ok)

	arr := make([]int, 3, 9)
	fmt.Println(len(arr))
	fmt.Println(arr)
}
