package finitefield

import (
	"bufio"
	"fmt"
	"math"
	"math.io/crath/mulfunc"
	ord2 "math.io/crath/ord"
	"math/big"
	"os"
	"strings"
	"testing"
)

// Constants for addition, multiplication, order tables

// Field modulus
const px = byte(0b0001_0111)
const testAddFile = "test.add.txt"
const testMulFile = "test.mul.txt"
const testOrdFile = "test.ord.txt"
const lineLen = 80

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

func TestPrintAdditionTable(t *testing.T) {
	fa, err := os.Create(testAddFile)
	if err != nil {
		t.Error(err)
	}

	wa := bufio.NewWriter(fa)

	for a := byte(0b0000_0000); ; a++ {

		for b := byte(0b0000_0000); ; b++ {
			result := AddG_2_8(a, b)

			_, _ = wa.WriteString(fmt.Sprintf("%08b", result))
			_, _ = wa.WriteString("\t")

			if b == byte(0b1111_1111) {
				_, _ = wa.WriteString("\n")
				break
			}
		}

		if a == byte(0b1111_1111) {
			_, _ = wa.WriteString("\n")
			break
		}
	}

	// Cleanup
	defer func() {
		if err := wa.Flush(); err != nil {
			t.Error(err)
		}

		if err := fa.Close(); err != nil {
			t.Error(err)
		}
	}()
}

func TestPrintMultiplicationTable(t *testing.T) {
	// Multiplication table
	fm, err := os.Create(testMulFile)
	if err != nil {
		t.Error(err)
	}

	wm := bufio.NewWriter(fm)

	for a := byte(0b0000_0000); ; a++ {

		for b := byte(0b0000_0000); ; b++ {
			_, _ = wm.WriteString(fmt.Sprintf("%08b", MulG_2_8(a, b, px)))
			_, _ = wm.WriteString("\t")

			if b == byte(0b1111_1111) {
				_, _ = wm.WriteString("\n")
				break
			}
		}

		if a == byte(0b1111_1111) {
			_, _ = wm.WriteString("\n")
			break
		}
	}

	// Cleanup
	defer func() {
		if err := wm.Flush(); err != nil {
			t.Error(err)
		}

		if err := fm.Close(); err != nil {
			t.Error(err)
		}
	}()
}

func TestPrintElementOrderTable(t *testing.T) {
	// Primitive element order
	primitiveOrder := big.NewInt(int64(px - 1))
	var primitiveElements []byte

	amountOfPrimitiveElements := mulfunc.Euler(big.NewInt(int64(math.Pow(2, 8)) - 1))

	// Elements order table
	fo, err := os.Create(testOrdFile)
	if err != nil {
		t.Error(err)
	}

	wo := bufio.NewWriter(fo)

	pxStr, err := bigBitsToPolynomial(*big.NewInt(int64(px)))
	if err != nil {
		pxStr = ""
	}

	_, err = wo.WriteString(fmt.Sprintf("PX Modulo: %40.40s \n", pxStr))
	if err != nil {
		t.Error(err)
	}

	_, err = wo.WriteString(fmt.Sprintf("Amount of primitive elements: %9.9s \n\n", amountOfPrimitiveElements.String()))
	if err != nil {
		t.Error(err)
	}

	_, err = wo.WriteString(fmt.Sprintf("%7.7s | %40.40s | %8.8s | %5.5s\n", "Regular", "Polynomial", "Vector", "Order"))
	if err != nil {
		t.Error(err)
	}

	// Line separator
	_, err = wo.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", lineLen)))
	if err != nil {
		t.Error(err)
	}

	for a := byte(0b0000_0000); ; a++ {
		order, err := ord2.Ord(big.NewInt(int64(a)), big.NewInt(int64(px)))
		ordStr := "0"

		if err == nil {
			ordStr = order.String()

			// Find all primitive elements
			if order.Cmp(primitiveOrder) == 0 {
				primitiveElements = append(primitiveElements, a)
			}
		}

		polynomial, err := bigBitsToPolynomial(*big.NewInt(int64(a)))
		if err != nil {
			polynomial = ""
		}

		_, err = wo.WriteString(fmt.Sprintf("%7.7d | %40.40s | %08b | %5.5s\n", int64(a), polynomial, a, ordStr))
		if err != nil {
			t.Error(err)
		}

		// Line separator
		_, err = wo.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", lineLen)))
		if err != nil {
			t.Error(err)
		}

		if a == byte(0b1111_1111) {
			_, err = wo.WriteString("\n")
			if err != nil {
				t.Error(err)
			}

			break
		}
	}

	// Print primitive elements
	if _, err := wo.WriteString(fmt.Sprintf("(%d) Primitive elements: \n", len(primitiveElements))); err != nil {
		t.Error(err)
	}

	for index, primitive := range primitiveElements {
		if _, err := wo.WriteString(fmt.Sprintf("%3.3d: %d\n", index, primitive)); err != nil {
			t.Error(err)
		}
	}

	// Cleanup
	defer func() {
		// Don't forget to flush C:
		if err := wo.Flush(); err != nil {
			t.Error(err)
		}

		if err := fo.Close(); err != nil {
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
