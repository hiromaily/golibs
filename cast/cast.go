// Package cast is just sample
package cast

import (
	"bytes"
)

//http://qiita.com/ikawaha/items/2d58f58e4ab12918e8c9
//http://qiita.com/mattn/items/176459728ff4f854b165

// StoB is to cast string to []byte
// performance is not good because of copy whole memory
func StoB(str string) []byte {
	return []byte(str)
}

// BtoS is to cast []byte to string
func BtoS(bt []byte) string {
	return string(bt)
}

// BufferStoB is to cast string to []byte using bytes.NewBuffer()
func BufferStoB(str string) []byte {
	b := bytes.NewBuffer(make([]byte, 0, 100))
	b.WriteString(str)
	//_ = b.String()
	return b.Bytes()
}

// BufferBtoS is to cast []byte to string using bytes.NewBuffer()
func BufferBtoS(bt []byte) string {
	b := bytes.NewBuffer(make([]byte, 0, 100))
	for i := 0; i < len(bt); i++ {
		b.WriteByte(bt[i])
	}
	return b.String()
}
