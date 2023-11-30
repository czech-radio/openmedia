package internal

import (
	"errors"
	"log/slog"
	"os"
)

type ErrorCode int8

// Error codes
const (
	ErrCodeSuccess        ErrorCode = 0
	ErrCodeInvalidCommand ErrorCode = 1
	ErrCodePermission     ErrorCode = 2 + iota
	ErrCodeExist
	ErrCodeNotExist
	ErrCodeClosed
	ErrCodeDataFormat
	// ErrCodeInternal
	// ErrCodeSystem
	ErrCodeUnknown
)

// ErrorCodeMap map ErrorCode and bas error message
var ErrorCodeMap map[ErrorCode]string = map[ErrorCode]string{
	ErrCodeSuccess:    "",
	ErrCodePermission: "permission denied",
	ErrCodeExist:      "file already exists",
	ErrCodeNotExist:   "file does not exist",
	ErrCodeClosed:     "file already closed",
	// ErrCodeDataFormat: "file has incompatible f",
	// ErrCodeUnknown:  errors.New("uknown error"),
}

// var (
// os
// ErrInvalid    = errors.New("invalid argument")
// ErrPermission = errors.New("permission denied")
// ErrExist      = errors.New("file already exists")
// ErrNotExist   = errors.New("file does not exist")
// ErrClosed     = errors.New("file already closed")
// )

func ErrorGetBaseMessage(err error) string {
	var baseErr error = err
	var unwrapErr error
	if err == nil {
		return ""
	}

	// Unwrap error as much as possible
	for {
		unwrapErr = errors.Unwrap(baseErr)
		if unwrapErr == nil {
			break
		} else {
			baseErr = unwrapErr
		}
	}
	return baseErr.Error()
}

func ErrorGetCode(err error) ErrorCode {
	var resultCode ErrorCode
	var resultCodeFound bool
	baseMsg := ErrorGetBaseMessage(err)
	for errCode, errMsg := range ErrorCodeMap {
		if baseMsg == errMsg {
			resultCode = errCode
			resultCodeFound = true
		}
	}
	if !resultCodeFound {
		resultCode = ErrCodeUnknown
	}

	return resultCode
}

func ErrorExitWithCode(err error) {
	code := ErrorGetCode(err)
	if err != nil {
		slog.Error(err.Error())
	}
	os.Exit(int(code))
}

// Alt method using errors.Is:
// for errCode, errBase := range ErrorCodeMap {
// if errors.Is(err, errBase) {
// resultCode = errCode
// errCodeFound = true
// }
// }
// if !errCodeFound {
// resultCode = ErrCodeUnknown
// }
