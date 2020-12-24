package messagepack

import (
	"testing"
)

var (
	convertTestValues = []string{
		"aah",
		"bobby",
		"coccyx",
		"diddled",
		"epee",
		"faff",
		"Shh",
	}
)

func TestMessageConvert(t *testing.T) {
	for _, testStr := range convertTestValues {
		if testStr != BigIntToString(StringToBigInt(testStr)) {
			t.Error("For string", testStr,
				"failed to pack message",
			)
		}
	}
}
