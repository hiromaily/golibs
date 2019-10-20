package stdoutcapture

import (
	"bytes"
	"io"
	"os"
)

// CaptureStdout is not thread safe
func CaptureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// ContinuousCaptureStdout is not thread safe
func ContinuousCaptureStdout() string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	//f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
