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
