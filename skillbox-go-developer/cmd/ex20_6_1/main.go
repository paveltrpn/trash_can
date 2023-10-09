package main

import "fmt"

type mtrx3 [3][3]float32

func mtrx3LU(m mtrx3) (lm, um mtrx3) {
	const (
		mrange int32 = 3
	)

	var (
		i, j, k int32
		sum     float32
	)

	for i = 0; i < mrange; i++ {
		for k = i; k < mrange; k++ {
			sum = 0
			for j = 0; j < i; j++ {
				sum += (lm[i][j] * um[j][k])
			}
			um[i][k] = m[i][k] - sum
		}

		for k = i; k < mrange; k++ {
			if i == k {
				lm[i][i] = 1.0
			} else {
				sum = 0
				for j = 0; j < i; j++ {
					sum += lm[k][j] * um[j][i]
				}
				lm[k][i] = (m[k][i] - sum) / um[i][i]
			}
		}
	}

	return lm, um
}

func mtrx3DetLU(m mtrx3) (rt float32) {
	var (
		i     int32
		u     mtrx3
		u_det float32
	)

	_, u = mtrx3LU(m)

	u_det = u[0][0]

	for i = 1; i < 3; i++ {
		u_det *= u[i][i]
	}

	return u_det
}

func main() {
	var foo = mtrx3{{-51, 3, 59}, {5, 67, -3}, {2, -21, 1}}

	fmt.Println(mtrx3DetLU(foo))
}
