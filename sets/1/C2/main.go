package main

import (
	"encoding/hex"
	"fmt"
)

func Xor(i []byte, o []byte) string {

	if len(i) != len(o) {
		return ""
	}

	result := make([]byte, len(i))

	for idx := range i {
		result[idx] = i[idx] ^ o[idx]
	}
	return hex.EncodeToString(result)
}

func main() {
	inputBytes, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		fmt.Println(err.Error())
	}
	testString, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(Xor(inputBytes, testString))

}
