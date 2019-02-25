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

func TestConvert(t *testing.T) {
	// 1 wei is 1,000,000,000,000,000,000 wei

	var tests = []struct {
		amount   string
		expected string
	}{
		{"1", "1000000000000000000"},
		{"0.35", "350000000000000000"},
		{"1.35", "1350000000000000000"},
		{"21.35", "21350000000000000000"},
		{"21.356", "21356000000000000000"},
		{"321.35", "321350000000000000000"},     //decimalでは、この桁からうまく返せない。。。
		{"99999.35", "99999350000000000000000"}, //
	}

	for _, val := range tests {
		d, _ := new(dc.Big).SetString(val.amount)
		fmt.Printf("%30f\n", d.Quantize(-30))
	}

}