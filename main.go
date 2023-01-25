package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
        "io"
        "archive/zip"
)

// VERSION of openmedia-minify
const VERSION = "0.0.1"

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
			log.Println("Warn: " + err.Error())
		}
	}

	return err

}

func ToXML(input_string string) string {
	var b bytes.Buffer
	xml.EscapeText(&b, []byte(input_string))
	return b.String()
}

// Minify reduces empty fields (whole lines) from XML file
func Minify(inpath string, outpath string, file os.FileInfo) error {

	fext := filepath.Ext(file.Name())

	if file.IsDir() || fext != ".xml" {
		return errors.New("Skipping folder or non-xml file")
	}

	fptr, _ := os.Open(filepath.Join(inpath, file.Name()))
	scanner := bufio.NewScanner(transform.NewReader(fptr, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))

	defer fptr.Close()

	var modded []string
	for scanner.Scan() {
		line := fmt.Sprintln(scanner.Text())

		if strings.Contains(line, `IsEmpty = "yes"`) || line == "" {
			continue
		} else {
			modded = append(modded, line)
		}
	}

	if inpath == outpath {
		return errors.New("This would rewrite the original file. Input and output paths must differ.")
	}

        // TODO: check validity of resulting XML file

	err := saveStringSliceToFile(filepath.Join(outpath, file.Name()), modded)
	if err != nil {
		return errors.New("Failed to save file " + filepath.Join(outpath, file.Name()))
	}


        err = zipFile(filepath.Join(inpath, file.Name()))
        if err != nil {
              return errors.New("Failed to create zip archive: "+filepath.Join(outpath, file.Name()))
        }

	return nil
}


func zipFile(input_filename string) error {
    dir, filename := filepath.Split(input_filename)
    name := strings.TrimSuffix(filename, filepath.Ext(filename))

    archive, err := os.Create(filepath.Join(dir, name +".zip") )
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

// function checks validity of XML
func IsValidXML(input []string) bool {
	return true
}
