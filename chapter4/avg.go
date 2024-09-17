package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func avg() float64 {
	fmt.Println("Please insert numbers.")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()

	s := reader.Text()

	total := 0.0
	length := 0.0
	for _, score := range strings.Split(s, " ") {
		scoreParsed, _ := strconv.ParseFloat(score, 64)

		total += scoreParsed
		length += 1
	}
	return total / length
}
