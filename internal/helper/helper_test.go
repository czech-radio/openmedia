package helper

import (
	"fmt"
	"testing"
)

func Test_LogTraceFunction(t *testing.T) {
	fmt.Println(TraceFunction(0))
	fmt.Println(TraceFunction(1))
}

func TestIter(t *testing.T) {
	abc := []string{"a", "b", "c", "d"}
	abc0 := []string{"a"}
	for i, str := range abc0 {
		fmt.Println(i, str, len(abc))
		if i == 3 {
			continue
		}
	}
}
