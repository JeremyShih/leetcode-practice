// 2020/5/24
package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	s, k := "abciiidef", 3
	fmt.Println(maxVowels(s, k) == 3)
	s, k = "aeiou", 2
	fmt.Println(maxVowels(s, k) == 2)
	s, k = "leetcode", 3
	fmt.Println(maxVowels(s, k) == 2)
	s, k = "rhythms", 4
	fmt.Println(maxVowels(s, k) == 0)
	s, k = "tryhard", 4
	fmt.Println(maxVowels(s, k) == 1)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxVowels(s string, k int) int {
	vows, maxCount := []string{"a", "e", "i", "o", "u"}, 0
	vowsSlice := []bool{}
	for i := 0; i < len(s); i++ {
		if ok, _ := Contain(string(s[i]), vows); ok {
			vowsSlice = append(vowsSlice, true)
		} else {
			vowsSlice = append(vowsSlice, false)
		}
	}
	for i := 0; i < len(s)-k+1; i++ {
		//fmt.Println(s[i : i+k])
		count := 0
		for _, c := range vowsSlice[i : i+k] {
			if c {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
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
