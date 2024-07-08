# openmedia

**The console application to archive and extract data from OpenMedia XML files.**

[![build](https://github.com/czech-radio/openmedia-archive/actions/workflows/main.yml/badge.svg)](https://github.com/czech-radio/openmedia-archive/actions/workflows/main.yml) ![version](https://img.shields.io/badge/version-1.0.1-blue.svg) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/a501e03269e1404fa677a0f6cecd7bfe)](https://app.codacy.com/gh/czech-radio/openmedia-archive/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

## Description

- Main features of this program is to archive and extract Openmedia XML rundown files.

- Weekly production of Openmedia rundowns files tends to be lots of data (5 GB/week) so with this program you can create zip archives from produced files, such that files which has same ISO week date are nested in one zip archive.  Resulting archive named for example `2023_W49_ORIGINAL.zip` stores exact copy of original rundown files has 10:1 size reduction. Resulting archive named for example `2023_W49_MINIFIED.zip` stores minified files and which has removed blocks of XML code which does not store any meaningful value. Rundown files in archives are renamed like `RD_05-09_Dvojka_Wednesday_W10_2020_03_04.xml`. The date and name of the week day in resulting filename is derived from XML time tag.

- The size reduction of minified files is 30:1. Minified archive/files are much faster to process or download.

- Next step is to extract, process and output useful data to csv (xlsx) table. The program contains various options for preprocessing, validation, transformation and filtering of data.

## Installation

- Linux

  ```bash
  ./scripts/build.sh
  ```

- Windows

  ```powershell
  .\scripts\build.ps1
  ```

## Help

[help](./docs/HELP.md)

## Usage

[usage](./docs/USAGE.md)

## Presets script

[presets](./scripts/run_main.sh)

  ```bash
  ./scripts/run_main.sh ArchiveExtractControl
  ```

  ```bash
  ./scripts/run_main.sh ArchiveExtractControlValidation
  ```

- make copy of the script, change main variables inside preset function form example: FROM, TO and FILTER_FILE.

- Make sure the output directory exists or change it in script inside function ArchiveExtractCommand OUTPUT_DIR.

- Mount rundowns folder R/.../cro/export-avo/Rundowns. Change variable SOURCE_DIR on line 3 to folder where the rundowns resides.

- Run the script.

    ```bash
  ./scripts/run_main.sh ArchiveExtractOpozice
  ```

  ```bash
  ./scripts/run_main.sh ArchiveExtractEurovolby
  ```

## Development

- Rundown files structure is described [here](<https://github.com/czech-radio/openmedia-extract/edit/main/docs/source/notes.md>).

- For XML rundown validation use program [`xmlint`](https://www.root.cz/man/1/xmllint/)[^1]

  ```bash
  xmllint --schema schema.xsd
  ```

[^1]: The XML does not validate when `schema.xsd` imports another XSD for common objects.
