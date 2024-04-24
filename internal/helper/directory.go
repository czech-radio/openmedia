package helper

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

// PathExists report whether path exitst
func PathExists(fileSystemPath string) (bool, error) {
	_, err := os.Stat(fileSystemPath)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, err
}

// DirectoryExists report wheter path exists and is directory. Returns error
// when path exists and is not directory
func DirectoryExists(fileSystemPath string) (bool, error) {
	fileInfo, err := os.Stat(fileSystemPath)
	if err == nil {
		if fileInfo.IsDir() {
			return true, nil
		}
		// is something else
		return true, fmt.Errorf(
			"path is not directory: %s", fileSystemPath)
	}
	// path does not exists
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	// unknown filesystem error
	return false, err
}

// ListDirFiles list files inside directory
func ListDirFiles(dir string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// If it's a file
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// DirectoryIsReadableOrPanic report whether selected path is directory and is readable with current permission or panic.
func DirectoryIsReadableOrPanic(fileSystemPath string) {
	// Get file info
	fileInfo, err := os.Stat(fileSystemPath)
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
			panic(fmt.Sprintf(errmsg, fileSystemPath, fileInfo.Mode()))
		}
	case "windows":
		if fileInfo.Mode()&os.ModePerm == 0 {
			panic(fmt.Sprintf(errmsg, fileSystemPath, fileInfo.Mode()))
		}
	}
	// NOTE: Not accounting for ACL or xattrs
}

// DirectoryCreateInRam
func DirectoryCreateInRam(base_name string) string {
	filepath, err := os.MkdirTemp("/dev/shm", base_name)
	if err != nil {
		panic(err)
	}
	return filepath
}

// DirectoryDeleteOrPanic
func DirectoryDeleteOrPanic(directory string) {
	err := os.RemoveAll(directory)
	if err == nil {
		msg := fmt.Sprintf("removed directory: %s", directory)
		slog.Debug(msg)
	} else {
		panic(err)
	}
}

// DirectoryTraverse
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

// DirectoryCopy copies directory contents from source directory to destination directory
func DirectoryCopy(
	srcDir string,
	dstDir string,
	recurse bool,
	overwrite bool,
	pathRegex string,
	verbose bool,
) error {
	var regex_patt *regexp.Regexp
	if pathRegex != "" {
		regex_patt = regexp.MustCompile(pathRegex)
	}
	// TODO: add count of copied/overwritten files or directories
	// var dirCount, fileCount int
	walk_func := func(fs_path string, d fs.DirEntry) error {
		// Get current relative from src_dir
		relDir, err := filepath.Rel(srcDir, fs_path)
		if err != nil {
			return err
		}
		srcFile := filepath.Join(fs_path, d.Name())
		dstDir := filepath.Join(dstDir, relDir)
		if d.Type().IsRegular() {
			dstFile := filepath.Join(dstDir, d.Name())
			if regex_patt != nil && !regex_patt.MatchString(srcFile) {
				return nil
			}
			err := os.MkdirAll(dstDir, 0700)
			if err != nil {
				return err
			}
			if verbose {
				slog.Debug("created", "path", dstDir)
			}
			err = CopyFile(srcFile, dstFile, overwrite, verbose)
			if err != nil {
				return err
			}
		}
		if d.Type().IsDir() {
			// also copy empty directory
			dst := filepath.Join(dstDir, d.Name())
			if err := os.MkdirAll(dst, 0700); err != nil {
				return err
			}
		}
		return nil
	}
	err := DirectoryTraverse(srcDir, walk_func, recurse)
	return err
}

// DirectoryCreateTemporaryOrPanic create temporary directory. Resulting name of directory will be baseDirName+random_string
func DirectoryCreateTemporaryOrPanic(baseDirName string) string {
	var err error
	var file_path string
	switch runtime.GOOS {
	case "linux":
		// Create temp directory in RAM
		// file_path, err = os.MkdirTemp("/dev/shm", base_name)
		file_path, err = os.MkdirTemp("/tmp/", baseDirName)
	default:
		// Create temp directory in system default temp directory
		file_path, err = os.MkdirTemp("", baseDirName)
	}
	if err != nil {
		panic(err)
	}
	slog.Debug("Temp directory created: " + file_path)
	return file_path
}
