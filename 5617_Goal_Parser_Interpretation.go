// 2020/12/6
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	command := "G()(al)"
	fmt.Println(strings.EqualFold(interpret(command), "Goal"))
	command = "G()()()()(al)"
	fmt.Println(strings.EqualFold(interpret(command), "Gooooal"))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func interpret(command string) string {
	command = strings.Replace(command, "()", "o", -1)
	command = strings.Replace(command, "(al)", "al", -1)
	return command
}
