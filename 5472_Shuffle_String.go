// 2020/7/26
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func restoreString(s string, indices []int) string {
	tmp := make([]rune, len(indices))
	for i, c := range s {
		tmp[indices[i]] = c
	}
	return string(tmp)
}
