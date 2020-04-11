// 2020/4/5
package main

import (
	"fmt"
	// "strings"
)

func main() {
	s := "1101"
	fmt.Println(numSteps(s) == 6)
	s = "10"
	fmt.Println(numSteps(s) == 1)
	s = "1"
	fmt.Println(numSteps(s) == 0)
	s = "1111110011101010110011100100101110010100101110111010111110110010"
	fmt.Println(numSteps(s) == 89)
	s = "100100001010010101101000111100100111101111000111111010010001100001000100011001"
	fmt.Println(numSteps(s) == 120)
}

func numSteps(s string) int {
	count, i := 0, len(s)-1
	for i > 0 {
		if s[i] == '0' {
			count++
			i--
		} else {
			count++
			for s[i] == '1' && i > 0 {
				count++
				i--
			}
			if i == 0 {
				count++
			}
			out := []rune(s)
			out[i] = '1'
			s = string(out)
		}
		// fmt.Println(s)
	}
	return count
}
