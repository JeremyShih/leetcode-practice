// 2020/11/15
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	os := Constructor(5)
	fmt.Println(reflect.DeepEqual(os.Insert(3, "ccccc"), []string{}))
	fmt.Println(reflect.DeepEqual(os.Insert(1, "aaaaa"), []string{"aaaaa"}))
	fmt.Println(reflect.DeepEqual(os.Insert(2, "bbbbb"), []string{"bbbbb", "ccccc"}))
	//fmt.Println(os.Insert(2, "bbbbb"))
	fmt.Println(reflect.DeepEqual(os.Insert(5, "eeeee"), []string{}))
	fmt.Println(reflect.DeepEqual(os.Insert(4, "ddddd"), []string{"ddddd", "eeeee"}))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

type OrderedStream struct {
	ptr      int
	sequence []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		ptr:      1,
		sequence: make([]string, n),
	}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.sequence[id-1] = value

	if id == this.ptr {
		i := this.ptr
		for i < len(this.sequence) && this.sequence[i] != "" {
			i++
		}
		this.ptr = i + 1

		return this.sequence[id-1 : i]
	}
	return []string{}
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(id,value);
 */
