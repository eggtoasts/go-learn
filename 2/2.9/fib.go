package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var prev = 0
	var cur = 1
	
	return func() int {
		if prev == 0 {
			prev = 1
			return 0
		}
		
		if prev == 1 && cur == 1 {
			cur = 2
			return 1
		}
		
		tmp := prev
		prev = cur
		
		cur = prev + tmp
		
		return tmp
		
	}
	
}

func main() {
	f := fibonacci() //<-- closure!
	
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
