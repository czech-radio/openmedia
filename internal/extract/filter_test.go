package extract

import (
	"testing"
)

func TestGetFilterByFilterFileName(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name    string
		args    args
		want    FilterFileCode
		wantErr bool
	}{
		{"eurovolby", args{
			"filtr_eurovolby_v1.xlsx"}, FilterFileEuroElection, false},
		{"opozice", args{
			"filtr_opozice_2024-04-01_2024-05-31_v1.xlsx"}, FilterFileOposition, false},
		{"unknown", args{
			"filtr_wrong_2024-04-01_2024-05-31_v1.xlsx"}, FilterFileCode(""), true},
		{"empty", args{
			"filtr_opozic_2024-04-01_2024-05-31_v1.xlsx"}, FilterFileCode(""), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FilterFileCodeMap.GetFilterFileCode(tt.args.fname)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilterByFilterFileName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFilterByFilterFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchStringElements(t *testing.T) {
	type args struct {
		str1 []string
		str2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"one", args{
			[]string{"kek", "jek", "tek"},
			[]string{"kek", "tek"}}, true,
		},
		{"two", args{
			[]string{"kek", "jek", "tek"},
			[]string{"kek", "sek"}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchStringElements(tt.args.str1, tt.args.str2, 2); got != tt.want {
				t.Errorf("MatchStringElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
