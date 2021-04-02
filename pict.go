package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	y := [][]uint8{}

	for j := 0; j < dy; j++ {
		x := []uint8{}
		for i := 0; i < dx; i++ {
			x = append(x, uint8((i * j)))
		}
		y = append(y, x)
	}
	return y

}

func main() {
	pic.Show(Pic)
}
