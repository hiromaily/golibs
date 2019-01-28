package stdout_capture_test

import (
	"fmt"
	"testing"

	. "github.com/hiromaily/golibs/stdout_capture"
)

func TestCaptureStdout(t *testing.T) {
	f := func() {
		fmt.Println("test")
	}
	captured := CaptureStdout(f)
	fmt.Printf("captured: %s\n", captured)
}
