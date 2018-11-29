package main

import (
	"golang.org/x/tour/reader"
	"fmt"
)

/*
Implementing this function 

https://golang.org/pkg/io/#Reader
*/

type MyReader struct{}

func (reader MyReader) Read(info []byte) (int, error){

    for i := range info {
		// Replacing value for 'A' character
        info[i] = 'A'
    }
	return len(info), nil
}

func main() {
	reader.Validate(MyReader{})
}


