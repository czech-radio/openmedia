package internal

import (
	"fmt"
	"testing"
)

func Test_StructFieldsList(t *testing.T) {
	StructFieldsList(OM_FIELD{})
}

func Test_StructFieldsMap(t *testing.T) {
	res := StructFieldsMap(OM_FIELD{})
	fmt.Println(res)
}
