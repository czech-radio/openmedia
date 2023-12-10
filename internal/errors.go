package internal

import (
	"errors"
	"log/slog"
	"os"
)

// Error codes for os.Exit
type ErrorCode int8

const (
	ErrCodeSuccess        ErrorCode = 0
	ErrCodeInvalidCommand ErrorCode = 1
	ErrCodePermission     ErrorCode = 2 + iota
	ErrCodeExist
	ErrCodeNotExist
	ErrCodeClosed
	ErrCodeDataFormat
	ErrCodeInvalid
	ErrCodeParsingField
	ErrCodeParsingXML
	// ErrCodeInternal
	// ErrCodeSystem
	ErrCodeUnknown
)

type ErrorsCodeMap map[ErrorCode]string

var Errors ErrorsCodeMap = ErrorsCodeMap{
	ErrCodeSuccess: "",
	// os.ErrPermission = errors.New("permission denied")
	ErrCodePermission:   "permission denied",   // os.ErrPermission
	ErrCodeExist:        "file already exists", // os.ErrExist
	ErrCodeNotExist:     "file does not exist", // os.ErrNotExist
	ErrCodeInvalid:      "invalid file",        // os.ErrInvalid
	ErrCodeClosed:       "file already closed", // os.ErrClosed
	ErrCodeUnknown:      "uknown error",
	ErrCodeParsingXML:   "cannot parse xml",
	ErrCodeParsingField: "cannot parse field",
	// ErrCodeDataFormat: "file has incompatible format",
}

func (ecm ErrorsCodeMap) CodeMsg(code ErrorCode) string {
	return ecm[code]
}

func (ecm ErrorsCodeMap) ErrorBaseMessage(err error) string {
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

func (ecm ErrorsCodeMap) ErrorCode(err error) ErrorCode {
	var resultCode ErrorCode
	var resultCodeFound bool
	baseMsg := ecm.ErrorBaseMessage(err)
	for errCode, errMsg := range Errors {
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

func (ecm ErrorsCodeMap) ErrorExitWithCode(err error) {
	code := ecm.ErrorCode(err)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(int(code))
	}
}

func ErrorAppend(errs []error, err error) []error {
	var resErrs []error
	for _, e := range errs {
		if e != nil {
			resErrs = append(resErrs, e)
		}
	}
	return resErrs
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
