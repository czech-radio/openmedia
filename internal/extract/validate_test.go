package extract

import "testing"

func TestFormatFieldBeforeValidation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"exmp1", args{"Akk Kek"}, "akk kek"},
		{"exmp2", args{" Akk  Kek "}, "akk kek"},
		{"exmp3", args{"Akk Kek"}, "akk kek"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatFieldBeforeValidation(tt.args.input); got != tt.want {
				t.Errorf("FormatFieldBeforeValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}
