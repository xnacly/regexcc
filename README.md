# regexcc

> Compile regular expressions to optimised C code 

## Usage

Executing 

```shell
regexcc '^([0-2][0-9]|3[01])-(0[1-9]|1[0-2])-[0-9]{4}$' -name date
```

Results in `date.h` and `date.c` being generated:

```c
#ifndef DATE_REGEX_H
#define DATE_REGEX_H

#include <stddef.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

// Returns true if the input matches '^([0-2][0-9]|3[01])-(0[1-9]|1[0-2])-[0-9]{4}$'
bool date_match(const char *input, size_t length);

#ifdef __cplusplus
}
#endif

#endif
```

```c
#include "date.h"
#include <stdbool.h>
#include <stddef.h>
#include <ctype.h>

// Compiled from regex: ^([0-2][0-9]|3[01])-(0[1-9]|1[0-2])-[0-9]{4}$
bool date_match(const char *s, size_t n) {
    if (n != 10) {
        return false 
    };

    // ([0-2][0-9]|3[01])
    if (s[0] < '0' || s[0] > '3') return false;
    if (s[0] == '3' && (s[1] < '0' || s[1] > '1')) return false;
    if (s[0] < '3' && !isdigit((unsigned char)s[1])) return false;

    // -
    if (s[2] != '-') {
        return false
    };

    // (0[1-9]|1[0-2])
    if (s[3] == '0') {
        if (s[4] < '1' || s[4] > '9') return false;
    } else if (s[3] == '1') {
        if (s[4] < '0' || s[4] > '2') return false;
    } else {
        return false;
    }

    // -
    if (s[5] != '-') return false;

    // [0-9]{4}
    for (int i = 6; i < 10; i++) {
        if (s[i] < '0' || s[i] > '9') {
            return false;
        }
    }

    return true;
}
```

### In C

These can be included as any other header source pair:

```c
#include <stdio.h>
#include <string.h>
#include "date.h"

int main(void) {
    const char *samples[] = {
        "05-11-2025",
        "31-02-2024",
        "5-11-2025",
        "2025-11-05"
    };

    for (size_t i = 0; i < sizeof(samples)/sizeof(samples[0]); i++) {
        const char *s = samples[i];
        bool ok = date_match(s, strlen(s));
        printf("%s => %s\n", s, ok ? "matches" : "doesnt match");
    }
}
```

Compiling this with `date.c`

```shell
gcc -O3 main.c date.c -o test_date
./test_date
```

Results in:

```text
05-11-2025 => matches
31-02-2024 => matches
5-11-2025 => doesnt match
2025-11-05 => doesnt match
```

### In Go

```go
// main.go
package main

/*
#cgo CFLAGS: -O3
#include "date.h"
*/
import "C"
import (
    "fmt"
    "unsafe"
)

// DateMatch wraps date_match from date.h
func DateMatch(s string) bool {
    p := unsafe.Pointer(&[]byte(s)[0])
    return bool(C.date_match((*C.char)(p), C.size_t(len(s))))
}

func main() {
    tests := []string{"05-11-2025", "31-02-2024", "5-11-2025", "2025-11-05"}
    for _, t := range tests {
        fmt.Printf("%s => %v\n", t, DateMatch(t))
    }
}
```
