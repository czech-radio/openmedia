package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/text/transform"
        "golang.org/x/text/encoding/unicode"
	"ioutil"
        "log"
	"os"
	"path/filepath"
)

// VERSION of openmedia-minify
const VERSION = "0.0.1"

func main() {

	log.Println(fmt.Sprintf("Openmedia-minify version: %s", VERSION))

	input := flag.String("i", "", "The input directory")
	output := flag.String("o", "", "Output directory")
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

}

func ProcessFolder(input string, output string) error {

	files, err := ioutil.ReadDir(input)

	if err != nil {
		return err
	}

	for file := range files {
		err := Minify(file)
		if err != nil {
			continue
		}
	}

	return err

}

func Minify(inpath string, outpath string, file os.Fileinfo) error {

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

		if strings.Contains(`Is Empty="yes"`) {
			continue
		} else {
			modded = append(modded, line)
		}


	}
	// TODO: modded to a file here

        return nil
}
