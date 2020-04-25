package main

import "fmt"

func main() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
}

// a =  1
// &a =  0xc000088008
// *&a =  1
// b =  0xc000088008
// &b =  0xc000082018
// *&b =  0xc000088008
// *b =  1
// c =  0xc000082018
// *c =  0xc000088008
// &c =  0xc000082020
// *&c =  0xc000082018
// **c =  1
// ***&*&*&*&c =  1
// x =  1
