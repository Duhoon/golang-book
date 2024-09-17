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
