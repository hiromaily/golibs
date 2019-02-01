package semaphore_test

import (
	. "github.com/hiromaily/golibs/semaphore"
	"testing"
)

func TestSemaphore(t *testing.T) {
	Semaphore(10)
}

func TestSemaphore2(t *testing.T) {
	Semaphore2(5)
}
