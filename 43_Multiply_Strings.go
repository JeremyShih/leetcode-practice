// 2020/4/4
package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "2"
	num2 := "3"
	fmt.Println(multiply(num1, num2) == "6")
	num1 = "123"
	num2 = "456"
	fmt.Println(multiply(num1, num2) == "56088")
	num1 = "498828660196"
	num2 = "840477629533"
	fmt.Println(multiply(num1, num2) == "419254329864656431168468")
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	result := make([]int, len(num1+num2))
	for i := len(num1) - 1; i > -1; i-- {
		for j := len(num2) - 1; j > -1; j-- {
			result[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}
	carry, sum := 0, 0
	for i := len(result) - 1; i > -1; i-- {
		sum = result[i] + carry
		result[i] = sum % 10
		carry = sum / 10
	}
	for result[0] == 0 {
		result = result[1:]
	}
	ans := ""
	for _, v := range result {
		ans += strconv.Itoa(v)
	}
	return ans
}
