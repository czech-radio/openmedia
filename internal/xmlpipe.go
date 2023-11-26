package internal

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"golang.org/x/text/encoding/unicode"
)

func PipePrint(input_reader *io.PipeReader) {
	var resultBuffer bytes.Buffer
	_, _ = io.Copy(&resultBuffer, input_reader)
	fmt.Println("Contents of PipeReader:", resultBuffer.String())
}

func PipeUTF16leToUTF8(r io.Reader) *io.PipeReader {
	utf8reader := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder().Reader(r)
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		_, _ = io.Copy(pw, utf8reader)
	}()
	return pr
}

func PipeRundownheaderAmmend(input_reader io.Reader) *io.PipeReader {
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
				_, _ = writer.WriteString(line + "\n")
			} else {
				_, _ = writer.WriteString(line + "\n")
			}
		}
		// Write remainig bytes wihtout scanning
		// _, _ = io.Copy(writer, input_reader)
		err := writer.Flush()
		if err != nil {
			slog.Error(err.Error())
		}
	}()
	return pr
}

func PipeRundownMinfiy(input_reader *io.PipeReader) (*OPENMEDIA, error) {
	var OM OPENMEDIA
	bufio.NewReader(input_reader)
	// byteData, err := io.ReadAll(input_reader)
	// if err != nil {
	// return nil, err
	// }
	err := xml.NewDecoder(input_reader).Decode(&OM)
	if err != nil {
		return nil, err
	}
	// fmt.Println("ke", byteData)
	// err = xml.Unmarshal(byteData, &OM)
	// io.ReadFull
	// err = xml.Unmarshal(byteData, &OM)
	// if err != nil {
	// return nil, err
	// }
	return &OM, nil
}

func RundownMarshal(om *OPENMEDIA, dst_file string) error {
	res, err := xml.MarshalIndent(om, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}
