package main

import (
	"fmt"
)

// https://leetcode.com/problems/jump-game
// Last Status: Failed on 76th test case due to out of memory

func main() {
	nums := []int{1, 0, 2}
	canJ := solve(nums)
	fmt.Println(canJ)
}

func solve(nums []int) bool {
	canReachLastIndex := canJump(nums)
	return canReachLastIndex
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	} else if nums[0] == 0 {
		return false
	}

	graph := makeGraph(nums)
	return bfs(nums, graph)
}

func bfs(nums []int, graph [][]bool) bool {
	visited := make(map[int]bool)
	queue := make([]int, 1, len(nums))

	queue[0] = 0
	queuePtr := 0
	for queuePtr < len(nums) && queuePtr < len(queue) {
		i := queue[queuePtr]
		if visited[i] == true {
			queuePtr += 1
			continue
		}
		for j, isAdj := range graph[i] {
			if !isAdj || i == j {
				continue
			}
			if j == len(nums)-1 {
				return true
			}
			queue = append(queue, j)
		}
		visited[i] = true
		queuePtr += 1
	}
	return false
}

func makeGraph(nums []int) [][]bool {
	graph := make([][]bool, len(nums))
	for i, _ := range graph {
		graph[i] = make([]bool, len(nums))
	}

	for ind, item := range nums {
		for i := ind; i <= ind+item; i += 1 {
			if i >= len(nums) {
				continue
			}
			graph[ind][i] = true
		}
	}

	return graph
}
