package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/charmap"
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

func XmlTagAttributesStruct(xs xml.StartElement, v any) {
}

func FileIsValidXmlToMinify(src_file_path string) (bool, error) {
	file_extension := filepath.Ext(src_file_path)
	if file_extension != ".xml" {
		return false,
			fmt.Errorf("file does not have xml extension: %s", src_file_path)
	}
	if !strings.Contains(src_file_path, "RD") {
		return false,
			fmt.Errorf("filename does not contaion 'RD' string")
	}
	srcFile, err := os.Open(src_file_path)
	if err != nil {
		return false, err
	}
	_, err = srcFile.Stat()
	if err != nil {
		return false, err
	}
	return true, nil
}

func BypassReader(label string, input io.Reader) (io.Reader, error) {
	return input, nil
}

func XmlDecoderValidate(decoder *xml.Decoder) (bool, error) {
	for {
		err := decoder.Decode(new(interface{}))
		if err != nil {
			return err == io.EOF, err
		}
	}
}

func XmlFileLinesValidate2(src_file_path string) (bool, error) {
	//##NOTE: DT:2023/11/12_20:13:10, Provide XML schema?
	file, err := os.Open(src_file_path)
	if err != nil {
		return false, err
	}
	defer file.Close()
	reader := charmap.Windows1252.NewDecoder().Reader(file)
	decoder := xml.NewDecoder(reader)
	return XmlDecoderValidate(decoder)
}

func XmlFileLinesValidate(src_file_path string) (bool, error) {
	// ##NOTE: DT:2023/11/12_20:13:10, Provide XML schema?
	file, err := os.Open(src_file_path)
	if err != nil {
		return false, err
	}
	defer file.Close()
	reader := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder().Reader(file)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = BypassReader
	return XmlDecoderValidate(decoder)
}

// func XMLunmarshallDynamicElement(d *xml.Decoder, start xml.StartElement, attrs_map map[string]string) {
// var tagValue strings.Builder
// var start_count int = 0
// var end_count int = 0
// var errUnexpectedTagStructure = fmt.Errorf("unexpected xml tag structure")
// }

func RundownUnmarshall(file_path string) (*OPENMEDIA, error) {
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
