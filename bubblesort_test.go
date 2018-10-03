package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestBubbleSortNil(t *testing.T) {
	sorter := BubbleSort(nil)
	if sorter != nil {
		t.Fail()
	}
}

func TestBubbleSortShortList(t *testing.T) {
	assert := assert.New(t)

	input := []int{5, 4, 3, 2, 1}
	output := BubbleSort(input)

	assert.Equal([]int{1, 2, 3, 4, 5}, output)
}

func TestBubbleSortRandomList(t *testing.T) {
	assert := assert.New(t)

	input := []int{2325, 2432, 8873, 34966, 4020, 5456, 42285}
	output := BubbleSort(input)

	assert.Equal([]int{2325, 2432, 4020, 5456, 8873, 34966, 42285}, output)
}
