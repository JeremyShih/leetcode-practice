// 2020/4/12
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s := "&amp; is an HTML entity but &ambassador; is not."
	fmt.Println(entityParser(s) == "& is an HTML entity but &ambassador; is not.")
	s = "and I quote: &quot;...&quot;"
	fmt.Println(entityParser(s) == "and I quote: \"...\"")
	s = "Stay home! Practice on Leetcode :)"
	fmt.Println(entityParser(s) == "Stay home! Practice on Leetcode :)")
	s = "x &gt; y &amp;&amp; x &lt; y is always false"
	fmt.Println(entityParser(s) == "x > y && x < y is always false")
	s = "leetcode.com&frasl;problemset&frasl;all"
	fmt.Println(entityParser(s) == "leetcode.com/problemset/all")
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func entityParser(text string) string {
	pair := map[string]string{"&quot;": "\"", "&apos;": "'", "&amp;": "&", "&gt;": ">", "&lt;": "<", "&frasl;": "/"}
	for k, v := range pair {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}
