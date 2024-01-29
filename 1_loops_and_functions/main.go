package main

import (
	"fmt"
	"math"
)

func SqrtNewton(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}

	var sqrtX float64 = float64(1)
	var oldSqrtX float64
	const stopCondition float64 = 1e-15

	var tmpDiff1, tmpDiff2 float64
	for {
		oldSqrtX = sqrtX
		sqrtX -= (sqrtX * sqrtX - x) / (2 * sqrtX)

		tmpDiff2 = tmpDiff1
		tmpDiff1 = oldSqrtX - sqrtX

		if tmpDiff1 < 0 {
			tmpDiff1 = -tmpDiff1
		}

		if tmpDiff1 < stopCondition || tmpDiff1 == tmpDiff2 {
			break
		}

	}

	return sqrtX
}

func getSqrtDiff(x float64) {
	if x < 0 {
		return
	}

	x1 := SqrtNewton(x)
	x2 := math.Sqrt(x)

	fmt.Println("------------------------------")
	fmt.Println("x             = ", x)
	fmt.Println("SqrtNewton(x) = ", x1)
	fmt.Println("math.Sqrt(x)  = ", x2)
	if x1 < x2 {
		fmt.Println("diff          = ", x2 - x1)
	} else {
		fmt.Println("diff          = ", x1 - x2)
	}
	fmt.Println("------------------------------")
}

func main() {
	getSqrtDiff(2)
	getSqrtDiff(100)
	getSqrtDiff(1337.123)
	getSqrtDiff(2e11 - 1)
	getSqrtDiff(2e23 - 1)
	getSqrtDiff(0.00000005)
	getSqrtDiff(2e-15)
}
