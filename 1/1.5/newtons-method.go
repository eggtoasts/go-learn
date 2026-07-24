package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	/* z := float64(1.0);
	for v := 1; v < 10; v++{
		z -= (z*z - x) / (2*z)
		fmt.Println(z);
	}

*/
	
	z := float64(1.0)
	
	for ; ; {
		oldZ := z;
		z -= (z*z - x) / (2*z)
		
		if math.Abs(oldZ - z) <= 1e-15 {
			return z
		}
		
	}
	return z;
}

func main() {
	fmt.Println(Sqrt(101))
}
