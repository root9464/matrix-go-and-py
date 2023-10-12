package main

import "fmt"

const N = 10 //обьявление матрицы x,y
const (
	TOP = iota
	LEFT
	DOWN
	RIGHT
)

func main() {
	var m [N][N]int
	var i, j = N / 2, N / 2
	var iMin, iMax = i, i
	var jMin, jMax = j, j

	d := TOP

	for step := 1; step < N*N+1; step++ {
		m[i][j] = step
		switch d {
		case TOP:
			i = i - 1
			if i < iMin {
				d = LEFT
				iMin = i
			}
		case LEFT:
			j = j - 1
			if j < jMin {
				d = DOWN
				jMin = j
			}
		case DOWN:
			i = i + 1
			if i > iMax {
				d = RIGHT
				iMax = i
			}
		case RIGHT:
			j = j + 1
			if j > jMax {
				d = TOP
				jMax = j
			}
		}
	}
	printSpiral(m)
}

func printSpiral(spiral [N][N]int) {
	fmt.Printf("matrix: %dx%d\n", N, N)
	fmt.Print("===========\n")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d\t", spiral[i][j])
		}
		fmt.Print("\n")
	}
}
