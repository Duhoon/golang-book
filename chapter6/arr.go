package main

import "fmt"

func arrExample() {
	// x 배열 할당 및 for 문 이용한 평균 계산
	x := [6]float64{
		98,
		93,
		77,
		82,
		83,
	}

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func copyExample() {
	arr := [5]int{
		56,
		13,
		44,
		78,
		91,
	}

	slice := arr[:]

	sliceCopy := make([]int, len(arr))
	copy(sliceCopy, slice)

	fmt.Printf("arr pointer: %p\n", &arr)
	fmt.Println(arr)
	fmt.Printf("slice pointer: %p\n", &slice)
	fmt.Println(slice)
	fmt.Printf("copied slice pointer: %p\n", &sliceCopy)
	fmt.Println(sliceCopy)

	slice[0] = 1
	sliceCopy[0] = 7

	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(sliceCopy)
}
