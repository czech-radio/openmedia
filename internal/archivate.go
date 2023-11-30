package internal

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type ArchiveResult struct {
	FilesProcessed int
	FilesSuccess   int
	Errors         []string
}

func ValidateFileName(src_path string) (bool, error) {
	file_extension := filepath.Ext(src_path)
	if file_extension != ".xml" {
		return false,
			fmt.Errorf("file does not have xml extension: %s", src_path)
	}
	if !strings.Contains(src_path, "RD") {
		return false,
			fmt.Errorf("filename does not contaion 'RD' string: %s", src_path)
	}
	_, err := os.Stat(src_path)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *ArchiveResult) AddError(err error) {
	b.Errors = append(b.Errors, err.Error())
}

func ZipArchive(sourceDir, zipFile string) (error, *ArchiveResult) {
	// Create or truncate the zip file
	archive, err := os.Create(zipFile)
	if err != nil {
		return err, nil
	}
	defer archive.Close()

	// Create a zip writer
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	// Walk through the source directory and add files to the zip archive
	var result ArchiveResult = ArchiveResult{}
	walk_func := func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}

		result.FilesProcessed++
		// Validate file that is valid xml rundown
		ok, err := ValidateFileName(filePath)
		if !ok {
			result.AddError(err)
			slog.Error(err.Error())
			return nil
		}

		slog.Debug("compressing", "file", filePath)
		err = ZipArchiveAdd(zipWriter, sourceDir, filePath)
		if err != nil {
			slog.Error("compressing", "file", filePath)
			result.AddError(err)
			return nil
		}
		result.FilesSuccess++
		return nil
	}
	filepath.Walk(sourceDir, walk_func)
	errorsCount := len(result.Errors)
	if errorsCount > 0 {
		return fmt.Errorf("there were %d errors, check results", errorsCount), &result
	}
	return nil, &result
}

func ZipArchiveAdd(zipWriter *zip.Writer, sourceDir, filePath string) error {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		// result.AddError(err)
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Create a new zip file header
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Method = zip.Deflate
	// Set the name of the file within the zip archive
	header.Name, err = filepath.Rel(sourceDir, filePath)
	if err != nil {
		return err
	}

	// Create a new entry in the zip archive
	entry, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// Copy the file content to the zip entry
	_, err = io.Copy(entry, file)
	if err != nil {
		return err
	}
	return nil

}

func TarGzArchive(sourceDir, tarGzFile string) (error, *ArchiveResult) {
	// Create or truncate output file
	archive, err := os.Create(tarGzFile)
	if err != nil {
		return err, nil
	}
	defer archive.Close()
	gzipWriter := gzip.NewWriter(archive)
	defer gzipWriter.Close()
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	var result ArchiveResult = ArchiveResult{}
	walk_func := func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}
		result.FilesProcessed++
		// Validate file that is valid xml rundown
		ok, err := ValidateFileName(filePath)
		if !ok {
			slog.Error("validate", "file", filePath)
			result.AddError(err)
			return nil
		}
		slog.Debug("compressing", "file", filePath)
		err = TarGzArchiveAdd(tarWriter, sourceDir, filePath)
		if err != nil {
			slog.Error("compressing", "file", filePath)
			result.AddError(err)
			return nil
		}
		result.FilesSuccess++
		return nil
	}
	filepath.Walk(sourceDir, walk_func)
	errorsCount := len(result.Errors)
	if errorsCount > 0 {
		return fmt.Errorf("there were %d errors, check results", errorsCount), &result
	}

	return nil, &result
}

func TarGzArchiveAdd(tw *tar.Writer, source_dir, source_filename string) error {
	file, err := os.Open(source_filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Get relative path in archive
	relPath, err := filepath.Rel(source_dir, source_filename)
	if err != nil {
		return err
	}

	// Create a tar Header from the FileInfo data
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	header.Name = relPath

	// Write file header to the tar archive
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	// Copy file content to tar archive
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}
