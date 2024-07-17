package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAXN = 200001

var (
	n, m    int
	a, b, c [MAXN]int
	results []int
)

func countPatterns(r int) int {
	count := 0
	for i := 1; i <= r-2; i++ {
		for j := i + 1; j <= r-1; j++ {
			if b[a[i]] != a[j] {
				continue
			}
			for k := j + 1; k <= r; k++ {
				if a[j] == c[a[k]] {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ = strconv.Atoi(parts[0])
	m, _ = strconv.Atoi(parts[1])

	for i := 0; i < 3; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		for j := 1; j <= n; j++ {
			val, _ := strconv.Atoi(parts[j-1])
			switch i {
			case 0:
				a[j] = val
			case 1:
				b[j] = val
			case 2:
				c[j] = val
			}
		}
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		query := scanner.Text()
		if strings.HasPrefix(query, "CHANGE") {
			parts := strings.Split(query[7:len(query)-1], ",")
			k, _ := strconv.Atoi(parts[0])
			x, _ := strconv.Atoi(parts[1])
			a[k] = x
		} else if strings.HasPrefix(query, "PRINT") {
			r, _ := strconv.Atoi(query[6 : len(query)-1])
			results = append(results, countPatterns(r))
		}
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
