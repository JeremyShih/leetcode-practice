package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "123   "
	fmt.Println("|" + strings.TrimSpace(s) + "|")
}
