package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"sync"
)

func readFile(path string) []string {
	readLines := []string{}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Unable to read file")
		return []string{""}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		readLines = append(readLines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return readLines
}

const commonChars string = "EeTtAaOoIiNn SsHhRrDdLlUu"

func generateValidAsciiBytes() []byte {
	var asciiBytes []byte
	for i := 32; i <= 126; i++ { // Printable ASCII characters
		asciiBytes = append(asciiBytes, byte(i))
	}

	// Add commonly used control characters
	controlChars := []byte{9, 10, 13} // Tab, Newline, Carriage Return
	asciiBytes = append(asciiBytes, controlChars...)

	return asciiBytes
}

func GetXorCommonChars(i []byte) []byte {

	byteCount := getMaxByte(i)
	commonCharB := []byte(commonChars)
	xorBytes := make([]byte, len(commonCharB)+1)

	for idx := range commonCharB {
		xorBytes[idx] = byteCount ^ commonCharB[idx]
	}

	return xorBytes

}

func getMaxByte(i []byte) byte {
	counts := make(map[byte]int)

	for idx := range i {
		counts[i[idx]]++
	}

	var maxByte byte
	var maxCount int
	for k, v := range counts {
		if v > maxCount {
			maxByte = k
			maxCount = v
		}
	}

	return maxByte

}

func XorDecode(i []byte) {

	d := GetXorCommonChars(i)
	validAscii := generateValidAsciiBytes()

	result := make([]byte, len(i))
	countNil := len(i)

	for dIdx := range d {
		// Try xor each byte against d
		for idx := range i {

			if !(bytes.Contains(validAscii, []byte{i[idx] ^ d[dIdx]})) {
				break
			}
			countNil--
			result[idx] = i[idx] ^ d[dIdx]
		}

		if countNil == 0 {
			fmt.Println(string(result))
		}

		countNil = len(i)
		result = make([]byte, len(i))
	}
}

func main() {
	var wg sync.WaitGroup

	lines := readFile("inputs.txt")
	for _, line := range lines {
		wg.Add(1)
		line, err := hex.DecodeString(line)
		if err != nil {
			fmt.Println("Unable to decode hex")
			continue
		}
		go func(line []byte) {
			defer wg.Done()
			XorDecode(line)
		}(line)
	}
	wg.Wait()
}
