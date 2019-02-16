# xerrors sample

[xerrors - GoDoc](https://godoc.org/golang.org/x/xerrors)

## Wrap

```go
package main

import (
    "fmt"
    "golang.org/x/xerrors"
)

func main() {
    baseErr := xerrors.New("base error")
    wrapErr1 := xerrors.Errorf("wrap error1: %v",baseErr)
    wrapErr2 := xerrors.Errorf("wrap error2: %v",wrapErr1)
    fmt.Printf("%+v\n", wrapErr2)
}
```

### Wrap Result

```text
wrap error2:
    main.main
        /Users/foo/Desktop/xerrors/main.go:11
  - wrap error1:
    main.main
        /Users/foo/Desktop/xerrors/main.go:10
  - base error:
    main.main
        /Users/foo/Desktop/xerrors/main.go:9
```
