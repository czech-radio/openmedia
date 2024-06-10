# openmedia

**The console application to archive and extract data from OpenMedia XML files.**

[![build](https://github.com/czech-radio/openmedia-archive/actions/workflows/main.yml/badge.svg)](https://github.com/czech-radio/openmedia-archive/actions/workflows/main.yml) ![version](https://img.shields.io/badge/version-0.9.10-blue.svg) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/a501e03269e1404fa677a0f6cecd7bfe)](https://app.codacy.com/gh/czech-radio/openmedia-archive/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

## Description

Program operates on XML rundown and contact files and creates ZIP archives. Archives will be named like `2023_W49_ORIGINAL.zip` for original files or `2023_W49_MINIFIED.zip` for minified files, where the part `W49` means the ISO week number. Each archive will contain only files corresponding to the same ISO week number. A date and name of the week day is derived from XML time tag. Rundown files in archives are renamed like `RD_05-09_Dvojka_Wednesday_W10_2020_03_04.xml`.

The program executes two operations:

- Archive original files
  - Rundown files archives will be created in `OUTPUT_FOLDER/Rundowns` directory and the archive will be named like `2023_W49_ORIGINAL.zip`
  - Contact files archives will be created in `OUTPUT_FOLDER/Contacts` directory and files will be named also like `2023_W49_ORIGINAL.zip`.

- Minify original files
  - original files will be minified, so that empty fields (fields that do not contain any value) will be removed. Original files are put in archive named like `2023_W49_MINIFIED.zip` after minification.

## Installation

- Linux

  ```bash
  ./scripts/build.sh
  ```

- Windows

  ```powershell
  .\scripts\build.ps1
  ```

## Usage

### Commands

#### Base

- running built binary (see Instalation)

    ```bash
    openmedia -h
    ```

- running without build (see Instalation)

    ```bash
    run main.go -h
    ```

#### Extract Rundown Archvie

    ```bash 
    go run main.go extractArchive -h
    ```

- providing all needed parameters on commandline is rather cumbersome

- use script default variables or change them in run_main.sh:

    -- globally in ArchiveExtractCommand or per preset function

    -- mount archive rundown folder to SOURCE_DIR as specified in run_main.sh

    -- create OUTPUT_DIR as specified in run_manin.sh

    ```bash
    ./scripts/run_main.sh ArchiveExtractOpozice
    ```

- avaiable presests:

    -- ArchiveExtractConntrolProductionHour

    -- ArchiveExtractConntrolProductionWeek

    -- ArchiveExtractControl

    -- ArchiveExtractConntrolProductionDecember

    -- ArchiveExtractControlValidation

    -- ArchiveExtractRange

    -- ArchiveExtractEurovolby

    -- ArchiveExtractOpozice


## Development

- Rundown files structure is describede [here](<https://github.com/czech-radio/openmedia-extract/edit/main/docs/source/notes.md>).

- Additional testing files are located in `R/GŘ/Strategický rozvoj/Analytická sekce/_Archiv/Projekty/OpenMedia/04_03_2020`.

- For XML rundown validation use program [`xmlint`](https://www.root.cz/man/1/xmllint/)[^1]

  ```bash
  xmllint --schema schema.xsd
  ```

[^1]: The XML does not validate when `schema.xsd` imports another XSD for common objects.
