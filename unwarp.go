package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	baseErr := xerrors.New("base error")
	wrapErr := xerrors.Errorf("error in main method: %w", baseErr)
	fmt.Printf("%+v\n", xerrors.Unwrap(wrapErr))
}
