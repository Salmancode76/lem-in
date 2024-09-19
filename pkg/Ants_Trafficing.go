package pkg

import "fmt"

type Ants_Path struct {
	Ants  [][]int
	Paths [][]string
}

func Ants_Traffic(Num_Ants int, All_Paths [][]string) Ants_Path {
	var antsPath Ants_Path

	antArry := make([][]int, len(All_Paths))
	// Loop over each ant
	for j := 1; j <= Num_Ants; j++ {
		// Loop over each path
		for i, p := range All_Paths {
			// Assign ant to the path if it's not the last path
			if i == len(All_Paths)-1 {
				antArry[i] = append(antArry[i], j)
				// Mark all intermediate rooms of this path as visited
				break
			}

			// Calculate path lengths and assign ants
			a := (len(p) + len(antArry[i]))
			b := (len(All_Paths[i+1]) + len(antArry[i+1]))

			if a <= b {
				antArry[i] = append(antArry[i], j)
				break
			}
		}
	}

	// Loop to collect paths and visited rooms
	for i, a := range antArry {
		if len(a) == 0 {
			continue
		}

		fmt.Println("The ants", a, "will go down path", All_Paths[i])

		antsPath.Ants = append(antsPath.Ants, a)
		antsPath.Paths = append(antsPath.Paths, All_Paths[i])
	}
	return antsPath
}
