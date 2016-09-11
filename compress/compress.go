package compress

import (
	"bytes"
	"compress/gzip"
)

// GZipString is compress string by GZip
func GZipString(baseStr string) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write([]byte(baseStr)); err != nil {
		return []byte{}, err
	}
	if err := gz.Flush(); err != nil {
		return []byte{}, err
	}
	if err := gz.Close(); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}
