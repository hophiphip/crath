package messagepack

import (
	"fmt"
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
		fmt.Println(testStr)
		fmt.Println(MessageToBigInt(testStr))
		fmt.Println(BigIntToMessage(MessageToBigInt(testStr)))

		if testStr != BigIntToMessage(MessageToBigInt(testStr)) {
			t.Error("For string", testStr,
				"failed to pack message",
			)
		}
	}
}
