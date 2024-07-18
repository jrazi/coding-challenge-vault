package main

import "fmt"

func main() {
	// test
	nums, k := []int{4, 8, 19, 48, 2, 4, 0, -32, 64}, 15
	rotate(nums, k)

	fmt.Printf("%v\n", nums)
}

func rotate(nums []int, k int) {
	n := len(nums)
	rotate_s := k % n

	var rotated_arr []int = make([]int, n)

	curr := rotate_s
	for i, val := range nums {
		ind := (curr + i) % n
		rotated_arr[ind] = val
	}

	for i, val := range rotated_arr {
		nums[i] = val
	}
}
