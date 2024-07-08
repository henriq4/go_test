package main

func NonParallelMultiply(matrix1 [][]int, matrix2 [][]int) {
	result := make([][]int, len(matrix1))
	for i := range result {
		result[i] = make([]int, len(matrix2[0]))
		for j := range result[i] {
			result[i][j] = 0
			for k := range matrix2 {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
}