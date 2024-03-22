package internal

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ncruces/go-strftime"

	enc_unicode "golang.org/x/text/encoding/unicode"
)

type VersionInfo struct {
	Version   string
	GitTag    string
	GitCommit string
	BuildTime string
}

// SetLogLevel: sets log level, default=0
func SetLogLevel(level string, logType ...string) {
	var logger *slog.Logger
	var loggerType string
	intlevel, err := strconv.Atoi(level)
	if err != nil {
		intlevel = 0
	}
	hopts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(intlevel),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				// Shorten the the filepath in log
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					source.File = filepath.Base(source.File)
				}
			}
			return a
		},
	}
	if logType != nil && logType[0] != "" {
		loggerType = logType[0]
	}
	switch loggerType {
	case "json":
		jhandle := slog.NewJSONHandler(os.Stderr, &hopts)
		logger = slog.New(jhandle)
	case "plain":
		thandle := slog.NewTextHandler(os.Stderr, &hopts)
		logger = slog.New(thandle)
	default:
		thandle := slog.NewTextHandler(os.Stderr, &hopts)
		logger = slog.New(thandle)
	}
	slog.SetDefault(logger)
}

func PrintObjJson(mark string, input any) {
	res, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		slog.Error("cannot marshal structure", "mark", mark, "input", input, "err", err.Error())
		return
	}
	fmt.Println(mark, string(res))
}

// Sleeper sleeps for specified durration
func Sleeper(duration int, time_unit string) {
	switch time_unit {
	case "ms":
		time.Sleep(time.Duration(duration) * time.Millisecond)
	case "s":
		time.Sleep(time.Duration(duration) * time.Second)
	case "m":
		time.Sleep(time.Duration(duration) * time.Minute)
	default:
		panic("Wrong time time_unit")
	}
}

func IsOlderThanOneISOweek(dateToCheck, dateNow time.Time) bool {
	year_check, week_check := dateToCheck.ISOWeek()
	year_now, week_now := dateNow.ISOWeek()
	if year_check < year_now {
		return true
	}
	return week_check < week_now
}

func TimeCurrent() string {
	tm := time.Now()
	return strftime.Format("%FT%T", tm)
}

var ErrFilePathExists = errors.New("file path exists")

func PathExists(fs_path string) bool {
	_, err := os.Stat(fs_path)
	return !os.IsNotExist(err)
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
			err = CopyFile(srcFile, dstFile, overwrite)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err := DirectoryTraverse(src_dir, walk_func, recurse)
	return err
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

func XOR(a, b bool) bool {
	return (a || b) && !(a && b)
}

func TraceFunction(depth int) (string, string, int) {
	pc, fileName, line, ok := runtime.Caller(depth)
	details := runtime.FuncForPC(pc)

	if ok && details != nil {
		return fileName, details.Name(), line
	}
	return "", "", -1
}

func EscapeCSVdelim(value string) string {
	// out := strings.TrimSpace(value)
	out := strings.ReplaceAll(value, "\t", "\\t")
	out = strings.ReplaceAll(out, "\n", "\\n")
	return out
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

func DateRangesIntersection(rA, rB [2]time.Time) ([2]time.Time, bool) {
	resrange := [2]time.Time{}
	if rA[0].IsZero() && rA[1].IsZero() {
		return rB, true
	}
	if rA[0].After(rB[1]) {
		return resrange, false
	}
	if rA[1].Before(rB[0]) {
		return resrange, false
	}
	var start time.Time
	if rA[0].Before(rB[0]) {
		start = rB[0]
	} else {
		start = rA[0]
	}
	var end time.Time
	if rA[1].Before(rB[1]) {
		end = rA[1]
	} else {
		end = rB[1]
	}
	resrange[0] = start
	resrange[1] = end
	return resrange, true
}

func DateInRange(interval [2]time.Time, dateToCheck time.Time) bool {
	if interval[0].Before(dateToCheck) && interval[1].After(dateToCheck) {
		return true
	}
	if dateToCheck.Equal(interval[0]) {
		return true
	}
	if dateToCheck.Equal(interval[1]) {
		return true
	}
	return false
}

type FileEncodingNumber int

const (
	UNKNOWN FileEncodingNumber = iota
	UTF8
	UTF16le
	UTF16be
)

func ZipFileExtractData(zf *zip.File, enc FileEncodingNumber) ([]byte, error) {
	fileHandle, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()
	var data []byte
	switch enc {
	case UTF8:
		data, err = io.ReadAll(fileHandle)
	case UTF16le:
		utf8reader := enc_unicode.UTF16(enc_unicode.LittleEndian, enc_unicode.IgnoreBOM).NewDecoder().Reader(fileHandle)
		data, err = io.ReadAll(utf8reader)
	default:
		err = fmt.Errorf("unknown encoding")
	}
	return data, err
}

func ZipXmlFileDecodeData(zf *zip.File, enc FileEncodingNumber) (*bytes.Reader, error) {
	data, err := ZipFileExtractData(zf, enc)
	if err != nil {
		return nil, err
	}
	breader := bytes.NewReader(data)
	switch enc {
	case UTF8:
	case UTF16le:
		breader, err = XmlAmendUTF16header(breader)
		if err != nil {
			return nil, err
		}
	default:
		err = fmt.Errorf("unknown encoding")
	}
	return breader, err
}

func PrintRows(rows map[int]CSVrowFields) {
	for i := 0; i < len(rows); i++ {
		fmt.Println(i, rows[i])
		fmt.Println()
	}
}

func JoinObjectPath(oldpath, newpath string) string {
	return oldpath + "/" + newpath
}
