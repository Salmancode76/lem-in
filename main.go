package main

import (
	"fmt"
	"lem-in/pkg"
	"log"
	"os"
)

type Ants_Path struct {
	Ants  [][]int
	Paths [][]string
}

func main() {
	steps := []int{}
	// var min int

	if len(os.Args) != 2 {
		log.Fatal("YOU NEED TO PASS ONLY A FILENAME")
	}
	Colony := pkg.ReadFile()

	fmt.Println(Colony)

	All_Paths := pkg.Edmonds_Karp(Colony.Rooms, Colony.Links)

	for i := 0; i < len(All_Paths); i++ {

		TAll_Paths := (pkg.RedundantPaths(All_Paths[i:]))
		steps = append(steps, pkg.Count(Colony.AntNum, TAll_Paths))
	}

	// fmt.Println("the steps for each path are =", steps)
	min := steps[0]
	index := 0
	for i := 0; i < len(steps); i++ {
		if steps[i] < min {
			min = steps[i]
			index = i
		}
	}

	All_Paths = (pkg.RedundantPaths(All_Paths[index:]))

	if len(All_Paths) <= 0 {
		log.Fatal("ERROR: invalid data format, PATHS ARE INVALID")
	}

	fmt.Println(All_Paths)

	pkg.Ants_Traffic(Colony.AntNum, All_Paths)

}
