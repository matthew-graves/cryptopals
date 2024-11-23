package main

import (
	"encoding/hex"
	"fmt"
)

func XorDecode(i []byte, d byte) {

	result := make([]byte, len(i))

	counts := make(map[byte]int)

	for idx := range i {
		result[idx] = i[idx] ^ d
		counts[result[idx]]++
	}

	fmt.Println(string(result))
	fmt.Println(counts)

}

func main() {
	inputString, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	if err != nil {
		fmt.Println("InputString Invalid Hex")
	}
	for i := 0; i < 256; i++ {
		fmt.Printf("i: %d, s: ", i)
		XorDecode(inputString, byte(i))
	}

}
