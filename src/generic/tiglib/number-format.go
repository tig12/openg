/** 
     Format a number with grouped thousands.
**/
package tiglib

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func NumberFormat(x interface{}, sep rune) string {
	switch x.(type) {
	case int:
		return numberFormatInt(x.(int), sep)
	case float64:
		return numberFormatFloat64(x.(float64), sep)
	default:
		panic(fmt.Sprintf("Unsupported type %T for NumberFormat", x))
	}
}

func numberFormatFloat64(x float64, sep rune) string {
	// note : does not use instructions like x - math.Trunc(x)
	// because sometimes supplementary digits are added
	s := strconv.FormatFloat(x, 'f', -1, 64)
	parts := strings.Split(s, ".")
	intPart, _ := strconv.Atoi(parts[0])
	var res = numberFormatInt(intPart, sep)
	if len(parts) == 2 {
		res += "." + parts[1]
	}
	return res
}

/**
    Adaptation from https://stackoverflow.com/questions/13020308/how-to-fmt-printf-an-integer-with-thousands-comma
**/
func numberFormatInt(n int, sep rune) string {

	s := strconv.Itoa(n)

	startOffset := 0
	var buff bytes.Buffer
	if n < 0 {
		startOffset = 1
		buff.WriteByte('-')
	}

	l := len(s)
	commaIndex := 3 - ((l - startOffset) % 3)
	if commaIndex == 3 {
		commaIndex = 0
	}

	for i := startOffset; i < l; i++ {
		if commaIndex == 3 {
			buff.WriteRune(sep)
			commaIndex = 0
		}
		commaIndex++
		buff.WriteByte(s[i])
	}

	return buff.String()
}
