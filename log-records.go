package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Function to process file and return tagMap and portProtocolMap
func ProcessFile(filename string) (map[string]int, map[string]int, error) {
	tagMap := make(map[string]int)
	portProtocolMap := make(map[string]int)

	lookups, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer lookups.Close()

	lookupReader := bufio.NewReader(lookups)
	for {
		line, err := lookupReader.ReadString('\n')
		if err != nil {
			break
		}

		trimmedS := strings.TrimSpace(line)
		splitLine := strings.Split(trimmedS, ",")
		var port, protocol, tag string

		port = strings.ToLower(strings.TrimSpace(splitLine[0]))
		protocol = strings.ToLower(strings.TrimSpace(splitLine[1]))

		if len(splitLine) == 2 {
			tag = "untagged"
		}

		if len(splitLine) == 3 {
			tag = strings.ToLower(strings.TrimSpace(splitLine[2]))
		}

		if tag != "" {
			tagMap[tag]++
		}

		key := port + "," + protocol
		if key != "," {
			portProtocolMap[key]++
		}
	}

	return tagMap, portProtocolMap, nil
}

// Main function calls the ProcessFile function
func main() {
	start := time.Now()
	args := os.Args
	filename := args[1]

	tagMap, portProtocolMap, err := ProcessFile(filename)
	if err != nil {
		fmt.Printf("unable to open file %s: %v", filename, err)
		return
	}

	fmt.Println("Count of matches for each tag: ", tagMap)
	fmt.Println("Count of matches for each port/protocol combination: ", portProtocolMap)

	fmt.Println("Time elapsed: ", time.Until(start))
}
