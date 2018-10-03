package main

// BubbleSort sorts a slice of ints using the bubble sort algorythm.
func BubbleSort(input []int) []int {
	swaped := true
	for swaped {
		// fmt.Println("")
		swaped = false
		for i := 0; i < len(input)-1; i++ {
			first := input[i]
			second := input[i+1]
			// fmt.Printf("Comparing %d and %d\n", first, second)
			if first > second {
				// fmt.Printf("\tswapping 2 values: %d and %d\n", first, second)
				swaped = true
				input[i] = second
				input[i+1] = first
			}
		}
	}
	return input
}
