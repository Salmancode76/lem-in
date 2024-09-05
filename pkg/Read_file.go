package pkg

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Colony struct holds the number of ants, room names, and links.
type Colony struct {
	AntNum int      // Number of ants
	Rooms  []string // List of room names
	Links  []string // List of links
}

// readFile reads the input file and populates the Colony struct.
func ReadFile() Colony {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	colony := Colony{}

	scanner := bufio.NewScanner(file)

	// Read the first line to get the number of ants
	if scanner.Scan() {
		antNum, err := strconv.Atoi(scanner.Text())
		if err != nil || antNum <= 0 {
			log.Fatal("Invalid number of ants in the file")
		}
		colony.AntNum = antNum
	} else {
		log.Fatal("No data found in the file")
	}

	isRoom := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check for the start and end of the room section
		switch line {
		case "##start":
			isRoom = true
		case "##end":
			isRoom = false
			if scanner.Scan() { // Capture the first line after ##end
				colony.Rooms = append(colony.Rooms, strings.Split(scanner.Text(), " ")[0])
			}
		default:
			// Skip comments and empty lines
			if line == "" || line[0] == '#' {
				continue
			}
			// Capture room names or links
			if isRoom {
				colony.Rooms = append(colony.Rooms, strings.Split(line, " ")[0])
			} else {
				colony.Links = append(colony.Links, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return colony
}
