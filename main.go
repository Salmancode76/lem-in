package main

import (
	"lem-in/pkg"
	"log"
	"os"
)

type Ants_Path struct {
	Ants  [][]int
	Paths [][]string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("YOU NEED TO PASS ONLY A FILENAME")
	}
	Colony := pkg.ReadFile()

	All_Paths := pkg.Edmonds_Karp(Colony.Rooms, Colony.Links)

	All_Paths = (pkg.RedundantPaths(All_Paths))

	if len(All_Paths) <= 0 {
		log.Fatal("ERROR: invalid data format, PATHS ARE INVALID")
	}

	pkg.Ants_Traffic(Colony.AntNum, All_Paths)

	//fmt.Println(AntsPaths)

}
