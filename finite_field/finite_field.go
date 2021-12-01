package main

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

const px = 0x1D // 00011101 -> a⁴ + a³ + a² + 1
// a⁸ = a⁴ + a³ + a² + 1

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
func bigBitsToPolynomial(n big.Int) (string, error) {
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

func fieldAdd(a, b *big.Int) big.Int {
	return *big.NewInt(0).Xor(a, b)
}

func fieldMul8Aplha(b *big.Int) (big.Int, error) {
	ret := big.NewInt(0)

	bBytes := b.Bytes()

	if len(bBytes) > 1 {
		return *ret, fmt.Errorf("element must be from field: 2⁸: %s", b.String())
	}

	if ((bBytes[0] >> 7) & 1) == 1 {
		ret.SetBytes([]byte{(bBytes[0] << 1) ^ 0x1D})
	} else {
		ret.SetBytes([]byte{bBytes[0] << 1})
	}

	return *ret, nil
}

func fieldMul8(a, b *big.Int) (big.Int, error) {
	ret := big.NewInt(0)

	aBytes := a.Bytes()
	bBytes := b.Bytes()

	if len(bBytes) > 1 || len(aBytes) > 1 {
		return *ret, fmt.Errorf("elements must be from field: 2⁸")
	}

	deg := aBytes[0]
	byteRes := byte(0)

	if bBytes[0]&1 == 1 {
		byteRes = aBytes[0]
	}

	for i := 1; i < 8; i++ {
		if ((deg >> 7) & 1) == 1 {
			deg = (deg << 1) ^ px
		} else {
			deg <<= 1
		}

		if (bBytes[0]>>i)&1 == 1 {
			byteRes ^= deg
		}
	}

	ret.SetBytes([]byte{byteRes})

	return *ret, nil
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

	if str, err := bigBitsToPolynomial(*big.NewInt(0xeeff)); err != nil {
		log.Println(err)
	} else {
		log.Println(str)
	}

	if mul, err := fieldMul8Aplha(big.NewInt(0xaa)); err != nil {
		log.Println(err)
	} else {
		if str, errS := bigBitsToPolynomial(mul); errS != nil {
			log.Println(errS)
		} else {
			log.Println(str)
		}
	}

	if mul, err := fieldMul8Aplha(big.NewInt(0xaabb)); err != nil {
		log.Println(err)
	} else {
		log.Println(bigBitsToPolynomial(mul))
	}

	if mul, err := fieldMul8(big.NewInt(0b0000111), big.NewInt(0b1001000)); err != nil {
		log.Println(err)
	} else {
		if str, errS := bigBitsToPolynomial(mul); errS != nil {
			log.Println(errS)
		} else {
			log.Println(str)
		}
	}
}
