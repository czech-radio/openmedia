// Package
package internal

import (
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

func DetectLinuxSytemOrPanic() {
	if runtime.GOOS != "linux" {
		msg := fmt.Sprintf("unsuported OS: %s, %s", runtime.GOOS, runtime.GOARCH)
		panic(msg)
	}
}

func DirectoryIsReadableOrPanic(file_path string) {
	// Get file info
	fileInfo, err := os.Stat(file_path)
	if err != nil {
		panic(err)
	}
	// Check if file_path is directory
	if !fileInfo.IsDir() {
		panic(err)
	}

	// Check file_path file mode or file permission
	errmsg := "directory not readable: %s, filemode: %s"
	switch runtime.GOOS {
	case "linux":
		// Check linux permission. Readable for current user has value > 0400
		if fileInfo.Mode().Perm()&0400 == 0 {
			// bitwise &:
			// 0700 & 0400 -> 100000000 -> 1
			// 0600 & 0400 -> 100000000
			// 0500 & 0400 -> 100000000
			// 0100 & 0400 -> 000000000
			// 0000 & 0400 -> 000000000 -> 0
			panic(fmt.Sprintf(errmsg, file_path, fileInfo.Mode()))
		}
	case "windows":
		if fileInfo.Mode()&os.ModePerm == 0 {
			panic(fmt.Sprintf(errmsg, file_path, fileInfo.Mode()))
		}
	}
	// NOTE: Not accounting for ACL or xattrs
}

func DirectoryCreateInRam(base_name string) string {
	filepath, err := os.MkdirTemp("/dev/shm", base_name)
	if err != nil {
		panic(err)
	}
	return filepath
}

func DirectoryCreateTemporaryOrPanic(base_name string) string {
	var err error
	var file_path string
	switch runtime.GOOS {
	case "linux":
		// Create temp directory in RAM
		file_path, err = os.MkdirTemp("/dev/shm", base_name)
	default:
		// Create temp directory in system default temp directory
		file_path, err = os.MkdirTemp("", base_name)
	}
	if err != nil {
		panic(err)
	}
	slog.Debug("Temp directory created: " + file_path)
	return file_path
}

func DirectoryDeleteOrPanic(directory string) {
	err := os.RemoveAll(directory)
	if err == nil {
		msg := fmt.Sprintf("removed directory: %s", directory)
		slog.Debug(msg)
	} else {
		panic(err)
	}
}

func DirectoryWalk(directory string) {
	walk_func := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		// fmt.Printf("dir: %t: name: %s\n", d.IsDir(), path)
		return nil
	}
	err := filepath.WalkDir(directory, walk_func)
	if err != nil {
		fmt.Println(err)
	}
}

func ListFsPath(fs_path string, d fs.DirEntry, err error) error {
	if err != nil {
		slog.Error(err.Error())
		return nil
	}
	fmt.Printf("list_func: %s, %s, %t\n", fs_path, d.Name(), d.IsDir())
	return nil
}

func DirectoryTraverse(
	fs_path string,
	fn func(path string, d fs.DirEntry, err error) error,
	recurse bool,
) error {
	dirs, err := os.ReadDir(fs_path)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}
	for _, dir := range dirs {
		err := fn(fs_path, dir, nil)
		if err != nil {
			return err
		}
		if dir.IsDir() {
			path_joined := filepath.Join(fs_path, dir.Name())
			if recurse {
				err := DirectoryTraverse(path_joined, fn, recurse)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func DirectoryFileList(file_path string) error {
	dirs, err := os.ReadDir(file_path)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		fmt.Println(dir.Name(), dir.Type(), dir.Type().IsRegular())
	}
	return nil
}

func CopyFile(src_file_path, dst_file_path string) error {
	slog.Debug(
		"copying file",
		"source_file", src_file_path,
		"dst_file", dst_file_path,
	)
	srcFile, err := os.Open(src_file_path)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file in the destination directory for writing
	dstFile, err := os.Create(dst_file_path)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func DirectoryCopyNoRecurse(src_dir, dst_dir string) (int, error) {
	files_count := 0 // number of files copied
	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(dst_dir, 0700); err != nil {
		return files_count, err
	}
	entries, err := os.ReadDir(src_dir)
	if err != nil {
		return files_count, err
	}

	for _, entry := range entries {
		if entry.Type().IsRegular() {
			srcFilePath := filepath.Join(src_dir, entry.Name())
			dstFilePath := filepath.Join(dst_dir, entry.Name())
			err = CopyFile(srcFilePath, dstFilePath)
			if err != nil {
				return files_count, err
			}
			files_count++
		}
	}
	return files_count, nil
}
