package helper

import (
	"fmt"
	"os"
	"testing"
)

func TestErrorGetBaseMessage(t *testing.T) {
	wrapErr := fmt.Errorf("%w: folder kek", os.ErrPermission)
	wrapErr2 := fmt.Errorf("%w: time", wrapErr)
	wrapErr3 := fmt.Errorf("%w: to go home", wrapErr2)
	msg := Errors.ErrorBaseMessage(wrapErr3)
	fmt.Println(msg)
}

func TestErrorGetCode(t *testing.T) {
	wrapErr := fmt.Errorf("%w: folder kek", os.ErrPermission)
	wrapErr2 := fmt.Errorf("%w: time", wrapErr)
	wrapErr3 := fmt.Errorf("%w: to go home", wrapErr2)
	code := Errors.ErrorCode(wrapErr3)
	fmt.Println(code)
}
