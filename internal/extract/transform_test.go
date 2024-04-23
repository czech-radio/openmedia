package extract

import "testing"

func TestTransformTema(t *testing.T) {
	type args struct {
		tema string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"normal_complex",
			args{"08-EU; 01-Politika, státní správa a samospráva; 07-Kultura a historie"}, "08;01;07;", false},
		{"normal_simple",
			args{"08-EU"}, "08;", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformTema(tt.args.tema)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformTema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TransformTema() = %v, want %v", got, tt.want)
			}
		})
	}
}
