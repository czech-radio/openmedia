package helper

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	enc_unicode "golang.org/x/text/encoding/unicode"
)

var ErrFilePathExists = errors.New("file path exists")

type FileEncodingNumber int

const (
	UNKNOWN FileEncodingNumber = iota
	UTF8
	UTF16le
	UTF16be
)

func HandleFileEncoding(
	enc FileEncodingNumber, ioReaderCloser io.ReadCloser) ([]byte, error) {
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

func CopyFile(
	src_file_path, dst_file_path string,
	overwrite bool,
) error {
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
	if !overwrite && PathExists(dst_file_path) {
		// return fmt.Errorf("destion path exists: %s", dst_file_path)
		return fmt.Errorf(
			"err: %w, filepath: %s", ErrFilePathExists, dst_file_path,
		)
	}
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

// FileExists: returns true if file exists, false when the filePath doesnot exists, error when it is directory
func FileExists(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if info.IsDir() {
		return false, fmt.Errorf("specified path is a directory, not a zip file")
	}
	return true, nil
}

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
