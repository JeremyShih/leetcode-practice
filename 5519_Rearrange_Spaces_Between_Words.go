// 2020/9/20
package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	text := "  this   is  a sentence "
	ans := "this   is   a   sentence"
	fmt.Println(strings.EqualFold(reorderSpaces(text), ans))
	text = " practice   makes   perfect"
	ans = "practice   makes   perfect "
	fmt.Println(strings.EqualFold(reorderSpaces(text), ans))
	text = "hello   world"
	ans = "hello   world"
	fmt.Println(strings.EqualFold(reorderSpaces(text), ans))
	text = "  walks  udp package   into  bar a"
	ans = "walks  udp  package  into  bar  a "
	fmt.Println(strings.EqualFold(reorderSpaces(text), ans))
	text = "a"
	ans = "a"
	fmt.Println(strings.EqualFold(reorderSpaces(text), ans))
	text = "   hello"
	ans = "hello   "
	fmt.Println(strings.EqualFold(reorderSpaces(text), ans))
	text = "a b c "
	ans = "a b c "
	fmt.Println(strings.EqualFold(reorderSpaces(text), ans))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func reorderSpaces(text string) string {
	space := regexp.MustCompile(" ")
	matches := space.FindAllStringIndex(text, -1)
	spaceCount := len(matches)
	if spaceCount == 0 {
		return text
	}
	//fmt.Println(len(matches))
	split := strings.Split(text, " ")
	//fmt.Println(strings.Join(split, ","))
	newSlice := []string{}
	for _, s := range split {
		if len(s) > 0 {
			newSlice = append(newSlice, s)
		}
	}
	if len(newSlice) == 1 {
		return newSlice[0] + strings.Repeat(" ", spaceCount)
	}
	divide := spaceCount / (len(newSlice) - 1)
	remain := spaceCount - divide*(len(newSlice)-1)
	//fmt.Println(divide, remain)
	newText := strings.Join(newSlice, strings.Repeat(" ", divide)) + strings.Repeat(" ", remain)
	return newText
}
