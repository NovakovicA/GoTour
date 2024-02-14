package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot calculate sqrt of negative number (%v)!", float64(err))
}

func SqrtNewton(x float64) (float64, error) {
	if x < 0 {
		return math.NaN(), ErrNegativeSqrt(x)
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

	return sqrtX, nil
}

func getSqrtDiff(x float64) {

	x1, err := SqrtNewton(x)

	if err != nil {
		fmt.Println(err)
		return
	}

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
	getSqrtDiff(-2)
}
