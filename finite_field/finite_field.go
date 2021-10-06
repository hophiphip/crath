package main

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

var numberToSuperscript = map[rune]string{
	'0': "⁰",
	'1': "¹",
	'2': "²",
	'3': "³",
	'4': "⁴",
	'5': "⁵",
	'6': "⁶",
	'7': "⁷",
	'8': "⁸",
	'9': "⁹",
}

// bigToSuperscript - converts big.Int to superscript string
func bigToSuperscript(n big.Int) (string, error) {
	var buffer bytes.Buffer

	str := n.String()
	for _, c := range str {
		if val, isIn := numberToSuperscript[c]; isIn {
			buffer.WriteString(val)
		} else {
			return "", fmt.Errorf("Incorrect integer: %d, this is not supposed to happen", c)
		}
	}

	return buffer.String(), nil
}

// intToSuperscript - converts int to superscript string
func intToSuperscript(n int) (string, error) {
	var buffer bytes.Buffer

	for _, c := range strconv.Itoa(n) {
		if val, isIn := numberToSuperscript[c]; isIn {
			buffer.WriteString(val)
		} else {
			return "", fmt.Errorf("Incorrect integer: %d, this is not supposed to happen", c)
		}
	}

	return buffer.String(), nil
}

// bigBitsToPolynomial - converts big.Int bits to polynomial representation string
func bitsToPolynomial(n big.Int) (string, error) {
	var buffer bytes.Buffer

	nBytes := n.Bytes()
	nBitCount := len(nBytes)*8 - 1
	wasSthPrinted := false

	for _, nByte := range nBytes {
		for _, b := range fmt.Sprintf("%08b", nByte) {
			if b == '1' {
				if wasSthPrinted {
					buffer.WriteString(" + ")
				}

				if nBitCount == 0 {
					buffer.WriteString("1")
				} else {
					buffer.WriteString("x")
					if nBitCount != 1 {
						if str, err := intToSuperscript(nBitCount); err != nil {
							return "", err
						} else {
							buffer.WriteString(str)
						}
					}
				}

				wasSthPrinted = true
			}

			nBitCount--
		}
	}

	return buffer.String(), nil
}

func main() {
	if str, err := bigToSuperscript(*big.NewInt(1234567890)); err != nil {
		log.Println(err)
	} else {
		log.Println(str)
	}

	if str, err := intToSuperscript(1234567890); err != nil {
		log.Println(err)
	} else {
		log.Println(str)
	}

	if str, err := bitsToPolynomial(*big.NewInt(0xeeff)); err != nil {
		log.Println(err)
	} else {
		log.Println(str)
	}
}
