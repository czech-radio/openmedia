package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/terminalstatic/go-xsd-validate"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// VERSION of openmedia-minify
const VERSION = "0.0.2"

func main() {

	log.Println(fmt.Sprintf("Openmedia-minify version: %s", VERSION))

	input := flag.String("i", "", "The input directory")
	output := flag.String("o", "", "The output directory")
	flag.Parse()

	if *input == "" {
		log.Fatal("Please specify the input folder.")
		os.Exit(1)
	}

	if *output == "" {
		log.Fatal("Please specify the output folder.")
		os.Exit(1)
	}

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("openmedia-minify -i input_folder -o output_folder")
	}

	ProcessFolder(*input, *output)

}

// ProcessFolder executes minify on each file in given folder outputs result to output folder
func ProcessFolder(input string, output string) error {

	files, err := ioutil.ReadDir(input)

	if err != nil {
		return err
	}

	for _, file := range files {
		err := Minify(input, output, file)
		if err != nil {
			log.Println(err.Error())
		}
	}

	return err

}

// ToXML helper function converts string to XML escaped string
func ToXML(input_string string) string {
	var b bytes.Buffer
	xml.EscapeText(&b, []byte(input_string))
	return b.String()
}

// Minify reduces empty fields (whole lines) from XML file
func Minify(inpath string, outpath string, file os.FileInfo) error {

	fext := filepath.Ext(file.Name())

	if file.IsDir() || fext != ".xml" || !strings.Contains(file.Name(), "RD") {
		return errors.New("Skipping folder, non-XML file or non-RD file")
	}

	fptr, _ := os.Open(filepath.Join(inpath, file.Name()))
	scanner := bufio.NewScanner(transform.NewReader(fptr, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))

	defer fptr.Close()

	var modded []string
        var counter int = 0
        var skipped int = 0
        for scanner.Scan() {
		line := fmt.Sprintln(scanner.Text())

		if (strings.Contains(line, `IsEmpty = "yes"`) && strings.Contains(line, "OM_FIELD")) || strings.Contains(line, "/OM_RECORD") {
			//log.Println("skipping: "+line)
			skipped++
                        continue
		} else {

		  // FIX encoding line to UTF-8
		  if counter == 0 {
                    line = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>`
		  }
                  counter++

                  modded = append(modded, line)
		}

	}

	log.Println("Minifying: " + filepath.Join(inpath, file.Name()))
        log.Printf("Document minified from %04d lines to %04d lines\n",skipped+counter,counter)

	// TODO: check validity of resulting XML file

	// new filename
	weekday, year, month, day, week := getDateFromFile(filepath.Join(inpath, file.Name()))
	split := strings.Split(file.Name(), "-")
	beginning := split[0] + "-" + split[1]
	new_filename := beginning + fmt.Sprintf("%s_W%02d_%04d_%02d_%02d", weekday, week, year, month, day)

	err := saveStringSliceToFile(filepath.Join(outpath, new_filename+".xml"), modded)
	if err != nil {
		return errors.New("Failed to save file " + filepath.Join(outpath, new_filename+".xml"))
	}

	log.Println("Validating source file: " + filepath.Join(inpath, file.Name()))
	err = IsValidXML(filepath.Join(inpath, file.Name()))
	if err != nil {
		return errors.New("Source file is invalid: " + filepath.Join(inpath, file.Name()) + " " + err.Error())
	}

	log.Println("Validating destination file: " + filepath.Join(outpath, new_filename))
	err = IsValidXML(filepath.Join(outpath, new_filename+".xml"))
	if err != nil {
		return errors.New("Resulting file is invalid: " + filepath.Join(outpath, new_filename+".xml") + " " + err.Error())
	}

	err = zipFile(filepath.Join(inpath, file.Name()), filepath.Join(outpath, new_filename+".zip"))
	if err != nil {
		return errors.New("Failed to create zip archive: " + filepath.Join(outpath, new_filename+".zip"))
	}

	return nil
}

// openmedia-check function to get date from xml file
func getDateFromFile(filepath string) (Weekday string, Year, Month, Day, Week int) {

	var weekday string
	var year, month, day, week = 0, 0, 0, 0
	var scanner bufio.Scanner

	handle, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	scanner = *bufio.NewScanner(transform.NewReader(handle, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))

	for scanner.Scan() {
		var line = fmt.Sprintln(scanner.Text())

		if strings.Contains(line, `FieldID = "1004"`) {
			reg := regexp.MustCompile("([0-9][0-9][0-9][0-9]{1})([0-9]{2})([0-9]{2})(T)")
			res := reg.FindStringSubmatch(line)

			date, err := time.Parse("20060102", res[1]+res[2]+res[3])

			if err != nil {
				log.Fatal(err)
			}

			year, month, day = date.Year(), int(date.Month()), date.Day()
			year, week = date.ISOWeek()

			t, err := time.Parse(time.RFC3339, fmt.Sprintf("%04d-%02d-%02dT00:00:00Z", date.Year(), int(date.Month()), date.Day()))
			if err != nil {
				log.Fatal(err)
			}
			weekday = t.Weekday().String()
			break // Find first ocurrance!
		}
	}

	return weekday, year, month, day, week
}

// zipFile zips incoming file to a new zipfile
func zipFile(input_filename string, output_filename string) error {

	_, filename := filepath.Split(input_filename)
	//name := strings.TrimSuffix(filename, filepath.Ext(filename))

	archive, err := os.Create(output_filename)
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	// reader
	f, err := os.Open(input_filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	log.Println("Zipping: " + output_filename)
	// writer
	w, err := zipWriter.Create(filename)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w, f); err != nil {
		panic(err)
	}

	zipWriter.Close()

	return nil
}

// saveStringSliceToFile saves given string slice to a file
func saveStringSliceToFile(filename string, input []string) error {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range input {
		_, _ = datawriter.WriteString(data)
	}

	datawriter.Flush()
	file.Close()

	return err
}

// IsValidXML checks validity of an XML
func IsValidXML(inputFile string) error {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()

	xmlFile, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	inXml, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}

	_, err = xsdvalidate.NewXmlHandlerMem(inXml, xsdvalidate.ValidErrDefault)
	if err != nil {
		switch err.(type) {
		case xsdvalidate.ValidationError:
			log.Println(err)
			log.Printf("Error in line: %d\n", err.(xsdvalidate.ValidationError).Errors[0].Line)
			log.Println(err.(xsdvalidate.ValidationError).Errors[0].Message)
			return err
		default:
			//log.Println(err)
			return err
		}
	}

	return err
}
