package main

import "fmt"

func DFS(graph map[int][]int, start int, visited map[int]bool) {
	if visited[start] {
		return
	}

	visited[start] = true
	for _, vertex := range graph[start] {
		DFS(graph, vertex, visited)
	}
}

func main() {
	graph := make(map[int][]int)
	graph[0] = []int{1, 2}
	graph[1] = []int{0, 3, 4}
	graph[2] = []int{0}
	graph[3] = []int{1}
	graph[4] = []int{1, 5}
	graph[5] = []int{4}

	visited := make(map[int]bool)
	DFS(graph, 0, visited)
	fmt.Println(visited)
}
