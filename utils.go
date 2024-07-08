package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomMatrix(rows, cols int) [][]int {
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(9) + 1
		}
	}
	return matrix
}

func TimeComparison(matrix1 [][]int, matrix2 [][]int) {
	startTime := time.Now()
	NonParallelMultiply(matrix1, matrix2)
	endTime := time.Now()
	fmt.Println("Sequential time: ", endTime.Sub(startTime).Milliseconds(), "ms")

	startTime = time.Now()
	ParallelMultiply(matrix1, matrix2)
	endTime = time.Now()
	fmt.Println("Concurrent time: ", endTime.Sub(startTime).Milliseconds(), "ms")
}