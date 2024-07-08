package main

import (
	"sync"
)

func ParallelMultiply(matrix1 [][]int, matrix2 [][]int) {
	result := make([][]int, len(matrix1))
	for i := range result {
		result[i] = make([]int, len(matrix2[0]))
	}

	var wg sync.WaitGroup
	wg.Add(len(matrix1) * len(matrix2[0]))

	for i := range matrix1 {
		for j := range matrix2[0] {
			go func(i, j int) {
				for k := range matrix1[0] {
					result[i][j] += matrix1[i][k] * matrix2[k][j]
				}
				wg.Done()
			}(i, j)
		}
	}

	wg.Wait()
}