package internal

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/go-xmlfmt/xmlfmt"
	"golang.org/x/text/encoding/unicode"
)

func PipeConsume(input_reader *io.PipeReader) {
	var resultBuffer bytes.Buffer
	_, err := io.Copy(&resultBuffer, input_reader)
	Errors.ExitWithCode(err)
}

func PipePrint(input_reader *io.PipeReader) {
	var resultBuffer bytes.Buffer
	_, err := io.Copy(&resultBuffer, input_reader)
	Errors.ExitWithCode(err)
	fmt.Println(resultBuffer.String())
}

func PipeUTF16leToUTF8(r io.Reader) *io.PipeReader {
	utf8reader := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder().Reader(r)
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		_, err := io.Copy(pw, utf8reader)
		Errors.ExitWithCode(err)
	}()
	return pr
}

func XmlFindBaseOpenMediaNode(breader *bytes.Reader,
) (*xmlquery.Node, error) {
	// Parse base xml node
	baseNode, err := xmlquery.Parse(breader)
	if err != nil {
		return nil, err
	}
	nodes := xmlquery.Find(baseNode, "/OPENMEDIA")
	if len(nodes) != 1 {
		return nil, fmt.Errorf(
			"unknown opendmedia file, nodes found count: %d,should be 1",
			len(nodes),
		)
	}
	return nodes[0], nil
}

func XmlAmendUTF16header(breader *bytes.Reader) (*bytes.Reader, error) {
	var buf bytes.Buffer
	var err error
	replace := "encoding=\"UTF-16\""
	scanner := bufio.NewScanner(breader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, replace) {
			line = strings.Replace(line, replace, "encoding=\"UTF-8\"", 1)
			_, err = fmt.Fprintln(&buf, line)
		} else {
			_, err = fmt.Fprintln(&buf, line)
		}
		if err != nil {
			return nil, err
		}
	}
	return bytes.NewReader(buf.Bytes()), nil
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
		Errors.ExitWithCode(err)
		// Write remainig bytes wihtout scanning?
		err = writer.Flush()
		Errors.ExitWithCode(err)

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
		Errors.ExitWithCode(err)
		_, err = io.Copy(writer, buffReader)
		Errors.ExitWithCode(err)
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
	// if OM == nil {
	// return nil, fmt.Errorf("xml cannot be unmarshaled")
	// }
	return &OM, nil
}

func PipeRundownMarshal(om *OPENMEDIA) *io.PipeReader {
	pr, pw := io.Pipe()
	writer := bufio.NewWriter(pw)
	go func() {
		defer pw.Close()
		xmlBytes, err := xml.MarshalIndent(om, "", "  ")
		Errors.ExitWithCode(err)
		_, err = writer.Write(xmlBytes)
		Errors.ExitWithCode(err)
		writer.Flush()
	}()
	return pr
}

type ArchiveResult struct {
	FilesCount     int
	FilesProcessed int
	FilesSuccess   int
	FilesFailure   int
	Errors         []error
	FilesValid     []string
}

type OpenMediaFileTypeCode int

const (
	OmFileTypeRundown = iota
	OmFileTypeContact
)

type OpenMediaFileType struct {
	Code         OpenMediaFileTypeCode
	ShortHand    string
	TemplateName string
	OutputDir    string
}

var OpenMediaFileTypeMap = map[OpenMediaFileTypeCode]*OpenMediaFileType{
	OmFileTypeRundown: {
		OmFileTypeRundown, "RD", "Radio Rundown", "Rundowns"},
	OmFileTypeContact: {
		OmFileTypeContact, "CT", "Contact Bin", "Contacts"},
}

func GetOMtypeByTemplateName(templateName string) (*OpenMediaFileType, error) {
	var result *OpenMediaFileType
	for _, t := range OpenMediaFileTypeMap {
		if t.TemplateName == templateName {
			result = t
			return result, nil
		}
	}
	return result, fmt.Errorf("unknown teplate type: %s", templateName)
}

