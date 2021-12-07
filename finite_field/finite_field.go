package finitefield

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

// Only for field 2^8

const defaultPx = 0x1D // 00011101 -> a⁴ + a³ + a² + 1
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
			return "", fmt.Errorf("incorrect integer: %d, this is not supposed to happen", c)
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
			return "", fmt.Errorf("incorrect integer: %d, this is not supposed to happen", c)
		}
	}

	return buffer.String(), nil
}

// byteToSuperscript - converts byte to superscript string
func byteToPolynomial(n byte) (string, error) {
	var buffer bytes.Buffer
	wasSthPrinted := false
	nBitCount := 7

	for _, b := range fmt.Sprintf("%08b", n) {
		if b == '1' {
			log.Println(nBitCount)

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
						log.Println(str)
						buffer.WriteString(str)
					}
				}
			}

			wasSthPrinted = true
		}

		nBitCount--
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
			deg = (deg << 1) ^ defaultPx
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

// Product if elementis in GF(2^8) with p(x)
// Example p(x) = x^8 + x^4 + x^3 + x^2 + 1
// 				--> px = 0001_1101 --> 0b00011101
// mulG_2_8 - multiplication of elements in field GF(2^8)
func mulG_2_8(a, b, px byte) byte {
	deg := a
	res := byte(0)

	if (b & 1) != 0 {
		res = a
	}

	for i := 1; i < 8; i++ {
		if ((deg >> 7) & 1) != 0 {
			deg = (deg << 1) ^ px
		} else {
			deg <<= 1
		}

		if ((b >> i) & 1) != 0 {
			res ^= deg
		}
	}

	return res
}

func addG_2_8(a, b byte) byte {
	return a ^ b
}

// bitStringToByte - Simple temporary bit string to byte converter
func bitStringToByte(bitString string) (byte, error) {
	if len(bitString) != 8 {
		return 0, fmt.Errorf("incorrect string length, expected: %d, but got: %d", 8, len(bitString))
	}

	var res byte = 0

	for idx, bit := range bitString {
		switch bit {
		case '0':
			continue
		case '1':
			res |= byte(1) << (7 - idx)
		default:
			return 0, fmt.Errorf("incorrect string symbol on position: %d", idx)
		}
	}

	return res, nil
}
