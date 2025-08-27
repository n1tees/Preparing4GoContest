package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const K = 2

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	firstLine := scanner.Text()
	sizes := strings.Fields(firstLine)

	if len(sizes) != 2 {
		fmt.Println("Invalid input")
		return
	}

	m, err1 := strconv.Atoi(sizes[0])
	n, err2 := strconv.Atoi(sizes[1])

	if err1 != nil || err2 != nil || m <= 0 || n <= 0 {
		fmt.Println(0)
		return
	}

	matrix := make([][]byte, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		line := scanner.Text()
		matrix[i] = []byte(line)
	}

	result := 0
	currentLen := 0

	for i := 0; i < m; i++ {

		currentLen = 0
		for j := 0; j < n; j++ {
			if matrix[i][j] != '#' {
				currentLen++
			} else {
				if currentLen >= K {
					result += countWords(currentLen, K)
				}
				currentLen = 0
			}
		}

		if currentLen >= K {
			result += countWords(currentLen, K)
		}

	}

	for j := 0; j < n; j++ {
		currentLen = 0
		for i := 0; i < m; i++ {
			if matrix[i][j] != '#' {
				currentLen++
			} else {
				if currentLen >= K {
					result += countWords(currentLen, K)
				}
				currentLen = 0
			}
		}

		if currentLen >= K {
			result += countWords(currentLen, K)
		}
	}

	fmt.Println(result)

}

func countWords(L, K int) int {
	return (L - K + 1) * (L - K + 2) / 2
}
