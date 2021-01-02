// 2020/6/7
package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	arr, k := []int{1, 2, 3, 4, 5}, 2
	ans := []int{5, 1}
	fmt.Println(ans)
	fmt.Println(reflect.DeepEqual(getStrongest(arr, k), ans))
	arr, k = []int{1, 1, 3, 5, 5}, 2
	ans = []int{5, 5}
	fmt.Println(ans)
	fmt.Println(reflect.DeepEqual(getStrongest(arr, k), ans))
	arr, k = []int{6, 7, 11, 7, 6, 8}, 5
	ans = []int{11, 8, 6, 6, 7}
	fmt.Println(ans)
	fmt.Println(reflect.DeepEqual(getStrongest(arr, k), ans))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func getStrongest(arr []int, k int) []int {
	n := len(arr)
	if k == n {
		return arr
	}
	sort.Ints(arr)
	fmt.Println(arr)
	m := arr[(n-1)/2]

	ans := []int{}
	for len(ans) < k {
		if math.Abs(float64(arr[n-1]-m)) >= math.Abs(float64(arr[0]-m)) {
			ans = append(ans, arr[n-1])
			i := n - 1
			copy(arr[i:], arr[i+1:]) // Shift a[i+1:] left one index.
			arr = arr[:len(arr)-1]   // Truncate slice.
		} else {
			ans = append(ans, arr[0])
			i := 0
			copy(arr[i:], arr[i+1:]) // Shift a[i+1:] left one index.
			arr = arr[:len(arr)-1]   // Truncate slice.
		}
		n -= 1
	}
	return ans
}
