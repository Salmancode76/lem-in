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
			log.Fatal("ERROR: invalid data format,Invalid number of ants in the file")
		}
		colony.AntNum = antNum
	} else {
		log.Fatal("ERROR: invalid data format, No data found in the file")
	}

	isEnd := false
	isStart := false
	var spawn string
	var target string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "##end" {
			isEnd = true
		}
		if line == "##start" {
			isStart = true
		}

		if (len(line) > 0 && (line[0] == '#' || line[0] == 'L')) || len(line) == 0 {
			continue
		}

		if isEnd {
			isEnd = false
			target = strings.Split(line, " ")[0]
			continue
		}
		if isStart {
			isStart = false
			spawn = strings.Split(line, " ")[0]
			continue
		}

		if strings.Contains(line, "-") {

			colony.Links = append(colony.Links, line)
		} else {
			colony.Rooms = append(colony.Rooms, strings.Split(line, " ")[0])
		}

	}
	if spawn != "" {
		colony.Rooms = append([]string{spawn}, colony.Rooms...) // Prepend spawn
	}
	colony.Rooms = append(colony.Rooms, target)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return colony
}
