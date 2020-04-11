// 2020/3/7
package main

import "fmt"

func main() {
	minStack := new(MinStack)
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin() == -3) //  --> Returns -3.
	minStack.Pop()
	minStack.Top()                       //     --> Returns 0.
	fmt.Println(minStack.GetMin() == -2) //  --> Returns -2.
}

// Stack save all the items
// Min save index of the min item
type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: []int{}, min: []int{}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 {
		this.min = append(this.min, 0)
	} else {
		if this.stack[this.min[len(this.min)-1]] < x {
			// min item is the original one
			this.min = append(this.min, this.min[len(this.min)-1])
		} else {
			// min item is the new added item
			this.min = append(this.min, len(this.stack)-1)
		}
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.min[len(this.min)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
