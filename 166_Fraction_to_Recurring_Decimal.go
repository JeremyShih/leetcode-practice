// 2020/5/3
package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	n, d := 1, 2
	//fmt.Println(fractionToDecimal(n, d))
	fmt.Println(fractionToDecimal(n, d) == "0.5")
	n, d = 2, 1
	fmt.Println(fractionToDecimal(n, d) == "2")
	n, d = 2, -1
	fmt.Println(fractionToDecimal(n, d) == "-2")
	n, d = 2, 3
	//fmt.Println(fractionToDecimal(n, d))
	fmt.Println(fractionToDecimal(n, d) == "0.(6)")
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return ""
	}
	if numerator == 0 {
		return "0"
	}

	ans := ""
	if numerator*denominator < 0 {
		ans += "-"
	}
	n, d := int(math.Abs(float64(numerator))), int(math.Abs(float64(denominator)))
	q, r := n/d, n%d
	ans += strconv.Itoa(q)
	if r == 0 {
		return ans
	}
	ans += "."

	hs := map[int]int{}
	for r != 0 {
		if i, ok := hs[r]; ok {
			return ans[:i] + "(" + ans[i:] + ")"
		}
		hs[r] = len(ans)
		q, r = r*10/d, r*10%d
		ans += strconv.Itoa(q)
	}
	return ans
}
