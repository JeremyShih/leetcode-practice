// 2020/11/15
package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	word1, word2 := "abc", "bca"
	fmt.Println(closeStrings(word1, word2))
	word1, word2 = "a", "aa"
	fmt.Println(!closeStrings(word1, word2))
	word1, word2 = "cabbba", "abbccc"
	fmt.Println(closeStrings(word1, word2))
	word1, word2 = "cabbba", "abbcca"
	fmt.Println(!closeStrings(word1, word2))
	word1, word2 = "cabbba", "aabbss"
	fmt.Println(!closeStrings(word1, word2))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	map1, map2 := make(map[uint8]int), make(map[uint8]int)
	for i := range word1 {
		c := word1[i]
		if v, ok := map1[c]; ok {
			map1[c] = v + 1
		} else {
			map1[c] = 1
		}

		c1 := word2[i]
		if v, ok := map2[c1]; ok {
			map2[c1] = v + 1
		} else {
			map2[c1] = 1
		}
	}

	if len(map1) != len(map2) {
		return false
	}

	for k := range map1 {
		if _, ok := map2[k]; !ok {
			return false
		}
	}

	_, map1v := []uint8{}, []int{}
	for _, v := range map1 {
		map1v = append(map1v, v)
	}
	_, map2v := []uint8{}, []int{}
	for _, v := range map2 {
		map2v = append(map2v, v)
	}

	sort.Ints(map1v)
	sort.Ints(map2v)
	if !reflect.DeepEqual(map1v, map2v) {
		return false
	}
	return true
}
