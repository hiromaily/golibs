#include "levenshtein.h"
#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>

unsigned int
levenshtein(int32_t *a, size_t a_size, int32_t *b, size_t b_size)
{
    unsigned int *m = calloc(a_size, sizeof(unsigned int));
    if (m == NULL) {
        return -1;
    }

    unsigned int result, i, j, d, t;

    for (i = 0; i < a_size; i++) {
        m[i] = i + 1;
    }

    for (i = 0; i < b_size; i++) {
        result = d = i;

        for (j = 0; j < a_size; j++) {
            t = (b[i] == a[j]) ? d : d + 1;
            d = m[j];

            m[j] = result = (d > result)
                ? (t > result ? result + 1 : t)
                : (t > d ? d + 1 : t);
        }
    }

    free(m);

    return result;
}

unsigned int
lcs_len(int32_t *a, size_t a_size, int32_t *b, size_t b_size)
{
    unsigned int *m = calloc(a_size, sizeof(unsigned int));
    if (m == NULL) {
        return -1;
    }

    unsigned int result, i, j, d, t;

    for (i = 0; i < b_size; i++) {
        result = d = 0;

        for (j = 0; j < a_size; j++) {
            t = (b[i] == a[j]) ? d + 1 : d;
            d = m[j];

            m[j] = result = (d > result)
                ? (t > d ? t : d)
                : (t > result ? t : result);
        }
    }

    free(m);

    return result;
}
