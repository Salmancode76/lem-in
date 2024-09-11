package main

import (
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

	pkg.Edmonds_Karp(Colony.Rooms, Colony.Links)

}
