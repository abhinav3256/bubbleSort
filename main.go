package main

import "fmt"

func main() {
	arr := []int{12, 4, 5, 1, 10, 45, 78, 15}
	resultArr := bubblesort(arr)
	fmt.Println(resultArr)

}

func bubblesort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr)-1; j++ {

			if arr[j] > arr[j+1] {

				arr = swap(arr, j, j+1)
			}

		}
	}
	return arr
}

func swap(arr []int, val1, val2 int) []int {

	temp := arr[val1]

	arr[val1] = arr[val2]
	arr[val2] = temp

	return arr

}
