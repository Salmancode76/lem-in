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
	if len(os.Args) != 2 {
		log.Fatal("YOU NEED TO PASS ONLY A FILENAME")
	}
	Colony := pkg.ReadFile()

	fmt.Println(Colony)

	All_Paths := pkg.Edmonds_Karp(Colony.Rooms, Colony.Links)

	//fmt.Println(All_Paths)
	All_Paths = (pkg.RedundantPaths(All_Paths))

	fmt.Println(All_Paths)

	pkg.Ants_Traffic(Colony.AntNum, All_Paths)

	//fmt.Println(AntsPaths)

}
