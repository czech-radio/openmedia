package internal

import (
	"encoding/xml"
	"io"
	"os"

	"golang.org/x/text/encoding/unicode"
)

func XmlTagAttributesMap(
	xs xml.StartElement,
	attrs_names map[string]string) map[string]string {
	result := make(map[string]string, len(attrs_names))
	for _, attr := range xs.Attr {
		result[attr.Name.Local] = attr.Value
	}
	return result
}

func FileUTF16leToUTF8(src_file, dst_file string) error {
	srcFile, err := os.Open(src_file)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	srcReader := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder().Reader(srcFile)

	dstWriter, err := os.Create(dst_file)
	if err != nil {
		return err
	}
	defer dstWriter.Close()
	_, err = io.Copy(dstWriter, srcReader)
	return err
}

func RundownUnmarshal(file_path string) (*OPENMEDIA, error) {
	xmlFile, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	byteData, err := io.ReadAll(xmlFile)
	var OM OPENMEDIA
	err = xml.Unmarshal(byteData, &OM)
	if err != nil {
		return nil, err
	}
	return &OM, nil
}
