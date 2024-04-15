package helper

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

func PathExists(fs_path string) bool {
	_, err := os.Stat(fs_path)
	return !os.IsNotExist(err)
}

func ListDirFiles(dir string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// If it's a file, print its path
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func DirectoryIsReadableOrPanic(file_path string) {
	// Get file info
	fileInfo, err := os.Stat(file_path)
	if err != nil {
		panic(err)
	}
	// Check if file_path is directory
	if !fileInfo.IsDir() {
		panic("filepath is directory")
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

func DirectoryDeleteOrPanic(directory string) {
	err := os.RemoveAll(directory)
	if err == nil {
		msg := fmt.Sprintf("removed directory: %s", directory)
		slog.Debug(msg)
	} else {
		panic(err)
	}
}

func DirectoryTraverse(
	directory string,
	fn func(directory string, d fs.DirEntry) error,
	recurse bool,
) error {
	dirs, err := os.ReadDir(directory)
	if err != nil {
		// Cannot traverse directory at all
		return err
	}
	for _, fsPath := range dirs {
		// slog.Info(dir.Name())
		err := fn(directory, fsPath)
		if err != nil {
			return err
		}
		if fsPath.IsDir() {
			path_joined := filepath.Join(directory, fsPath.Name())
			if recurse {
				err := DirectoryTraverse(path_joined, fn, recurse)
				if err != nil {
					// Cannot traverse nested directory
					// slog.Error(err.Error())
					return err
				}
			}
		}
	}
	return nil
}

func DirectoryCopy(
	src_dir string,
	dst_dir string,
	recurse bool,
	overwrite bool,
	path_regex string,
	verbose bool,
) error {
	var regex_patt *regexp.Regexp
	if path_regex != "" {
		regex_patt = regexp.MustCompile(path_regex)
	}

	walk_func := func(fs_path string, d fs.DirEntry) error {
		if d.Type().IsRegular() {
			// Get current relative from src_dir
			relDir, err := filepath.Rel(src_dir, fs_path)
			if err != nil {
				return err
			}
			srcFile := filepath.Join(fs_path, d.Name())
			dstDir := filepath.Join(dst_dir, relDir)
			dstFile := filepath.Join(dstDir, d.Name())
			if regex_patt != nil && !regex_patt.MatchString(srcFile) {
				return nil
			}

			if err := os.MkdirAll(dstDir, 0700); err != nil {
				return err
			}
			slog.Debug("created", "path", dstDir)
			err = CopyFile(srcFile, dstFile, overwrite, verbose)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err := DirectoryTraverse(src_dir, walk_func, recurse)
	return err
}

func DirectoryCreateTemporaryOrPanic(base_name string) string {
	var err error
	var file_path string
	switch runtime.GOOS {
	case "linux":
		// Create temp directory in RAM
		// file_path, err = os.MkdirTemp("/dev/shm", base_name)
		file_path, err = os.MkdirTemp("/tmp/", base_name)
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
