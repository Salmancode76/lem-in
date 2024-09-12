package main

import (
	"fmt"
	"lem-in/pkg"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("YOU NEED TO PASS ONLY A FILENAME")
	}
	Colony := pkg.ReadFile()
	//fmt.Print(Colony)

	All_Paths := pkg.Edmonds_Karp(Colony.Rooms, Colony.Links)
	antArry := [][]int{}

	//fmt.Println("the ant number is", Colony.AntNum)
	for j := 1; j <= Colony.AntNum; j++ {
		antArry = append(antArry, []int{})
		antArry = append(antArry, []int{})
		//fmt.Println(len(antArry))
		for i, p := range All_Paths {
			if i == len(All_Paths)-1 {
				antArry[i] = append(antArry[i], j)
				break
			}
			a := (len(p) + len(antArry[i]))
			b := (len(All_Paths[i+1]) + len(antArry[i+1]))
			//fmt.Println(len(p), len(antArry[i]), "vs", len(All_Paths[i+1]), len(antArry[i+1]), "i is equal to ", i)
			if a <= b {
				antArry[i] = append(antArry[i], j)
				break
			}
		}
	}

	for i, a := range antArry {
		if i > len(All_Paths)-1 {
			break
		}
		fmt.Println("the ants", a, "will go down path", All_Paths[i])
	}
}
