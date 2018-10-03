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

	output := BubbleSort(input)

	fmt.Println("================")
	for _, i := range output {
		fmt.Printf("%d\n", i)
	}
}
