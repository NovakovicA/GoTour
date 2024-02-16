package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var resultPic [][]uint8 = make([][]uint8, dy)

	for i := range resultPic {
		resultPic[i] = make([]uint8, dx)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			resultPic[i][j] = uint8(i*j);
		}
	}

	return resultPic
}

func main() {
	pic.Show(Pic)
}
