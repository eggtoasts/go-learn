package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	What float64
}

//Makes it an error.
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number : %v", e.What);
}

// Only for non-complex numbers
func Sqrt(x float64) (float64, error) {
	if(x < 0){
		//Returns an error value when given a neg. number
		return x, ErrNegativeSqrt{x}
	}
	z := float64(1.0)
	
	for ; ; {
		oldZ := z;
		z -= (z*z - x) / (2*z)
		
		if math.Abs(oldZ - z) <= 1e-15 {
			return z, nil
		}
		
	}
	return z, nil
}

func check(x float64) {
	z, err := Sqrt(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(z)
	}
	return
}

func main() {
	//fmt.Println(Sqrt(2))
	//fmt.Println(Sqrt(-2))
	
	check(2)
	check(-2)
}
