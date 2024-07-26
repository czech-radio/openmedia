# openmedia

**The console application to archive and extract data from OpenMedia XML files.**

[![build](https://github.com/czech-radio/openmedia-archive/actions/workflows/main.yml/badge.svg)](https://github.com/czech-radio/openmedia-archive/actions/workflows/main.yml) ![version](https://img.shields.io/badge/version-1.0.1-blue.svg) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/a501e03269e1404fa677a0f6cecd7bfe)](https://app.codacy.com/gh/czech-radio/openmedia-archive/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

## Description

- The main features of this program are to archive and extract Openmedia XML rundown files.

- Weekly production of Openmedia rundown files tends to be lots of data (5 GB/week). With this program, you can create zip archives from rundown files, such that files with the same ISO week date are nested in one ZIP archive. Two types of archives are produced: original and minimal.

- Rundown files stored in both original and modified archives are renamed, like `RD_05-09_Dvojka_Wednesday_W10_2020_03_04.xml`. The date and name of the weekday in the resulting filename are derived from the XML time tag.

- The original archive, e.g., '2023_W49_ORIGINAL.zip', stores an exact copy of the original rundown files, and after compression, the size is reduced by a factor of 10. 2023_W49 stands for the year 2023 and ISO week number 49. 

- A minified archive named, for example, `2023_W49_MINIFIED.zip` stores minified files. Minified files have removed blocks of XML data that do not store any meaningful value. The size of minified files is reduced after minification and compression by a factor of ~30. Minified archives and files are much faster to process or download. 

- The XML rundown files make it difficult to process and analyze the data they contain. So the next step is to extract, process, and output useful data to a CSV (XLSX) table. The program contains various options for preprocessing, validation, transformation, and filtering of data.

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

- Make a copy of the script and change the main variables inside the preset function, for example, FROM, TO, and FILTER_FILE.

- Make sure the output directory exists, or change it in the script inside the function ArchiveExtractCommand OUTPUT_DIR.

- Mount the rundowns folder in `R/.../cro/export-avo/Rundowns`. Change the variable SOURCE_DIR on line 3 to the folder where the rundowns reside.

- Run the script.

    ```bash
  ./scripts/run_main.sh ArchiveExtractOpozice
  ```

  ```bash
  ./scripts/run_main.sh ArchiveExtractEurovolby
  ```

## Development

- Rundown files structure is described [here](<https://github.com/czech-radio/extractor/blob/main/docs/source/notes.md>).

- For XML rundown validation use program [`xmlint`](https://www.root.cz/man/1/xmllint/)[^1]

  ```bash
  xmllint --schema schema.xsd
  ```

[^1]: The XML does not validate when `schema.xsd` imports another XSD for common objects.
