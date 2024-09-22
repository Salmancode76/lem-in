package pkg

import (
	"sort"
	"strings"
)

// Find all valid paths using DFS
func Edmonds_Karp(rooms []string, links []string) [][]string {
	graph := make(map[string][]string)

	for _, link := range links {
		line := strings.Split(link, "-")
		u := line[0]
		v := line[1]

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	/*

		USED TO PRINT ALL THE ROOMS CONNECTIONS

		for room, connections := range graph {
			fmt.Println(room, connections)
		}

	*/

	start := rooms[0]

	target := rooms[len(rooms)-1]

	All_Paths := DSF(graph, start, target)

	sort.Slice(All_Paths, func(i, j int) bool {
		return len(All_Paths[i]) < len(All_Paths[j])
	})

	/*
		USED TO PRINT ALL PATHS

		for i, path := range All_Paths {
			fmt.Printf("Path %d: %v\n", i+1, path)
		}
	*/

	return All_Paths
}

func DSF(graph map[string][]string, start string, target string) [][]string {

	var All_Paths [][]string

	stack := [][]string{{start}}

	for len(stack) > 0 {
		path := stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		current := path[len(path)-1]

		if current == target {
			All_Paths = append(All_Paths, path)
		}

		for _, neighbor := range graph[current] {
			if !contains(path, (neighbor)) {
				newPath := append([]string{}, path...)
				newPath = append(newPath, (neighbor))
				stack = append(stack, newPath)
			}
		}

	}

	return All_Paths

}

func contains(path []string, room string) bool {
	for _, element := range path {
		if element == room {
			return true
		}
	}
	return false
}
