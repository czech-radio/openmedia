package internal

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"golang.org/x/text/encoding/unicode"
)

func PipePrint(input_reader *io.PipeReader) {
	var resultBuffer bytes.Buffer
	_, err := io.Copy(&resultBuffer, input_reader)
	ErrorExitWithCode(err)
	fmt.Println(resultBuffer.String())
}

func PipeUTF16leToUTF8(r io.Reader) *io.PipeReader {
	utf8reader := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder().Reader(r)
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		_, err := io.Copy(pw, utf8reader)
		ErrorExitWithCode(err)
	}()
	return pr
}

func PipeRundownHeaderAmmend(input_reader io.Reader) *io.PipeReader {
	var err error
	pr, pw := io.Pipe()
	writer := bufio.NewWriter(pw)
	scanner := bufio.NewScanner(input_reader)
	go func() {
		defer pw.Close()
		replace := "encoding=\"UTF-16\""
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, replace) {
				line = strings.Replace(line, replace, "encoding=\"UTF-8\"", 1)
				_, err = writer.WriteString(line + "\n")
			} else {
				_, err = writer.WriteString(line + "\n")
			}
		}
		ErrorExitWithCode(err)
		// Write remainig bytes wihtout scanning?
		// _, _ = io.Copy(writer, input_reader)
		err = writer.Flush()
		ErrorExitWithCode(err)

	}()
	return pr
}

func PipeRundownHeaderAdd(input_reader io.Reader) *io.PipeReader {
	pr, pw := io.Pipe()
	buffReader := bufio.NewReader(input_reader)
	writer := bufio.NewWriter(pw)
	go func() {
		defer pw.Close()
		defer writer.Flush()
		_, err := writer.Write(openMediaXmlHeader)
		ErrorExitWithCode(err)
		_, err = io.Copy(writer, buffReader)
		ErrorExitWithCode(err)
	}()
	return pr
}

func PipeRundownUnmarshal(input_reader *io.PipeReader) (*OPENMEDIA, error) {
	var OM OPENMEDIA
	buffReader := bufio.NewReader(input_reader)
	// io.ReadFull
	// byteData, err := io.ReadAll(input_reader)
	// err = xml.Unmarshal(byteData, &OM)
	err := xml.NewDecoder(buffReader).Decode(&OM)
	if err != nil {
		return nil, err
	}
	return &OM, nil
}

func PipeRundownMarshal(om *OPENMEDIA) *io.PipeReader {
	pr, pw := io.Pipe()
	writer := bufio.NewWriter(pw)
	go func() {
		defer pw.Close()
		xmlBytes, err := xml.MarshalIndent(om, "", "\t")
		ErrorExitWithCode(err)
		_, err = writer.Write(xmlBytes)
		ErrorExitWithCode(err)
		writer.Flush()
	}()
	return pr
}

type ProcessFolderOptions struct {
	SourceDirectory        string
	DestinationDirectory   string
	InputEncoding          string
	OutputEncoding         string
	ValidateFilenames      bool
	ValidateWithDefaultXSD bool   // validate with bundled file
	ValidateWithXSD        string // path to XSD file
	ValidatePre            bool
	ValidatePost           bool
	ArchiveType            string
	InvalidFileRename      bool
	InvalidFileContinue    bool
}

type ProcessResults struct {
	Weeks        int
	Files        int
	SizeOriginal int
	SizeBackup   int
	SizeMinified int
	Errors       []error
}

func (pr *ProcessResults) AddError(err ...error) {
	pr.Errors = append(pr.Errors, err...)
}

func ProcesFolder(opts ProcessFolderOptions) (*ProcessResults, error) {
	var results = &ProcessResults{}
	if opts.ValidateFilenames {
		result, err := ValidateFilenamesInDirectory(opts.SourceDirectory)
		if err != nil && opts.InvalidFileContinue {
			results.AddError(err)
			results.AddError(result.Errors...)
			return results, err
		}
	}
	return results, nil
}

func ProcessWeek() {
}

// new_filename := beginning + fmt.Sprintf("%s_W%02d_%04d_%02d_%02d", weekday, week, year, month, day)
