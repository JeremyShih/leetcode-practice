package main

import "fmt"

func main() {
	inp := "aaabb"
	a := [][]string{{"a", "a", "a", "b", "b"}, {"a", "a", "a", "bb"}, {"a", "aa", "b", "b"}, {"a", "aa", "bb"}, {"aa", "a", "b", "b"}, {"aa", "a", "bb"}, {"aaa", "b", "b"}, {"aaa", "bb"}}
	fmt.Println(inp, String2DSliceEqual(partition(inp), a))
	inp = "efe"
	a = [][]string{{"e", "f", "e"}, {"efe"}}
	fmt.Println(inp, String2DSliceEqual(partition(inp), a))
	inp = "ababbbabbaba"
	a = [][]string{{"a", "b", "a", "b", "b", "b", "a", "b", "b", "a", "b", "a"}, {"a", "b", "a", "b", "b", "b", "a", "b", "b", "aba"}, {"a", "b", "a", "b", "b", "b", "a", "b", "bab", "a"}, {"a", "b", "a", "b", "b", "b", "a", "bb", "a", "b", "a"}, {"a", "b", "a", "b", "b", "b", "a", "bb", "aba"}, {"a", "b", "a", "b", "b", "b", "abba", "b", "a"}, {"a", "b", "a", "b", "b", "bab", "b", "a", "b", "a"}, {"a", "b", "a", "b", "b", "bab", "b", "aba"}, {"a", "b", "a", "b", "b", "bab", "bab", "a"}, {"a", "b", "a", "b", "b", "babbab", "a"}, {"a", "b", "a", "b", "bb", "a", "b", "b", "a", "b", "a"}, {"a", "b", "a", "b", "bb", "a", "b", "b", "aba"}, {"a", "b", "a", "b", "bb", "a", "b", "bab", "a"}, {"a", "b", "a", "b", "bb", "a", "bb", "a", "b", "a"}, {"a", "b", "a", "b", "bb", "a", "bb", "aba"}, {"a", "b", "a", "b", "bb", "abba", "b", "a"}, {"a", "b", "a", "b", "bbabb", "a", "b", "a"}, {"a", "b", "a", "b", "bbabb", "aba"}, {"a", "b", "a", "bb", "b", "a", "b", "b", "a", "b", "a"}, {"a", "b", "a", "bb", "b", "a", "b", "b", "aba"}, {"a", "b", "a", "bb", "b", "a", "b", "bab", "a"}, {"a", "b", "a", "bb", "b", "a", "bb", "a", "b", "a"}, {"a", "b", "a", "bb", "b", "a", "bb", "aba"}, {"a", "b", "a", "bb", "b", "abba", "b", "a"}, {"a", "b", "a", "bb", "bab", "b", "a", "b", "a"}, {"a", "b", "a", "bb", "bab", "b", "aba"}, {"a", "b", "a", "bb", "bab", "bab", "a"}, {"a", "b", "a", "bb", "babbab", "a"}, {"a", "b", "a", "bbb", "a", "b", "b", "a", "b", "a"}, {"a", "b", "a", "bbb", "a", "b", "b", "aba"}, {"a", "b", "a", "bbb", "a", "b", "bab", "a"}, {"a", "b", "a", "bbb", "a", "bb", "a", "b", "a"}, {"a", "b", "a", "bbb", "a", "bb", "aba"}, {"a", "b", "a", "bbb", "abba", "b", "a"}, {"a", "b", "abbba", "b", "b", "a", "b", "a"}, {"a", "b", "abbba", "b", "b", "aba"}, {"a", "b", "abbba", "b", "bab", "a"}, {"a", "b", "abbba", "bb", "a", "b", "a"}, {"a", "b", "abbba", "bb", "aba"}, {"a", "bab", "b", "b", "a", "b", "b", "a", "b", "a"}, {"a", "bab", "b", "b", "a", "b", "b", "aba"}, {"a", "bab", "b", "b", "a", "b", "bab", "a"}, {"a", "bab", "b", "b", "a", "bb", "a", "b", "a"}, {"a", "bab", "b", "b", "a", "bb", "aba"}, {"a", "bab", "b", "b", "abba", "b", "a"}, {"a", "bab", "b", "bab", "b", "a", "b", "a"}, {"a", "bab", "b", "bab", "b", "aba"}, {"a", "bab", "b", "bab", "bab", "a"}, {"a", "bab", "b", "babbab", "a"}, {"a", "bab", "bb", "a", "b", "b", "a", "b", "a"}, {"a", "bab", "bb", "a", "b", "b", "aba"}, {"a", "bab", "bb", "a", "b", "bab", "a"}, {"a", "bab", "bb", "a", "bb", "a", "b", "a"}, {"a", "bab", "bb", "a", "bb", "aba"}, {"a", "bab", "bb", "abba", "b", "a"}, {"a", "bab", "bbabb", "a", "b", "a"}, {"a", "bab", "bbabb", "aba"}, {"a", "babbbab", "b", "a", "b", "a"}, {"a", "babbbab", "b", "aba"}, {"a", "babbbab", "bab", "a"}, {"aba", "b", "b", "b", "a", "b", "b", "a", "b", "a"}, {"aba", "b", "b", "b", "a", "b", "b", "aba"}, {"aba", "b", "b", "b", "a", "b", "bab", "a"}, {"aba", "b", "b", "b", "a", "bb", "a", "b", "a"}, {"aba", "b", "b", "b", "a", "bb", "aba"}, {"aba", "b", "b", "b", "abba", "b", "a"}, {"aba", "b", "b", "bab", "b", "a", "b", "a"}, {"aba", "b", "b", "bab", "b", "aba"}, {"aba", "b", "b", "bab", "bab", "a"}, {"aba", "b", "b", "babbab", "a"}, {"aba", "b", "bb", "a", "b", "b", "a", "b", "a"}, {"aba", "b", "bb", "a", "b", "b", "aba"}, {"aba", "b", "bb", "a", "b", "bab", "a"}, {"aba", "b", "bb", "a", "bb", "a", "b", "a"}, {"aba", "b", "bb", "a", "bb", "aba"}, {"aba", "b", "bb", "abba", "b", "a"}, {"aba", "b", "bbabb", "a", "b", "a"}, {"aba", "b", "bbabb", "aba"}, {"aba", "bb", "b", "a", "b", "b", "a", "b", "a"}, {"aba", "bb", "b", "a", "b", "b", "aba"}, {"aba", "bb", "b", "a", "b", "bab", "a"}, {"aba", "bb", "b", "a", "bb", "a", "b", "a"}, {"aba", "bb", "b", "a", "bb", "aba"}, {"aba", "bb", "b", "abba", "b", "a"}, {"aba", "bb", "bab", "b", "a", "b", "a"}, {"aba", "bb", "bab", "b", "aba"}, {"aba", "bb", "bab", "bab", "a"}, {"aba", "bb", "babbab", "a"}, {"aba", "bbb", "a", "b", "b", "a", "b", "a"}, {"aba", "bbb", "a", "b", "b", "aba"}, {"aba", "bbb", "a", "b", "bab", "a"}, {"aba", "bbb", "a", "bb", "a", "b", "a"}, {"aba", "bbb", "a", "bb", "aba"}, {"aba", "bbb", "abba", "b", "a"}}
	// fmt.Println(len(partition(inp)) == len(a))
	fmt.Println(inp, String2DSliceEqual(partition(inp), a))
}

var ans [][]string

func partition(s string) [][]string {
	ans = [][]string{}
	dfs([]string{}, s)
	return ans
}

func dfs(l []string, s string) {
	if len(s) == 0 {
		ans = append(ans, l)
	} else {
		for i := 1; i < len(s)+1; i++ {
			if isPalindrome(s[:i]) {
				tmp := make([]string, len(l))
				copy(tmp, l)
				dfs(append(tmp, s[:i]), s[i:])
			}
		}
	}
}

func isPalindrome(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func String2DSliceEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if !StringSliceEqual(v, b[i]) {
			return false
		}
	}

	return true
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
