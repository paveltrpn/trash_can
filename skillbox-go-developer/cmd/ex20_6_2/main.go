package main

import "fmt"

func mtrx3Mult(a [4][5]float32, b [5][4]float32) (rt [4][4]float32) {
	var (
		i, j, k int32
		tmp     float32
	)

	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			tmp = 0.0
			for k = 0; k < 5; k++ {
				tmp = tmp + a[j][k]*b[k][i]
			}
			rt[i][j] = tmp
		}
	}

	return rt
}
func main() {
	var (
		foo = [4][5]float32{{56, 77, 109, 3, -9}, {4, 5, -60, 23, 88}, {-5, 32, 19, 22, -60}, {109, 65, -501, 22, 24}}
		bar = [5][4]float32{{56, 34, -55, 12}, {5, -60, 23, 88}, {22, 69, 22, -60}, {-25, -801, 22, 24}, {439, 23, -11, 23}}
	)

	fmt.Println(mtrx3Mult(foo, bar))
}
