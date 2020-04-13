// 2020/4/11 give up and use python instead
package main

import "fmt"

func main() {
	s := "1 + 1"
	fmt.Println(calculate(s) == 2)
	s = " 2-1 + 2 "
	fmt.Println(calculate(s) == 3)
	s = "(1+(4+5+2)-3)+(6+8)"
	fmt.Println(calculate(s) == 23)
}

func calculate(s string) int {
	tmp := []rune(s)
	fmt.Println(tmp)
	stack, i := []rune{}, 0
	for i < len(s) {
		if s[i] == '+' || s[i] == '-' {
			stack = append(stack, s[i])
		} else if s[i] <= '9' && s[i] >= '0' {
			start := i
			for i < len(s) && s[i] <= '9' && s[i] >= '0' {
				i += 1
			}
			i -= 1
			num := s[start:i+1] - '0'
			if len(stack) == 0 || s[len(s)-1] != '+' || s[len(s)-1] != '-' {
				stack = append(stack, num)
			} else {
				stack = append(stack, evalSingle(num, stack))
			}
		}
		i += 1
	}
	return stack[0]
}

func evalSingle(num rune, stack []rune) rune {
	oper := stack[0]
	numA := stack[1]
	if oper == '+' {
		num += numA
	} else {
		num -= numA
	}
	stack = stack[2 : len(stack)-1]
	return num
}
