package decimal_test

import (
	"fmt"
	"testing"

	dc "github.com/ericlagergren/decimal"
)

//https://github.com/ericlagergren/decimal
//https://github.com/ericlagergren/decimal/blob/master/example_decimal_test.go

func TestDecimalString(t *testing.T) {
	var tests = []struct {
		strDecimal  string
		format      string
		description string
	}{
		{"123456789.12345678901234567890", "%s", ""},
		{"123456789.12345678901234567890", "%.1f", "decimal point but no exponent, e.g. 123.456"},
		{"123456789.12345678901234567890", "%10.f", ""},
		{"123456789.12345678901234567890", "%10.10f", ""},
		{"123456789.12345678901234567890", "%10.20f", ""},
		//{"123456789.123456789",  "%.3g", "%e for large exponents, %f otherwise. Precision is discussed below."},
		//{"123456789.123456789",  "%6.4g", ""},
	}

	for _, val := range tests {
		d, _ := new(dc.Big).SetString(val.strDecimal)
		fmt.Printf(val.format+"\n", d)
	}
}
