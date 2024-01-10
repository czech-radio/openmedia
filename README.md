# openmedia-archive

**The console application archive OpenMedia XML files.**

[![build](https://github.com/czech-radio/openmedia-reduce/actions/workflows/main.yml/badge.svg)](https://github.com/czech-radio/openmedia-reduce/actions/workflows/main.yml) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/a501e03269e1404fa677a0f6cecd7bfe)](https://app.codacy.com/gh/czech-radio/openmedia-archive/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

## Description

Program operates on XML rundown and contact files and creates ZIP archives. Archives will be named like `2023_W49_ORIGINAL.zip` for original files or `2023_W49_MINIFIED.zip` for minified files, where the part `W49` means the ISO week number. Each archive will contain only files corresponding to the same ISO week number. A date and name of the week day is derived from XML time tag. Rundown files in archives are renamed like `RD_05-09_Dvojka_Wednesday_W10_2020_03_04.xml`.

The program executes two operations:

- Archive original files
  - Rundown files archives will be created in `OUTPUT_FOLDER/Rundowns` directory and the archive will be named like `2023_W49_ORIGINAL.zip`
  - Contact files archives will be created in `OUTPUT_FOLDER/Contacts` directory and files will be named also like `2023_W49_ORIGINAL.zip`.

- Minify original files
  - original files will be minified, so that empty fields (fields that do not contain any value) will be removed. Original files are put in archive named like `2023_W49_MINIFIED.zip` after minification.

## Build

- linux
  ```bash
  ./scripts/build.sh
  ```

- win
  ```ps
  ./scripts/build.ps1
  ```

## Usage

Use `openmedia-archive -h` to see all available options.

- Production mode: Halts when unprocessable file encountered.

  ```bash
  openmedia-archive -i <SOURCE_FOLDER> -o <OUTPUT_FOLDER>
  ```

- Dry run mode: output files will be created in a temporary directory

  ```bash
  openmedia-archive -n -i <SOURCE_FOLDER> [-o <OUTPUT_FOLDER>]
  ```

- Continue processing folder when unprocessable file encountered
  (useful for example in dry-mode).
  
  ```bash
  openmedia-archive -ifr -i <SOURCE_FOLDER> -o <OUTPUT_FOLDER>
  ```

  ```bash
  openmedia-archive -n -ifr -i <SOURCE_FOLDER> [-o <OUTPUT_FOLDER>]
  ```

- By default, we log in structured JSON e.g.

  ```bash
  {"time":"2024-01-08T20:08:41.365154585+01:00","level":"INFO","source":{"function":"github/czech-radio/openmedia-archive/internal.(*Process).WorkerLogInfo","file":"process.go","line":346},"msg":"2020_W10_MINIFIED.zip","ArhiveRatio":"0.024","MinifyRatio":"0.327","original":13552180,"compressed":319488,"minified":4430320,"file":"test/testdata/rundowns_mix/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml"}
  {"time":"2024-01-08T20:08:41.519854699+01:00","level":"INFO","source":{"function":"github/czech-radio/openmedia-archive/internal.(*Process).WorkerLogInfo","file":"process.go","line":346},"msg":"2020_W10_ORIGINAL.zip","ArhiveRatio":"0.066","MinifyRatio":"1.000","original":18364782,"compressed":1204224,"minified":18364782,"file":"test/testdata/rundowns_mix/RD_12-19_ČRo_Ostrava_-_Středa_04_03_2020_2_1603282_20200304234540.xml"}
  {"time":"2024-01-08T20:08:42.244667764+01:00","level":"INFO","source":{"function":"github/czech-radio/openmedia-archive/internal.(*Process).WorkerLogInfo","file":"process.go","line":346},"msg":"2020_W10_MINIFIED.zip","ArhiveRatio":"0.023","MinifyRatio":"0.314","original":18364782,"compressed":421888,"minified":5772375,"file":"test/testdata/rundowns_mix/RD_12-19_ČRo_Ostrava_-_Středa_04_03_2020_2_1603282_20200304234540.xml"}
  {"time":"2024-01-08T20:08:42.244735372+01:00","level":"INFO","source":{"function":"github/czech-radio/openmedia-archive/internal.(*Process).WorkerLogInfo","file":"process.go","line":346},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.063","MinifyRatio":"1.000","original":449249600,"compressed":28196864,"minified":449249600,"file":"test/testdata/rundowns_mix/"}
  {"time":"2024-01-08T20:08:42.244753922+01:00","level":"INFO","source":{"function":"github/czech-radio/openmedia-archive/internal.(*Process).WorkerLogInfo","file":"process.go","line":346},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.021","MinifyRatio":"0.287","original":449249600,"compressed":9629696,"minified":129017125,"file":"test/testdata/rundowns_mix/"}
  ```

- Plain logging output can be invoked with `-lt` e.g.

  ```bash
  openmedia_archive -n -lt plain -i <SOURCE_FOLDER> [-o <OUTPUT_FOLDER>]
  ```

  which outputs

  ```bash
  time=2024-01-08T20:05:22.542+01:00 level=INFO source=process.go:346 msg=2020_W10_ORIGINAL.zip ArhiveRatio=0.067 MinifyRatio=1.000 original=13552180 compressed=905216 minified=13552180 file=test/testdata/rundowns_mix/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml
  time=2024-01-08T20:05:23.457+01:00 level=INFO source=process.go:346 msg=2020_W10_MINIFIED.zip ArhiveRatio=0.024 MinifyRatio=0.327 original=13552180 compressed=319488 minified=4430320 file=test/testdata/rundowns_mix/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml
  time=2024-01-08T20:05:23.662+01:00 level=INFO source=process.go:346 msg=2020_W10_ORIGINAL.zip ArhiveRatio=0.066 MinifyRatio=1.000 original=18364782 compressed=1204224 minified=18364782 file=test/testdata/rundowns_mix/RD_12-19_ČRo_Ostrava_-_Středa_04_03_2020_2_1603282_20200304234540.xml
  time=2024-01-08T20:05:24.500+01:00 level=INFO source=process.go:346 msg=2020_W10_MINIFIED.zip ArhiveRatio=0.023 MinifyRatio=0.314 original=18364782 compressed=421888 minified=5772375 file=test/testdata/rundowns_mix/RD_12-19_ČRo_Ostrava_-_Středa_04_03_2020_2_1603282_20200304234540.xml
  time=2024-01-08T20:05:24.500+01:00 level=INFO source=process.go:346 msg=GLOBAL_ORIGINAL ArhiveRatio=0.063 MinifyRatio=1.000 original=449249600 compressed=28196864 minified=449249600 file=test/testdata/rundowns_mix/
  time=2024-01-08T20:05:24.500+01:00 level=INFO source=process.go:346 msg=GLOBAL_MINIFY ArhiveRatio=0.021 MinifyRatio=0.287 original=449249600 compressed=9629696 minified=129017125 file=test/testdata/rundowns_mix/
  ```


## Development

- Rundown files structure is describede [here](<https://github.com/czech-radio/openmedia-extract/edit/main/docs/source/notes.md>).

- Additional testing files are located in `R/GŘ/Strategický rozvoj/Analytická sekce/_Archiv/Projekty/OpenMedia/04_03_2020`.

- For XML rundown validation use program [`xmlint`](https://www.root.cz/man/1/xmllint/)[^1]

  ```bash
  xmllint --schema schema.xsd
  ```

[^1]: The XML does not validate when `schema.xsd` imports another XSD for common objects.
