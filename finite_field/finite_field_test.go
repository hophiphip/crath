package main

import (
	"math/big"
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

var intToSuperscriptTestCases = []intToString{
	{1234567890, "¹²³⁴⁵⁶⁷⁸⁹⁰"},
}

var bigToSuperscriptTestCases = []bigToString{
	{*big.NewInt(1234567890), "¹²³⁴⁵⁶⁷⁸⁹⁰"},
}

func TestIntToSuperscript(t *testing.T) {
}

func TestBigToSuperscript(t *testing.T) {
}

func TestBitsToPolynomial(t *testing.T) {
}
