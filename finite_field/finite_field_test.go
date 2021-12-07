package finitefield

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"testing"
)

type intToString struct {
	input  int
	output string
}

type bigToString struct {
	input  big.Int
	output string
}

type field_G_2_8_AddCase struct {
	a, b, output byte
}

type field_G_2_8_MulCase struct {
	a, b, px, output byte
}

var intToSuperscriptTestCases = []intToString{
	{1234567890, "¹²³⁴⁵⁶⁷⁸⁹⁰"},
}

var bigToSuperscriptTestCases = []bigToString{
	{*big.NewInt(1234567890), "¹²³⁴⁵⁶⁷⁸⁹⁰"},
}

var bigBitsToPolynomialTestCases = []bigToString{
	{*big.NewInt(0xeeff), "x¹⁵ + x¹⁴ + x¹³ + x¹¹ + x¹⁰ + x⁹ + x⁷ + x⁶ + x⁵ + x⁴ + x³ + x² + x + 1"},
	{*big.NewInt(0xffff), "x¹⁵ + x¹⁴ + x¹³ + x¹² + x¹¹ + x¹⁰ + x⁹ + x⁸ + x⁷ + x⁶ + x⁵ + x⁴ + x³ + x² + x + 1"},
}

var fieldAddOperationTestCases = []field_G_2_8_AddCase{
	{1, 2, 1 ^ 2},
	{3, 2, 3 ^ 2},
	{5, 22, 5 ^ 22},
}

var fieldMulOperationTestCases = []field_G_2_8_MulCase{
	{0b0000_1011, 0b0000_1011, 0b0001_0111, 69},
	{0b1111_1111, 0b1111_1111, 0b0001_0111, 242},
	{0b1000_1011, 0b1000_1011, 0b0001_0111, 206},
	{0b0100_1011, 0b0000_1011, 0b0001_0111, 171},
}

func TestIntToSuperscript(t *testing.T) {
	for _, testCase := range intToSuperscriptTestCases {
		if res, err := intToSuperscript(testCase.input); err != nil {
			t.Error(
				"For value: ", testCase.input,
				" received error: ", err,
			)
		} else {
			if res != testCase.output {
				t.Error(
					"For value: ", testCase.input,
					" expected: ", testCase.output,
					" but got: ", res,
				)
			}
		}
	}
}

func TestBigToSuperscript(t *testing.T) {
	for _, testCase := range bigToSuperscriptTestCases {
		if res, err := bigToSuperscript(testCase.input); err != nil {
			t.Error(
				"For value: ", testCase.input.String(),
				" received error: ", err,
			)
		} else {
			if res != testCase.output {
				t.Error(
					"For value: ", testCase.input.String(),
					" expected: ", testCase.output,
					" but got: ", res,
				)
			}
		}
	}
}

func TestBitsToPolynomial(t *testing.T) {
	for _, testCase := range bigBitsToPolynomialTestCases {
		if res, err := bigBitsToPolynomial(testCase.input); err != nil {
			t.Error(
				"For value: ", testCase.input.String(),
				" received error: ", err,
			)
		} else {
			if res != testCase.output {
				t.Error(
					"For value: ", testCase.input.String(),
					" expected: ", testCase.output,
					" but got: ", res,
				)
			}
		}
	}
}

func TestAddG_2_8(t *testing.T) {
	for _, testCase := range fieldAddOperationTestCases {
		if res := AddG_2_8(testCase.a, testCase.b); res != testCase.output {
			t.Error(
				"For values: ", testCase.a,
				" and: ", testCase.b,
				" expected: ", testCase.output,
				" but got: ", res,
			)
		}
	}
}

func TestMulG_2_8(t *testing.T) {
	for _, testCase := range fieldMulOperationTestCases {
		if res := MulG_2_8(testCase.a, testCase.b, testCase.px); res != testCase.output {
			t.Error(
				"For values: ", testCase.a,
				" and: ", testCase.b,
				" , and field: ", testCase.px,
				" expected: ", testCase.output,
				" but got: ", res,
			)
		}
	}
}

func TestField(t *testing.T) {
	testAddFile := "test.add.txt"
	testMulFile := "test.mul.txt"

	px := byte(0b0001_0111)

	fa, err := os.Create(testAddFile)
	if err != nil {
		t.Error(err)
	}

	wa := bufio.NewWriter(fa)

	for a := byte(0b0000_0000); ; a++ {

		for b := byte(0b0000_0000); ; b++ {
			wa.WriteString(fmt.Sprintf("%08b", AddG_2_8(a, b)))
			wa.WriteString("\t")

			if b == byte(0b1111_1111) {
				wa.WriteString("\n")
				break
			}
		}

		if a == byte(0b1111_1111) {
			wa.WriteString("\n")
			break
		}
	}

	fm, err := os.Create(testMulFile)
	if err != nil {
		t.Error(err)
	}

	wm := bufio.NewWriter(fm)

	for a := byte(0b0000_0000); ; a++ {

		for b := byte(0b0000_0000); ; b++ {
			wm.WriteString(fmt.Sprintf("%08b", MulG_2_8(a, b, px)))
			wm.WriteString("\t")

			if b == byte(0b1111_1111) {
				wm.WriteString("\n")
				break
			}
		}

		if a == byte(0b1111_1111) {
			wm.WriteString("\n")
			break
		}
	}

	defer func() {
		if err := fa.Close(); err != nil {
			t.Error(err)
		}

		if err := fm.Close(); err != nil {
			t.Error(err)
		}
	}()
}

func TestBitStringToByte(t *testing.T) {
	for num := byte(0); ; num++ {
		input := fmt.Sprintf("%08b", num)

		if res, err := BitStringToByte(input); err != nil {
			t.Error(
				"For input: ", input,
				" got an error: ", err,
			)
		} else {
			if res != num {
				t.Error(
					"For input: ", input,
					" expected: ", num,
					" but got: ", res,
				)
			}
		}

		if num == byte(0b1111_1111) {
			break
		}
	}

	for _, input := range []string{"0000", "0ada", "11111111111", "adasdjasl", "asdasdasd", "0000111a", ""} {
		if res, err := BitStringToByte(input); err == nil {
			t.Error(
				"For input: ", input,
				" expected error but got: ", res,
			)
		}
	}
}
