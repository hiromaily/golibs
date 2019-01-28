package stdout_capture_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/hiromaily/golibs/stdout_capture"
)

func TestCaptureStdout(t *testing.T) {
	f := func() {
		fmt.Println("test")
	}
	captured := CaptureStdout(f)
	fmt.Printf("captured: %s\n", captured)
}

// FIXME:このやり方では全部captureできるわけではない
func TestCaptureStdout2(t *testing.T) {
	go continuousRun()

	for i := 0; i < 10; i++ {
		fmt.Printf("test%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func continuousRun() {
	for {
		captured := ContinuousCaptureStdout()
		if captured != "" {
			fmt.Printf("captured: %s\n", captured)
		}
	}
}
