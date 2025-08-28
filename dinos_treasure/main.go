package main

import (
	"fmt"
)

var (
	bestSum   int
	incMatrix [][]int
	visited   []bool
	treasures []int

	n, m int
)

func main() {
	getInput()

	visited[1] = true

	findBest(1, treasures[1])
	fmt.Println(bestSum)

}

func findBest(island, currentSum int) {

	if currentSum > bestSum {
		bestSum = currentSum
	}

	for _, neighbour := range incMatrix[island] {
		if !visited[neighbour] {

			visited[neighbour] = true
			findBest(neighbour, currentSum+treasures[neighbour])
			visited[neighbour] = false
		}

	}

}

func getInput() {

	fmt.Scan(&n, &m)
	treasures = make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&treasures[i])
	}

	incMatrix = make([][]int, n+1)

	for i := 0; i < m; i++ {
		var n1, n2 int
		fmt.Scan(&n1, &n2)
		incMatrix[n1] = append(incMatrix[n1], n2)
		incMatrix[n2] = append(incMatrix[n2], n1)

	}

	visited = make([]bool, n+1)
}
