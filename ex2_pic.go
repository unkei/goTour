package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	dat := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		dat[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// dat[y][x] = uint8((x+y)/2)
			// dat[y][x] = uint8(x * y)
			dat[y][x] = uint8(x ^ y)
		}
	}
	return dat
}

func main() {
	pic.Show(Pic)
}

