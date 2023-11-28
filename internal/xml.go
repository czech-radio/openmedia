package internal

import (
	"embed"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"

	// "encoding/xml/schema"

	"golang.org/x/text/encoding/unicode"
	// "gorm.io/gorm/schema"
)

var content embed.FS

func EmbededOpenMediaXSD() ([]byte, error) {
	fmt.Println(os.Getwd())
	filePath := filepath.Join("../", "test", "testdata", "rundowns_schemas", "OM_LV7_schema.xsd")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fmt.Println(bytes)

	// data, err := content.ReadFile(filePath)
	return bytes, err
}

func XmlValidateWithXSD(filePath string) (bool, error) {
	xsdData, err := EmbededOpenMediaXSD()
	if err != nil {
		return false, err
	}
	fmt.Println(xsdData)
	// schema := &schema.Schema{}
	// if err := schema.Unmarshal(xsdData); err != nil {
	// return false, fmt.Errorf("error unmarshaling XSD: %v", err)
	// }
	return true, nil
}

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
