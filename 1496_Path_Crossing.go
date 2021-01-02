// 2020/6/29
package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	path := "NES"
	fmt.Println(!isPathCrossing(path))
	path = "NESWW"
	fmt.Println(isPathCrossing(path))
	path = "NS"
	fmt.Println(isPathCrossing(path))
	path = "W"
	fmt.Println(isPathCrossing(path))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func isPathCrossing(path string) bool {
	start := []int{0, 0}
	var rec []string
	for _, c := range path {
		rec = append(rec, strconv.Itoa(start[0])+strconv.Itoa(start[1]))
		if string(c) == "N" {
			start[1] += 1
		} else if string(c) == "E" {
			start[0] += 1
		} else if string(c) == "S" {
			start[1] -= 1
		} else if string(c) == "W" {
			start[0] -= 1
		}

		if passed, _ := Contain(strconv.Itoa(start[0])+strconv.Itoa(start[1]), rec); passed {
			return true
		}
	}
	return false
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
