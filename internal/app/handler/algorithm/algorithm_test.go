package algorithm

import (
	"fmt"
	"testing"
)

func TestBubbling(c *testing.T) {
	arr := []int{10, 8, 9, 1, 2, 5, 3, 4, 6, 7}
	fmt.Println(Bubbling(arr))
}

func TestFastRow(t *testing.T) {
	arr := []int{10, 8, 9, 1, 2, 5, 3, 4, 6, 7}
	fmt.Println(FastRow(arr, 0, len(arr)-1))
}
