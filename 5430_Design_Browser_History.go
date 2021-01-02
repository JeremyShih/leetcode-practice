// 2020/6/7
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	obj := Constructor("homepage")
	obj.Visit("url")
	param_2 := obj.Back(1)
	fmt.Println(param_2)
	param_3 := obj.Forward(1)
	fmt.Println(param_3)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

type BrowserHistory struct {
	index   int
	history []string
}

func Constructor(homepage string) BrowserHistory {
	p := new(BrowserHistory)
	p.index = 0
	p.history = []string{homepage}
	return *p
}

func (this *BrowserHistory) Visit(url string) {
	this.history = this.history[:this.index+1]
	this.history = append(this.history, url)
	this.index += 1
}

func (this *BrowserHistory) Back(steps int) string {
	if this.index-steps < 0 {
		this.index = 0
	} else {
		this.index -= steps

	}
	return this.history[this.index]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.index+steps > len(this.history) {
		this.index = len(this.history) - 1
	} else {
		this.index += steps
	}
	return this.history[this.index]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
