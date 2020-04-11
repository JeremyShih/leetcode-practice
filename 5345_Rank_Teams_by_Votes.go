// 2020/3/1
package main

import (
	"fmt"
	"strings"
)

func rankTeams(votes []string) string {
	stat := make(map[string][]int)
	score := make(map[string]map[int]int)
	for i := 0; i < len(votes[0]); i++ {
		stat[string(votes[0][i])] = []int{}
		score[string(votes[0][i])] = make(map[int]int)
	}
	for i := 0; i < len(votes); i++ {
		// for j := 0; j < len(votes[0]); j++ {
		// 	stat[votes[0][j]] = append(stat[votes[0][j]], j)
		// }
		for k := range stat {
			stat[k] = append(stat[k], strings.Index(votes[i], k))
		}
	}
	fmt.Println(stat)
	// for i := 0; i < len(votes[0]); i++ {

	// }
	return votes[0]
}

func main() {
	inp := []string{"ABC", "ACB", "ABC", "ACB", "ACB"}
	fmt.Println(inp)
	fmt.Println(rankTeams(inp))
}
