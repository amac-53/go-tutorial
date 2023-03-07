package main

import (
	"golang.org/x/tour/pic"
	// "fmt"
)

func Pic(dx, dy int) [][]uint8 {
	// fmt.Println(dx, dy) // 256, 256
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	} 
	return pic
}

func main() {
	pic.Show(Pic)
}
