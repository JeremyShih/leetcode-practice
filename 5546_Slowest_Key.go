// 2020/10/25
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	releaseTimes := []int{9, 29, 49, 50}
	keysPressed := "cbcd"
	fmt.Println(slowestKey(releaseTimes, keysPressed) == 'c')
	releaseTimes = []int{12, 23, 36, 46, 62}
	keysPressed = "spuda"
	fmt.Println(slowestKey(releaseTimes, keysPressed) == 'a')
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	duration := []int{releaseTimes[0]}
	maxDuration := duration[0]
	for i := 1; i < len(releaseTimes); i++ {
		d := releaseTimes[i] - releaseTimes[i-1]
		duration = append(duration, d)
		if d > maxDuration {
			maxDuration = d
		}
	}
	//fmt.Println(duration, maxDuration)

	maxIndex := []int{}
	for i, d := range duration {
		if d == maxDuration {
			maxIndex = append(maxIndex, i)
		}
	}
	//fmt.Println(maxIndex)

	if len(maxIndex) == 1 {
		return keysPressed[maxIndex[0]]
	}

	ans := keysPressed[maxIndex[0]]
	for _, i := range maxIndex {
		if keysPressed[i] > ans {
			ans = keysPressed[i]
		}
	}

	return ans
}
