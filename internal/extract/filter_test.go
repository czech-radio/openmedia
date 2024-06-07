package extract

import "testing"

// func TestUniqSliceInt(t *testing.T) {
// 	type args struct {
// 		A []int
// 		B []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []int
// 	}{
// 		{"one", args{[]int{1, 3, 5, 10}, []int{2, 3, 10, 11}}, []int{3, 10}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := UniqSliceInt(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UniqSliceInt() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

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
			got, err := GetFilterFileCode(tt.args.fname)
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
