// 2020/8/9
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	n, k := 3, 1
	fmt.Println(findKthBit(n, k) == '0')
	n, k = 4, 11
	fmt.Println(findKthBit(n, k) == '1')
	n, k = 1, 1
	fmt.Println(findKthBit(n, k) == '0')
	n, k = 2, 3
	fmt.Println(findKthBit(n, k) == '1')
	n, k = 3, 2
	fmt.Println(findKthBit(n, k) == '1')
	n, k = 3, 5
	fmt.Println(findKthBit(n, k) == '0')
	n, k = 3, 7
	fmt.Println(findKthBit(n, k) == '1')
	n, k = 4, 12
	fmt.Println(findKthBit(n, k) == '0')
	//n, k = 20, 11
	//fmt.Println(findKthBit(n, k))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func findKthBit(n int, k int) byte {
	sl := []int{0, 1, 3, 7, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095, 8191, 16383, 32767, 65535, 131071, 262143, 524287, 1048575, 2097151}

	invertTimes := 0
	for n > 1 {
		//fmt.Println(n, k, invertTimes)
		if k == sl[n-1]+1 {
			if invertTimes == 0 {
				return '1'
			} else {
				return '0'
			}
		} else {
			if k > sl[n-1]+1 {
				k = sl[n] - k + 1
				invertTimes = (invertTimes + 1) % 2
			}
		}
		n--
	}
	//fmt.Println(n, k, invertTimes)
	if invertTimes == 0 {
		return '0'
	} else {
		return '1'
	}
}
