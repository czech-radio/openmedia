package extract

import (
	"reflect"
	"testing"
)

func TestUniqSliceInt(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"one", args{[]int{1, 3, 5, 10}, []int{2, 3, 10, 11}}, []int{3, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqSliceInt(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqSliceInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
