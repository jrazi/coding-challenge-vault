package main

import "fmt"

func main() {
	// simple test
	nums := []int{-1, 1, 0, -3, 3}
	answer := productExceptSelf(nums)
	fmt.Printf("%v\n", answer)
}

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	// fast path: more than one zero in the array
	zeroCount := 0
	lastZeroIndex := -1
	for i, val := range nums {
		if val == 0 {
			zeroCount += 1
			lastZeroIndex = i
		}
	}
	if zeroCount > 1 {
		return answer
	} else if zeroCount == 1 {
		answer[lastZeroIndex] = productExceptIndex(nums, lastZeroIndex)
		return answer
	}

	leftProds := make([]int, len(nums)+1)
	rightProds := make([]int, len(nums)+1)

	productSteps(0, 1, nums, leftProds)
	productSteps(len(rightProds)-2, -1, nums, rightProds)

	for i := 0; i < len(nums); i++ {
		answer[i] = leftProds[i] * rightProds[i]
	}

	return answer
}

func productSteps(strInd int, step int, nums []int, prods []int) {
	prods[strInd] = 1
	for i := strInd + step; i < len(prods) && i >= 0; i += step {
		prods[i] = prods[i-step] * nums[i-step]
	}
}

func productExceptIndex(nums []int, index int) int {
	product := 1
	for i, val := range nums {
		if i != index {
			product *= val
		}
	}
	return product
}
