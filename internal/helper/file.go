package helper

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	enc_unicode "golang.org/x/text/encoding/unicode"
)

var ErrFilePathExists = errors.New("file path exists")

// FileEncodingCode is code, which specifies the file encoding.
type FileEncodingCode int

const (
	UNKNOWN FileEncodingCode = iota
	UTF8
	UTF16le
	UTF16be
)

// HandleFileEncoding read file according to specified file encoding code.
func HandleFileEncoding(
	enc FileEncodingCode, ioReaderCloser io.ReadCloser) ([]byte, error) {
	var data []byte
	var err error
	switch enc {
	case UTF8:
		data, err = io.ReadAll(ioReaderCloser)
	case UTF16le:
		utf8reader := enc_unicode.UTF16(enc_unicode.LittleEndian, enc_unicode.IgnoreBOM).NewDecoder().Reader(ioReaderCloser)
		data, err = io.ReadAll(utf8reader)
	default:
		err = fmt.Errorf("unknown encoding")
	}
	return data, err
}

// CopyFile copies file from source path to destination path.
func CopyFile(
	srcFilePath, dstFilePath string,
	overwrite bool, verbose bool,
) error {
	if verbose {
		slog.Debug(
			"copying file",
			"source_file", srcFilePath,
			"dst_file", dstFilePath,
		)
	}
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file in the destination directory for writing
	pathExists, err := PathExists(dstFilePath)
	if err != nil {
		return err
	}
	if !overwrite && pathExists {
		return fmt.Errorf(
			"err: %w, filepath: %s",
			ErrFilePathExists, dstFilePath,
		)
	}
	dstFile, err := os.Create(dstFilePath)
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

// FileExists: returns true if file exists, false when the filePath doesnot exists, error when it is directory
func FileExists(filePath string) (bool, error) {
	fileInfo, err := os.Stat(filePath)
	if err == nil {
		if fileInfo.IsDir() {
			return false, fmt.Errorf("specified path is a directory, not a file: %s", filePath)
		}
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

// ProcessedFileRename adds prfix to original filepath.
func ProcessedFileRename(originalPath string) error {
	fileName := filepath.Base(originalPath)
	directory := filepath.Dir(originalPath)
	newPath := filepath.Join(directory, "processed_"+fileName)
	err := os.Rename(originalPath, newPath)
	if err != nil {
		return fmt.Errorf("Error renaming file: %s", err)
	}
	return nil
}

// ReadCSVfile
func ReadCSVfile(filePath string) ([][]string, error) {
	file, err := os.Open("Students.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	// reads all the records from the CSV file and return [][]string
	return reader.ReadAll()
}
