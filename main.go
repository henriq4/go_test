package main

import (
	"fmt"
)

func main() {
	var rowsA, colsA, rowsB, colsB int 

	if colsA != rowsB {
		fmt.Println("Error...")
		return
	}

	fmt.Print("Digite o tamanho das matrizes: ")
	fmt.Scan(&rowsA)
	rowsB = rowsA
	colsA = rowsA
	colsB = rowsA

	fmt.Println("Gerando matriz A...")
	matrixA := generateRandomMatrix(rowsA, colsA)
	fmt.Println("Matriz A gerada")

	fmt.Println("Gerando matriz B...")
	matrixB := generateRandomMatrix(rowsB, colsB)
	fmt.Println("Matriz A gerada")

	fmt.Println("Multiplicando matrizes...")

	TimeComparison(matrixA, matrixB)
}