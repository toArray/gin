package algorithm

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	arr := []int{10, 8, 9, 1, 2, 5, 3, 4, 6, 7}

	// 冒泡
	fmt.Println(Bubbling(arr))

	// 快排
	fmt.Println(FastRow(arr, 0, len(arr)-1))
}

// Bubbling 冒泡
func Bubbling(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {

			// 判断对比两两交换
			if arr[i] > arr[j] {
				tam := arr[i]
				arr[i] = arr[j]
				arr[j] = tam
			}

		}
	}

	return arr
}

// FastRow 快排
func FastRow(arr []int, left int, right int) []int {

	if left < right {

		// 单指针
		//index := PartitionOne(arr, left, right)

		// 双指针
		index := PartitionTwo(arr, left, right)
		FastRow(arr, left, index-1)
		FastRow(arr, index+1, right)

	}

	return arr
}

// PartitionOne 单指针
func PartitionOne(arr []int, left int, right int) int {

	first := left
	index := first + 1

	for i := index; i < right; i++ {

		// 判断 n+1 小于 n   n++
		if arr[index] < arr[first] {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}

	}

	// 循环结束 n 与 n++ 数据交换
	arr[index-1], arr[first] = arr[first], arr[index-1]
	return index
}

// PartitionTwo 双指针
func PartitionTwo(arr []int, left int, right int) int {

	//基准
	tam := arr[left]
	for left < right {

		// 当尾部元素大于等于基准数据时， 向前挪动right指针
		for left < right && arr[right] >= tam {
			right--
		}

		// 如果尾部元素小于 tam 需要将其赋值给 left
		arr[left] = arr[right]

		// 当首位元素小等于基准数据时， 向前挪动 left指针
		for left < right && arr[left] <= tam {
			left++
		}

		// 当首位元素大于 tam，需要将其赋值给 right
		arr[right] = arr[left]
	}

	arr[left] = tam
	return left
}
