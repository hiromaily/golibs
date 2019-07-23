package decimal

import (
	dc "github.com/ericlagergren/decimal"
	"math/big"
)

// Notes: about scale
// How do I interpret precision and scale of a number in a database?
// https://stackoverflow.com/questions/2377174/how-do-i-interpret-precision-and-scale-of-a-number-in-a-database
// ie 1234567.89 has a precision of 9
// ie 123456.789 has a scale of 3

type Decimal struct {
	ctx dc.Context
}

// if 12345.12345 is expected,
// New(10, 5, 0) //as precision 9,16,19,34
func New(prec, maxScale, minScale int) *Decimal {
	ctx := dc.Context{
		Precision:     prec,
		MaxScale:      maxScale,
		MinScale:      minScale,
		OperatingMode: dc.GDA,
	}

	return &Decimal{
		ctx: ctx,
	}
}

func (d *Decimal) String(v string) *dc.Big {
	//TODO: error should be checked though
	//b, _ := new(dc.Big).SetString(v)
	b, _ := dc.WithContext(d.ctx).SetString(v)
	return b
}

func (d *Decimal) Uint64(v uint64) *dc.Big {
	//return new(dc.Big).SetUint64(v)
	return dc.WithContext(d.ctx).SetUint64(v)
}

func (d *Decimal) Float64(v float64) *dc.Big {
	//return new(dc.Big).SetFloat64(v)
	return dc.WithContext(d.ctx).SetFloat64(v)
}

func (d *Decimal) BigFloat(v *big.Float) *dc.Big {
	//return new(dc.Big).SetFloat(v)
	return dc.WithContext(d.ctx).SetFloat(v)
}

func (d *Decimal) Add(x0, y0 *dc.Big) *dc.Big {
	x := dc.WithContext(d.ctx).Set(x0)
	y := dc.WithContext(d.ctx).Set(y0)
	dc.WithContext(d.ctx).Int64()

	return x.Add(x, y)
}

func (d *Decimal) Sub(x0, y0 *dc.Big) *dc.Big {
	x := dc.WithContext(d.ctx).Set(x0)
	y := dc.WithContext(d.ctx).Set(y0)

	return x.Sub(x, y)
}

func (d *Decimal) Mul(x0, y0 *dc.Big) *dc.Big {
	x := dc.WithContext(d.ctx).Set(x0)
	y := dc.WithContext(d.ctx).Set(y0)

	return x.Mul(x, y)
}

func (d *Decimal) Div(x0, y0 *dc.Big) *dc.Big {
	x := dc.WithContext(d.ctx).Set(x0)
	y := dc.WithContext(d.ctx).Set(y0)

	return x.Quo(x, y)
}
