package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r * rot13Reader) Read(b []byte) (int, error) {
	//Reads from a io.Reader from rot13Reader
	n, err := r.r.Read(b)
	
	//Modifies the stream by applying rot13
	for i := range b{
		
		if b[i] < 'A' || b[i] > 'z' {
			continue
		}
		
		
		if b[i] > ('Z' - 13) && b[i] <= 'Z' && b[i] >= 'A' {
			// Minus 1 because we only want the distance BETWEEN b[i] and Z
			// so we exclude b[i] from the interval
			diff := b[i] - ('Z' - 13) - 1
			b[i] = 'A' + diff
		} else if b[i] > ('z' - 13) && b[i] <= 'z' && b[i] >= 'a' {
			 diff := b[i] - ('z' - 13) - 1
			b[i] = 'a' + diff
		} else {
			b[i] = b[i] + 13
		}
	}
	
	return n, err
	
}

func main() {
	//Creats new reader
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	fmt.Printf("%t", s);
	
	//Object now has reader s
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
