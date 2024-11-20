package main

import (
	"encoding/hex"
	"fmt"
)

const base64Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Base64Encode(src []byte) []byte {

	length := len(src)
	// Not sure how to calculate length here.
	dst := make([]byte, length*2)
	if len(src) == 0 {
		return nil
	}
	// enc is a pointer receiver, so the use of enc.encode within the hot
	// loop below means a nil check at every operation. Lift that nil check
	// outside of the loop to speed up the encoder.

	di, si := 0, 0
	n := (len(src) / 3) * 3
	for si < n {
		// Convert 3x 8bit source bytes into 4 bytes
		val := uint(src[si+0])<<16 | uint(src[si+1])<<8 | uint(src[si+2])

		dst[di+0] = base64Alphabet[val>>18&0x3F]
		dst[di+1] = base64Alphabet[val>>12&0x3F]
		dst[di+2] = base64Alphabet[val>>6&0x3F]
		dst[di+3] = base64Alphabet[val&0x3F]

		si += 3
		di += 4
	}

	remain := len(src) - si
	if remain == 0 {
		return dst
	}
	// Add the remaining small block
	val := uint(src[si+0]) << 16
	if remain == 2 {
		val |= uint(src[si+1]) << 8
	}

	dst[di+0] = base64Alphabet[val>>18&0x3F]
	dst[di+1] = base64Alphabet[val>>12&0x3F]

	switch remain {
	case 2:
		dst[di+2] = base64Alphabet[val>>6&0x3F]
		dst[di+3] = byte('=')
	case 1:
		dst[di+2] = byte('=')
		dst[di+3] = byte('=')
	}
	return dst
}

func main() {
	inputBytes, err := hex.DecodeString("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		fmt.Println(err.Error())
	}
	enc := Base64Encode(inputBytes)
	fmt.Println(string(enc))
}
