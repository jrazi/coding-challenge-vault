package main

// https://quera.org/problemset/237883
// wtf why am I doing this

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nStr := scanner.Text()

	n, _ := strconv.ParseInt(nStr, 10, 64)

	nThDigit := getDigitN(n)
	fmt.Println(string(nThDigit))
}

func getDigitN(n int64) (s rune) {
	if n%2 == 0 {
		return 'b'
	}
	return 'a'
}
