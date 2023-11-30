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
	fmt.Println("Contents of PipeReader:", resultBuffer.String())
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
