# Error handling in Go


```go
package main

import (
    "errors"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        // TAKE: you do not need to declare a var for error BC it is already defined in the return signature above
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```