func ValidateFileName(src_path string) (bool, error) {
	file_extension := filepath.Ext(src_path)
	if file_extension != ".xml" {
		return false,
			fmt.Errorf("file does not have xml extension: %s", src_path)
	}
	var isOpenMediaFile bool
	for _, t := range OpenMediaFileTypeMap {
		if strings.Contains(src_path, t.ShortHand) {
			isOpenMediaFile = true
		}
	}
	if !isOpenMediaFile {
		return false, fmt.Errorf("filename is not valid OpenMedia file: %s", src_path)
	}

	_, err := os.Stat(src_path)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ValidateFilesInDirectory(rootDir string, recursive bool) (*ArchiveResult, error) {
	var result *ArchiveResult = &ArchiveResult{}
	dirWalker := func(filePath string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filePath == rootDir {
			return nil
		}
		if file.IsDir() && !recursive {
			return filepath.SkipDir
		}
		if file.IsDir() {
			return nil
		}
		_, err = ValidateFileName(filePath)
		if err != nil {
			result.FilesFailure++
			result.AddError(err)
			return nil
		}
		result.FilesProcessed++
		result.FilesValid = append(result.FilesValid, filePath)
		result.FilesCount = result.FilesProcessed
		return nil
	}

	err := filepath.WalkDir(rootDir, dirWalker)
	if err != nil {
		return result, err
	}
	if len(result.Errors) > 0 {
		err := fmt.Errorf("%s, count %d", Errors.CodeMsg(ErrCodeInvalid), len(result.Errors))
		return result, err
		// errors.New("invalid files count: %d", len(result.Errors))
	}
	return result, nil
}

func (ar *ArchiveResult) AddError(err ...error) {
	if err == nil {
		return
	}
	if len(err) > 0 {
		ar.Errors = append(ar.Errors, err...)
	}
}

func XMLprint(node *xmlquery.Node) {
	ex := xmlfmt.FormatXML(node.OutputXML(true), "", "\t")
	fmt.Println(ex)
}

func GetFieldValueByName(attrs []xmlquery.Attr, id string) (string, bool) {
	for _, i := range attrs {
		if i.Name.Local == id {
			return i.Value, true
		}
	}
	return "NO_VALUE", false
}

func XMLbuildAttrQuery(attrName string, ids []string) string {
	if len(ids) == 0 {
		return ""
	}
	if ids[0] == "*" {
		return "/*"
	}
	// if ids[0] == "" {
	// return attrName
	// }
	var expr strings.Builder
	attrQuery := "@" + attrName + "='"
	expr.WriteString("[")
	for i, id := range ids {
		if i != len(ids)-1 {
			expr.WriteString(attrQuery + id + "' or ")
		} else {
			// expr.WriteString("@FieldID='" + id + "']@FieldID")
			expr.WriteString(attrQuery + id + "']")
		}
	}
	return expr.String()
}

func GetPathGlobPrefix(objectName string) (string, string) {
	pathPrefix := "/"
	pattern := "^\\*"
	regex := regexp.MustCompile(pattern)
	parts := regex.Split(objectName, -1)
	if len(parts) > 1 {
		objectName = parts[1]
		pathPrefix = "//"
	}
	return objectName, pathPrefix
}

// /Radio Rundown/<OM_RECORD>/Hourly Rundown"

func XMLqueryFromPath(path string) string {
	var out strings.Builder
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "" {
			continue
		}
		object, globPrefix := GetPathGlobPrefix(part)
		attrName, ok := OmTagStructureMap[object]
		if ok {
			fmt.Fprintf(
				&out, "%s%s", globPrefix, attrName.XMLtagName)
			continue
		}
		attrQuery := XMLbuildAttrQuery(
			"TemplateName", []string{object})
		fmt.Fprintf(
			&out, "%sOM_OBJECT%s", globPrefix, attrQuery)
	}
	return out.String()
}

func QueryObject(objectName string) (string, error) {
	//ObjectName: <OM_RECORD>, "Radio Rundow"
	var XMLobjectName string
	var XMLattrValue string
	var attrquery string
	var pathPrefix string

	// Check glob path
	objectName, pathPrefix = GetPathGlobPrefix(objectName)
	xmlTag, ok := OmTagStructureMap[objectName]
	if ok {
		XMLobjectName = xmlTag.XMLtagName
	}
	if !ok {
		XMLobjectName = "OM_OBJECT"
		XMLattrName := "TemplateName"
		XMLattrValue = objectName
		attrquery = XMLbuildAttrQuery(
			XMLattrName, []string{XMLattrValue})
	}
	objquery := pathPrefix + XMLobjectName + attrquery
	return objquery, nil
}

func XMLqueryFields(fieldsPath string, IDs []string) string {
	attrQuery := XMLbuildAttrQuery("FieldID", IDs)
	return fieldsPath + attrQuery
}
