package semaphore_test

import (
	"testing"

	. "github.com/hiromaily/golibs/semaphore"
)

func TestSemaphore(t *testing.T) {
	Semaphore(10)
}

func TestSemaphore2(t *testing.T) {
	Semaphore2(5)
}
