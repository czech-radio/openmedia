package helper

import (
	"testing"
)

func TestFileExists(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, "helper")

	tp := testerConfig.TempSourcePathGeter("helper")
	// Prepare test pairs
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"file_exits", args{tp("some_file.txt")}, true, false},
		{"file_not_exits", args{tp("nonexisten_file.txt")}, false, false},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileExists(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"FileExists() error = %v, wantErr %v, args %v",
					err, tt.wantErr, tt.args)
				return
			}
			if got != tt.want {
				t.Errorf(
					"FileExists() = %v, want %v, args %v",
					got, tt.want, tt.args)
			}
		})
	}
}
