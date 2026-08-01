package main

import (
	"fmt"
)

type MyReader struct{}

type ReaderError struct{
	What []byte	
}

func (r ReaderError) Error() string {
	return fmt.Sprintf("Infinite Loop.")
}

//Rewrite all the values in []byte into 'A's.
func (r MyReader) Read(b []byte) (int, error){
	for i := range b{
		b[i] = 'A'
	}
	
	//Returns nil because if we reached the end, then there should be no error
	return len(b), nil
}

func main() {
	// reader.Validate(MyReader{})
}
