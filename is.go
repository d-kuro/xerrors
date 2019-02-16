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
