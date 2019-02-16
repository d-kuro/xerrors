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
	wrapErr1 := xerrors.Errorf("wrap error1: %w", baseErr)
	wrapErr2 := xerrors.Errorf("wrap error2: %w", wrapErr1)
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

## Unwrap

```go
package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	baseErr := xerrors.New("base error")
	wrapErr := xerrors.Errorf("wrap error: %w", baseErr)
	fmt.Printf("%+v\n", xerrors.Unwrap(wrapErr))
}
```

### Unwrap Result

```text
base error:
    main.main
        /Users/foo/Desktop/xerrors/unwarp.go:9
```

## Is

```go
package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	baseErr := xerrors.New("base error")
	fmt.Println("xerrors.Is(baseErr, baseErr): ", xerrors.Is(baseErr, baseErr))

	wrapErr := xerrors.Errorf("wrap error: %w", baseErr)
	fmt.Println("xerrors.Is(wrapErr, baseErr): ", xerrors.Is(wrapErr, baseErr))
}
```

### Is Result

```text
xerrors.Is(baseErr, baseErr):  true
xerrors.Is(wrapErr, baseErr):  true
```

## As

```go
package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

type err struct {
	message string
}

func (e *err) Error() string {
	return e.message
}

func main() {
	baseErr := &err{"base error"}
	wrapErr1 := xerrors.Errorf("wrap error1: %w", baseErr)
	wrapErr2 := xerrors.Errorf("wrap error2: %w", wrapErr1)

	var baseErr2 *err
	if xerrors.As(wrapErr2, &baseErr2) {
		fmt.Println("xerrors.As(wrapErr2, baseErr2): ", xerrors.As(wrapErr2, &baseErr2))
		fmt.Printf("baseErr2: %+v", baseErr2)
	}
}
```

### As Result

```text
xerrors.As(wrapErr2, baseErr2):  true
baseErr2: base error
```
