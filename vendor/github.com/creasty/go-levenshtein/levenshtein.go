package levenshtein

/*
#cgo CFLAGS: -I .

#include <stdint.h>
#include <stddef.h>

#include "levenshtein.h"
*/
import "C"

import (
	"unsafe"
)

func Distance(a, b string) int {
	if a == b {
		return 0
	}

	aRune := []rune(a)
	bRune := []rune(b)
	aLen := len(aRune)
	bLen := len(bRune)

	if aLen == 0 {
		return bLen
	}
	if bLen == 0 {
		return aLen
	}

	aPtr := (*C.int32_t)(unsafe.Pointer(&aRune[0]))
	bPtr := (*C.int32_t)(unsafe.Pointer(&bRune[0]))

	cdist := C.levenshtein(aPtr, C.size_t(aLen), bPtr, C.size_t(bLen))
	return int(cdist)
}

func LcsLen(a, b string) int {
	aRune := []rune(a)
	bRune := []rune(b)
	aLen := len(aRune)
	bLen := len(bRune)

	if aLen == 0 {
		return 0
	}
	if bLen == 0 {
		return 0
	}

	aPtr := (*C.int32_t)(unsafe.Pointer(&aRune[0]))
	bPtr := (*C.int32_t)(unsafe.Pointer(&bRune[0]))

	cdist := C.lcs_len(aPtr, C.size_t(aLen), bPtr, C.size_t(bLen))
	return int(cdist)
}
