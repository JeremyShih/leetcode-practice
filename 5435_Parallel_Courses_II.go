// 2020/6/27
package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	n := 4
	dependencies := [][]int{{2, 1}, {3, 1}, {1, 4}}
	k := 2
	fmt.Println(minNumberOfSemesters(n, dependencies, k) == 3)
	n = 5
	dependencies = [][]int{{2, 1}, {3, 1}, {4, 1}, {1, 5}}
	k = 2
	fmt.Println(minNumberOfSemesters(n, dependencies, k) == 4)
	n = 11
	dependencies = [][]int{}
	k = 2
	fmt.Println(minNumberOfSemesters(n, dependencies, k) == 6)
	n = 4
	dependencies = [][]int{{2, 1}, {2, 4}}
	k = 2
	fmt.Println(minNumberOfSemesters(n, dependencies, k) == 2)
	n = 5
	dependencies = [][]int{{1, 5}, {1, 3}, {1, 2}, {4, 2}, {4, 5}, {2, 5}, {1, 4}, {4, 3}, {3, 5}, {3, 2}}
	k = 3
	fmt.Println(minNumberOfSemesters(n, dependencies, k) == 5)
	n = 5
	dependencies = [][]int{{5, 1}, {3, 1}, {5, 4}, {4, 1}, {2, 3}}
	k = 3
	fmt.Println(minNumberOfSemesters(n, dependencies, k) == 3)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	m := make(map[int][]int)
	for i := 1; i <= n; i++ {
		m[i] = []int{}
	}
	for _, v := range dependencies {
		m[v[1]] = append(m[v[1]], v[0])
	}
	var nopre []int
	for k, v := range m {
		if len(v) == 0 {
			nopre = append(nopre, k)
		}
	}

	taken := nopre
	if len(nopre)%k != 0 {
		taken = nopre[:len(nopre)-len(nopre)%k]
		taken = append(taken, classesForThisSeme(m, taken, k-len(taken)%k)...)
	}
	seme := len(taken) / k
	if len(taken)%k != 0 {
		seme++
	}

	for len(taken) < n {
		toTake := classesForThisSeme(m, taken, k)
		taken = append(taken, toTake...)
		seme++
	}

	return seme
}

func classesForThisSeme(preq map[int][]int, taken []int, count int) []int {
	ans := []int{}
	for k, v := range preq {
		if ok, _ := Contain(k, taken); ok {
			continue
		}
		pass := true
		for _, c := range v {
			if ok, _ := Contain(c, taken); !ok {
				pass = false
				break
			}
		}
		if pass {
			ans = append(ans, k)
		}
		//fmt.Println("ans:", ans)
		if len(ans) == count {
			break
		}
	}
	return ans
}

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}
