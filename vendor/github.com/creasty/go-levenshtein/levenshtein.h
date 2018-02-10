#ifndef __LEVENSHTEIN_H_LOADED__
#define __LEVENSHTEIN_H_LOADED__

#include <stddef.h>
#include <stdint.h>

unsigned int
levenshtein(int32_t *a, size_t a_size, int32_t *b, size_t b_size);

unsigned int
lcs_len(int32_t *a, size_t a_size, int32_t *b, size_t b_size);

#endif // __LEVENSHTEIN_H_LOADED__
