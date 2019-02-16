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