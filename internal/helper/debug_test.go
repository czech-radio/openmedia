package helper

import (
	"fmt"
	"testing"
)

func Test_LogTraceFunction(t *testing.T) {
	fmt.Println(TraceFunction(0))
	fmt.Println(TraceFunction(1))
}
