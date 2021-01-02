// 2020/5/17
package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	favoriteCompanies := [][]string{{"leetcode", "google", "facebook"}, {"google", "microsoft"}, {"google", "facebook"}, {"google"}, {"amazon"}}
	ans := []int{0, 1, 4}
	fmt.Println(reflect.DeepEqual(peopleIndexes(favoriteCompanies), ans))
	favoriteCompanies = [][]string{{"leetcode", "google", "facebook"}, {"leetcode", "amazon"}, {"facebook", "google"}}
	ans = []int{0, 1}
	fmt.Println(reflect.DeepEqual(peopleIndexes(favoriteCompanies), ans))
	favoriteCompanies = [][]string{{"leetcode"}, {"google"}, {"facebook"}, {"amazon"}}
	ans = []int{0, 1, 2, 3}
	fmt.Println(reflect.DeepEqual(peopleIndexes(favoriteCompanies), ans))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func peopleIndexes(favoriteCompanies [][]string) []int {
	ans := []int{}
	for i := 0; i < len(favoriteCompanies); i++ {
		cur := favoriteCompanies[i]
		test := false
		for j := 0; j < len(favoriteCompanies); j++ {
			if i == j {
				continue
			}
			if len(cur) < len(favoriteCompanies[j]) {
				if isSubSet(cur, favoriteCompanies[j]) {
					test = true
					break
				}
			}
		}
		if !test {
			ans = append(ans, i)
		}
	}
	return ans
}

func isSubSet(cur []string, companies []string) bool {
	for _, c := range cur {
		if ok, _ := Contain(c, companies); !ok {
			return false
		}
	}
	return true
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
