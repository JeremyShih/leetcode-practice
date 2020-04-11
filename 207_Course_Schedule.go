// 2020/4/4
package main

import "fmt"

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(canFinish(numCourses, prerequisites))
	numCourses = 2
	prerequisites = [][]int{{1, 0}, {0, 1}}
	fmt.Println(!canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	dependencies := make([]int, numCourses)
	for _, v := range prerequisites {
		course := v[0]
		preq := v[1]
		dependencies[course] += 1
		if _, ok := graph[preq]; ok {
			graph[preq] = append(graph[preq], course)
		} else {
			graph[preq] = []int{course}
		}
	}
	fmt.Println(dependencies)
	fmt.Println(graph)
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if dependencies[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		course := queue[0]
		subCourses := graph[course]
		for _, c := range subCourses {
			dependencies[c]--
			// if there's no dependency for the subCourses means we can take the course now
			if dependencies[c] == 0 {
				queue = append(queue, c)
			}
		}
		queue = queue[1:]
	}
	for _, c := range dependencies {
		if c != 0 {
			return false
		}
	}
	return true
}
