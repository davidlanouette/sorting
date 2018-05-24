package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Printf("Hello World\n")
	input := make([]int, 0, 10)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i <= 10; i++ {
		num := r.Intn(1000)
		input = append(input, num)
		fmt.Printf("%d\n", num)
	}

	fmt.Println("================")

	output := bubbleSort(input)

	fmt.Println("================")
	for _, i := range output {
		fmt.Printf("%d\n", i)
	}
}

func bubbleSort(input []int) []int {
	swaped := true
	for swaped {
		fmt.Println("")
		swaped = false
		for i := 0; i < len(input)-1; i++ {
			first := input[i]
			second := input[i+1]
			fmt.Printf("Comparing %d and %d\n", first, second)
			if first > second {
				fmt.Printf("\tswapping 2 values: %d and %d\n", first, second)
				swaped = true
				input[i] = second
				input[i+1] = first
			}
		}
	}
	return input
}
