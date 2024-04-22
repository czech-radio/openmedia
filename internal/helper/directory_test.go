package helper

import (
	"errors"
	"os"
	"testing"
)

func TestDirectoryCreateTemporary(t *testing.T) {
	directory := DirectoryCreateTemporaryOrPanic("golang_test")
	defer os.RemoveAll(directory)
}

func Test_CurrentDir(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Log(err)
	}
	if err == nil {
		t.Log(dir)
	}
}

func TestCreateDirectory(t *testing.T) {
	testSubdir := "helper"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	tpDst := testerConfig.TempDestinationPathGeter(testSubdir)
	dstDir := tpDst("hello/jello")
	err := os.MkdirAll(dstDir, 0700)
	if err != nil {
		t.Error(err)
	}
	dstDir = tpDst("")
	err = os.MkdirAll(dstDir, 0700)
	if err != nil {
		t.Error(err)
	}
}

func Test_DirectoryCreateInRam(t *testing.T) {
	directory := DirectoryCreateInRam("golang_test")
	defer os.RemoveAll(directory)
}

func Test_DirectoryCopy(t *testing.T) {
	testSubdir := "helper"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	tpSrc := testerConfig.TempSourcePathGeter(testSubdir)
	srcDir := tpSrc("")
	tpDst := testerConfig.TempDestinationPathGeter(testSubdir)
	dstDir := tpDst("")

	// Test copy matching files
	err := DirectoryCopy(
		srcDir, dstDir,
		true, false, "hello", false)
	if err != nil {
		t.Error(err)
	}

	// Test copy recursive and overwrite destination files
	err = DirectoryCopy(
		srcDir, dstDir, true, true, "", false)
	if err != nil && errors.Unwrap(err) != ErrFilePathExists {
		t.Error(err)
	}
}

func TestPathExists(t *testing.T) {
	testSubdir := "helper"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	tpSrc := testerConfig.TempSourcePathGeter(testSubdir)

	type args struct {
		fs_path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"path_exists_file", args{tpSrc("some_file.txt")},
			true, false},
		{"path_exists_dir", args{tpSrc("")},
			true, false},
		{"path_not_exists_dir", args{tpSrc("kek/")},
			false, false},
		{"path_not_exists_file", args{tpSrc("kek")},
			false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PathExists(tt.args.fs_path)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"PathExists() error = %v, wantErr %v, args %v",
					err, tt.wantErr, tt.args)
				return
			}
			if got != tt.want {
				t.Errorf(
					"PathExists() = %v, want %v, args %v",
					got, tt.want, tt.args)
			}
		})
	}
}

func TestDirectoryExists(t *testing.T) {
	testSubdir := "helper"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	tpSrc := testerConfig.TempSourcePathGeter(testSubdir)
	type args struct {
		fs_path string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{

		{"path_exists_file", args{tpSrc("some_file.txt")},
			true, true},
		{"path_exists_dir", args{tpSrc("")},
			true, false},
		{"path_not_exists_dir", args{tpSrc("kek/")},
			false, false},
		{"path_not_exists_file", args{tpSrc("kek")},
			false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DirectoryExists(tt.args.fs_path)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"PathExists() error = %v, wantErr %v, args %v",
					err, tt.wantErr, tt.args)
				return
			}
			if got != tt.want {
				t.Errorf(
					"PathExists() = %v, want %v, args %v",
					got, tt.want, tt.args)
			}
		})
	}
}
