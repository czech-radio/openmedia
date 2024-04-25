package helper

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// List of errors
// /usr/lib/go/src/syscall/zerrors_linux_amd64.go:1490
// ControlFlowAction
type ControlFlowAction int

const (
	Continue ControlFlowAction = iota
	Skip
	Break
)

// ErrorCode codes for os.Exit
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
	ErrCodeParseField
	ErrCodeParseXML
	// ErrCodeInternal
	// ErrCodeSystem
	ErrCodeUnknown
)

type ErrorsCodeMap map[ErrorCode]string

// e.g. from package os: os.ErrPermission = errors.New("permission denied")
var Errors ErrorsCodeMap = ErrorsCodeMap{
	ErrCodeSuccess:    "",
	ErrCodePermission: "permission denied",   // os.ErrPermission
	ErrCodeExist:      "file already exists", // os.ErrExist
	ErrCodeNotExist:   "file does not exist", // os.ErrNotExist
	ErrCodeInvalid:    "invalid file",        // os.ErrInvalid
	ErrCodeClosed:     "file already closed", // os.ErrClosed
	ErrCodeUnknown:    "uknown error",
	ErrCodeParseXML:   "cannot parse xml",
	ErrCodeParseField: "cannot parse field",
	// ErrCodeDataFormat: "file has incompatible format",
}

// ErrorWrap
func ErrorWrap(fieldName, fieldValue string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %s, %w", fieldName, fieldValue, err)
	}
	return nil
}

// CodeMsg
func (ecm ErrorsCodeMap) CodeMsg(code ErrorCode) string {
	return ecm[code]
}

// ErrorBaseMessage
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

// ErrorCode
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

// ExitWithCode
func (ecm ErrorsCodeMap) ExitWithCode(err error) {
	code := ecm.ErrorCode(err)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(int(code))
	}
}

type ErrorsAgregate struct {
	Errors   []error
	Messages []string
}

// MessageAdd
func (ea *ErrorsAgregate) MessageAdd(msg string) {
	if msg != "" {
		ea.Messages = append(ea.Messages, msg)
	}
}

// MessagesJoin
func (ea *ErrorsAgregate) MessagesJoin() string {
	return strings.Join(ea.Messages, ", ")
}

// ErrorAdd
func (ea *ErrorsAgregate) ErrorAdd(err error) {
	if err != nil {
		ea.Errors = append(ea.Errors, err)
	}
}

// GetError
func (ea *ErrorsAgregate) GetError() {
}

// GetMessage
func (ea *ErrorsAgregate) GetMessage() {
}

// ErrorAppend
func ErrorAppend(errs []error, err error) []error {
	var resErrs []error
	for _, e := range errs {
		if e != nil {
			resErrs = append(resErrs, e)
		}
	}
	return resErrs
}
