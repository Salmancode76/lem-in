package pkg

// CheckCollision checks if any room in the path has been visited.
func CheckCollision(path []string, visited map[string]bool) bool {
	for _, room := range path {
		if visited[room] {
			return true
		}
	}
	return false
}

// RedundantPaths filters out paths that contain visited rooms.
func RedundantPaths(All_Paths [][]string) [][]string {
	visited := make(map[string]bool)
	var All_Paths_New [][]string

	for _, path := range All_Paths {
		// Check if the path contains any visited rooms
		if CheckCollision(path[1:len(path)-1], visited) {
			continue // Skip this path if it has visited rooms
		}

		// Mark rooms in this path as visited
		for _, room := range path[1 : len(path)-1] {
			visited[room] = true
		}

		// Add the path to the new list
		All_Paths_New = append(All_Paths_New, path)
	}

	/*
		MARKED FOR POTENTIAL CHANGE
	*/

	for _, Path := range All_Paths {
		if len(Path) == 2 {
			All_Paths_New = [][]string{}
			All_Paths_New = append(All_Paths_New, Path)
			return All_Paths_New
		}
	}

	return All_Paths_New
}
