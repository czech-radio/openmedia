# OPENMEDIA-ARCHIVE

**The console application archive OpenMedia XML files.**

[![build](https://github.com/czech-radio/openmedia-reduce/actions/workflows/main.yml/badge.svg)](https://github.com/czech-radio/openmedia-reduce/actions/workflows/main.yml) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/a501e03269e1404fa677a0f6cecd7bfe)](https://app.codacy.com/gh/czech-radio/openmedia-archive/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)


## Description
Program operates on xml Rundown and Contact files and creates zip archives from them. Archives will be named like 2023_W49_ORIGINAL.zip for original files or 2023_W49_MINIFIED.zip for minified original files. W49 means the iso week number. Each archive will contain only files corresponding to same iso week number.  The date and name-day is derived from XML time tag and it is the same as given folder on `ANNOVA` disk. Rundown files in archives are renamed like 'RD_05-09_Dvojka_Wednesday_W10_2020_03_04.xml'.


1.  archivate original files
-   Rundown files archives will be created in 'OUTPUT_FOLDER/Rundowns' directory and the archive will be named like 2023_W49_ORIGINAL.zip
-   Contact files archives will be created in OUTPUT_FOLDER/Contacts/ directory and files will be named also like 2023_W49_ORIGINAL.zip

2.  minify original files
-   original files will be minified such that empty fields (fields that does not contain any valu) will be filter out. After minification the files are put in archive named like 2023_W49_MINIFIED.zip

## Example usage

**normal mode**
```bash
openmedia-archive archivate -i SOURCE_FOLDER -o OUTPUT_FOLDER
```

**dry run mode**
-   output files will be created in temporary directory
```bash
openmedia-archive --dry_run archivate -i SOURCE_FOLDER
```

**additional flags**
```bash
openmedia-archive -h
openmedia-archive archivate -h
```

## Error handling

**halts when unprocessable file encountered**
```bash
openmedia-archive archivate -i SOURCE_FOLDER -o OUTPUT_FOLDER
```

**continue processing folder when unprocessable file encountered**
```bash
openmedia-archive archivate -ifr -i SOURCE_FOLDER -o OUTPUT_FOLDER
openmedia-archive -n archivate -ifr -i SOURCE_FOLDER -o OUTPUT_FOLDER
```

-   useful for example in dry-mode

## Logging
### logging level
**info level (default)**
```shell
openmedia-archive archivate -i SOURCE_FOLDER -o OUTPUT_FOLDER
openmedia-archive -v=0 archivate -i SOURCE_FOLDER -o OUTPUT_FOLDER
```

**debug level**
```shell
openmedia-archive -v=-4 archivate -i SOURCE_FOLDER -o OUTPUT_FOLDER
```

**error level (show only errors)**
```shell
openmedia-archive -v=2 archivate -i SOURCE_FOLDER -o OUTPUT_FOLDER
```

### txt logging output

```shell
openmedia_archive -n -lt json archivate -i test/testdata/rundowns_mix/
```

```shell
time=2024-01-08T20:05:22.542+01:00 level=INFO source=process.go:346 msg=2020_W10_ORIGINAL.zip ArhiveRatio=0.067 MinifyRatio=1.000 original=13552180 compressed=905216 minified=13552180 file=test/testdata/rundowns_mix/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml
time=2024-01-08T20:05:23.457+01:00 level=INFO source=process.go:346 msg=2020_W10_MINIFIED.zip ArhiveRatio=0.024 MinifyRatio=0.327 original=13552180 compressed=319488 minified=4430320 file=test/testdata/rundowns_mix/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml
time=2024-01-08T20:05:23.662+01:00 level=INFO source=process.go:346 msg=2020_W10_ORIGINAL.zip ArhiveRatio=0.066 MinifyRatio=1.000 original=18364782 compressed=1204224 minified=18364782 file=test/testdata/rundowns_mix/RD_12-19_ČRo_Ostrava_-_Středa_04_03_2020_2_1603282_20200304234540.xml
time=2024-01-08T20:05:24.500+01:00 level=INFO source=process.go:346 msg=2020_W10_MINIFIED.zip ArhiveRatio=0.023 MinifyRatio=0.314 original=18364782 compressed=421888 minified=5772375 file=test/testdata/rundowns_mix/RD_12-19_ČRo_Ostrava_-_Středa_04_03_2020_2_1603282_20200304234540.xml
time=2024-01-08T20:05:24.500+01:00 level=INFO source=process.go:346 msg=GLOBAL_ORIGINAL ArhiveRatio=0.063 MinifyRatio=1.000 original=449249600 compressed=28196864 minified=449249600 file=test/testdata/rundowns_mix/
time=2024-01-08T20:05:24.500+01:00 level=INFO source=process.go:346 msg=GLOBAL_MINIFY ArhiveRatio=0.021 MinifyRatio=0.287 original=449249600 compressed=9629696 minified=129017125 file=test/testdata/rundowns_mix/
```

### json logging output

```shell
openmedia_archive -n -lt json archivate -i test/testdata/rundowns_mix/
```

```shell
{"time":"2024-01-08T20:08:41.365154585+01:00","level":"INFO","source":{"function":"github/czech-radio/openmedia-archive/internal.(*Process).WorkerLogInfo","file":"process.go","line":346},"msg":"2020_W10_MINIFIED.zip","ArhiveRatio":"0.024","MinifyRatio":"0.327","original":13552180,"compressed":319488,"minified":4430320,"file":"test/testdata/rundowns_mix/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml"}
{"time":"2024-01-08T20:08:41.519854699+01:00","level":"INFO","source":{"function":"github/czech-radio/openmedia-archive/internal.(*Process).WorkerLogInfo","file":"process.go","line":346},"msg":"2020_W10_ORIGINAL.zip","ArhiveRatio":"0.066","MinifyRatio":"1.000","original":18364782,"compressed":1204224,"minified":18364782,"file":"test/testdata/rundowns_mix/RD_12-19_ČRo_Ostrava_-_Středa_04_03_2020_2_1603282_20200304234540.xml"}
{"time":"2024-01-08T20:08:42.244667764+01:00","level":"INFO","source":{"function":"github/czech-radio/openmedia-archive/internal.(*Process).WorkerLogInfo","file":"process.go","line":346},"msg":"2020_W10_MINIFIED.zip","ArhiveRatio":"0.023","MinifyRatio":"0.314","original":18364782,"compressed":421888,"minified":5772375,"file":"test/testdata/rundowns_mix/RD_12-19_ČRo_Ostrava_-_Středa_04_03_2020_2_1603282_20200304234540.xml"}
{"time":"2024-01-08T20:08:42.244735372+01:00","level":"INFO","source":{"function":"github/czech-radio/openmedia-archive/internal.(*Process).WorkerLogInfo","file":"process.go","line":346},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.063","MinifyRatio":"1.000","original":449249600,"compressed":28196864,"minified":449249600,"file":"test/testdata/rundowns_mix/"}
{"time":"2024-01-08T20:08:42.244753922+01:00","level":"INFO","source":{"function":"github/czech-radio/openmedia-archive/internal.(*Process).WorkerLogInfo","file":"process.go","line":346},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.021","MinifyRatio":"0.287","original":449249600,"compressed":9629696,"minified":129017125,"file":"test/testdata/rundowns_mix/"}
```

## Developement guide and discussion
### Rundown files structure
#### Object/subobject structure

```plain
OM_OBJECT "Radio Rundown"
 OM_OBJECT "Hourly Rundown"
  OM_OBJECT "Sub Rundown"
   OM_OBJECT "Radio Story"
    OM_OBJECT "Contact item" [Optional]
    OM_OBJECT "Audio clip"   [Optional]
```

#### Object/field structure

```plain
OPENMEDIA
 OM_SERVER
 OM_OBJECT "Radio Rundown"
  OM_HEADER
   OM_FIELD [1-548]
  OM_RECORD
   OM_FILED [1-5012,-11,-12]
   OM_OBJECT "Hourly Rundown"

OM_OBJECT "Hourly Rundown"
 OM_HEADER
  OM_FIELD [1-548]
 OM_UPLINK
 OM_RECORD
  OM_FILED [1-5012,-11,-12]
  OM_OBJECT "Sub Rundown"
```

### Testing

#### Additional testing files

```
:/GŘ/Strategický rozvoj/Analytická sekce/_Archiv/Projekty/OpenMedia/04_03_2020
```

#### XML rundown validation

```
xmllint --schema schema.xsd
```
The XML does not validate when schema.xsd imports another XSD for common objects.
