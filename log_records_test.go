package main

import (
	"os"
	"testing"
)

func TestProcessFile(t *testing.T) {

	test := []struct {
		filename        string
		testData        []byte
		tagMap          map[string]int
		portProtocolMap map[string]int
		expectedError   error
	}{
		{
			filename: "valid test input",
			testData: []byte(`80,tcp,web
53,udp,dns
22,tcp,ssh
80,tcp
`),
			tagMap: map[string]int{
				"web":      1,
				"dns":      1,
				"ssh":      1,
				"untagged": 1,
			},
			portProtocolMap: map[string]int{
				"80,tcp": 2,
				"53,udp": 1,
				"22,tcp": 1,
			},
		},
		{
			filename:        "valid test input-empty file",
			testData:        []byte(``),
			tagMap:          map[string]int{},
			portProtocolMap: map[string]int{},
		},
	}

	for _, tt := range test {
		// Create a temporary file to simulate the input file
		tmpfile, err := os.CreateTemp("", "testfile")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name()) // clean up
		if _, err := tmpfile.Write([]byte(tt.testData)); err != nil {
			t.Fatal(err)
		}
		if err := tmpfile.Close(); err != nil {
			t.Fatal(err)
		}

		// Call ProcessFile function with the temp file
		tagMap, portProtocolMap, err := ProcessFile(tmpfile.Name())
		if err != nil {
			t.Fatalf("ProcessFile returned an error: %v", err)
		}

		// Validate tagMap
		for k, v := range tt.tagMap {
			if tagMap[k] != v {
				t.Errorf("tagMap[%s] = %d; want %d", k, tagMap[k], v)
			}
		}

		for k, v := range tt.portProtocolMap {
			if portProtocolMap[k] != v {
				t.Errorf("portProtocolMap[%s] = %d; want %d", k, portProtocolMap[k], v)
			}
		}
	}
}
