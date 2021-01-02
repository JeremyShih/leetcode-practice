// 2020/5/10
package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	target, n := []int{1, 3}, 3
	fmt.Println(reflect.DeepEqual(buildArray(target, n), []string{"Push", "Push", "Pop", "Push"}))
	target, n = []int{1, 2, 3}, 3
	fmt.Println(reflect.DeepEqual(buildArray(target, n), []string{"Push", "Push", "Push"}))
	target, n = []int{1, 2}, 4
	fmt.Println(reflect.DeepEqual(buildArray(target, n), []string{"Push", "Push"}))
	target, n = []int{2, 3, 4}, 4
	fmt.Println(reflect.DeepEqual(buildArray(target, n), []string{"Push", "Pop", "Push", "Push", "Push"}))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func buildArray(target []int, n int) []string {
	ans, nums := []string{}, []int{}
	push, pop := "Push", "Pop"
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
		if ok, _ := Contain(i, target); ok {
			ans = append(ans, push)
			//fmt.Println(ans)
			target = target[1:]
		} else {
			if len(target) > 0 {
				ans = append(ans, push)
				ans = append(ans, pop)
			}
		}
		//fmt.Println(i, target)

		if len(target) == 0 {
			return ans
		}
	}
	fmt.Println(ans)
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
