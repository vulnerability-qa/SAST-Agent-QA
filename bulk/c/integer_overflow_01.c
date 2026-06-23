/* CWE-190: Integer Overflow in malloc size */
#include <stdlib.h>
void *alloc_items(int count, int item_size) {
    return malloc(count * item_size); /* overflows if large values */
}
