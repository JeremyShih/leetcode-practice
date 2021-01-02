// 2020/12/20
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	number := "1-23-45 6"
	fmt.Println(reformatNumber(number) == "123-456")
	number = "123 4-567"
	fmt.Println(reformatNumber(number) == "123-45-67")
	number = "123 4-5678"
	fmt.Println(reformatNumber(number) == "123-456-78")
	number = "12"
	fmt.Println(reformatNumber(number) == "12")
	number = "--17-5 229 35-39475 "
	fmt.Println(reformatNumber(number) == "175-229-353-94-75")
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func reformatNumber(number string) string {
	number = strings.Replace(number, " ", "", -1)
	number = strings.Replace(number, "-", "", -1)
	//fmt.Println(number)

	res := ""
	for len(number) > 0 {
		l := len(number)
		if l > 4 {
			res += number[:3] + "-"
			number = number[3:]
			//fmt.Println(res, number)
		} else if l == 4 {
			res += number[:2] + "-" + number[2:]
			number = number[4:]
			//fmt.Println(res, number)
		} else if l == 3 {
			res += number[:3]
			number = number[3:]
			//fmt.Println(res, number)
		} else {
			res += number
			number = ""
		}
	}
	//fmt.Println(res)
	return res
}
