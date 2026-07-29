package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	
	//var x [][]uint8
	x := make([][]uint8, dy)
	
	for i := 0; i < dy; i++ {
		x[i] = make([]uint8,dx)
		
		for j := 0; j < dx; j++ {
			x[i][j] = uint8((2^i) + (i+j)/2)
		}
	}
 
	return x;
}

func main() {
	pic.Show(Pic)
}
