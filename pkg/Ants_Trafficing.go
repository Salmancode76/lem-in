package pkg

import (
	"fmt"
	"strconv"
)

type Ants_Path struct {
	Ants  [][]int
	Paths [][]string
}

func Ants_Traffic(Num_Ants int, All_Paths [][]string) {
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
	antQ := 0
	// Loop to collect paths and visited rooms
	for i, a := range antArry {
		if len(a) == 0 {
			continue
		}
		antQ++
		All_Paths[i] = All_Paths[i][1:]
		fmt.Println("The ants", a, "will go down path", All_Paths[i])

		antsPath.Ants = append(antsPath.Ants, a)
		antsPath.Paths = append(antsPath.Paths, All_Paths[i])

	}

	outPath := []string{}
	for m := len(All_Paths[0]) + len(antArry[0])*2; m > 0; m-- {
		outPath = append(outPath, "")
	}
	for z := 0; z < len(antArry); z++ {
		w := 0
		line := 0
		for i := 0; i <= len(All_Paths[z])+w; i++ {

			for j := 0; j <= w; j++ {
				if i-j < len(All_Paths[z]) && (len(antArry[z]) > 0) {
					outPath[line] = outPath[line] + " L" + strconv.Itoa(antArry[z][j]) + "-" + All_Paths[z][i-j]
				}

			}
			if w < len(antArry[z])-1 {
				w++
				//println(w)
			}

			line++
		}
	}
	var moves int
	for i := 0; len(outPath) > i; i++ {

		if outPath[i] != "" {
			moves++

			fmt.Println(outPath[i][1:])
		}
	}
}
