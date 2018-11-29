package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
	"unicode"
)

func WordCount(s string) map[string]int {
	// Defining function to get values all letters in s
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	
	// Initialize map 
	result := make(map[string]int)	
	// Iterate over words plited by FieldsFunc
	for index, val := range strings.FieldsFunc(s, f) {
		fmt.Println(index, val)
		if val_, ok := result[val]; ok {
			// If value exists, increase by 1
			fmt.Println("existe en map", val_, ok)
			result[val] +=1			
		} else {
			// Otherwise append to map 
			result[val] = 1
		} 
	}	
	return result // map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
