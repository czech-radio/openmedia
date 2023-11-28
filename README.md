# openmedia-reduce

**The console application to reduce size of OpenMedia XML exports.**

[![build](https://github.com/czech-radio/openmedia-reduce/actions/workflows/main.yml/badge.svg)](https://github.com/czech-radio/openmedia-reduce/actions/workflows/main.yml)

## Description

Program operates on Rundown files. It strips down unnecessary or empty fields and produces light version of an original file.
There are two flags needed to run the program `-i` for input folder and `-o` for output folder. Whole command would look like this:

```bash
openmedia-reduce -i /path/to/source/folder -o /path/to/destination
```

When program runs it creates two files in output folder. Two files are:

- zipped all minified (stripped down) version of original XML in `UTF-8` format. It is named: `2023_W02_MINIFIED.zip`
- zipped all XML originals in `UTF-16` named similarly (inside of zip archive there are all of the original files with its original file-names): `2023_W02_ORIGINAL.zip`

The date and name-day is derived from XML timetag and it is the same as given folder on `ANNOVA` disk. Input folder remains unchanged.

## Logging

When everything works well you should see something similar in console output:

```shell
2023/01/31 15:07:35 Minifying: /mnt/cro.cz/Rundowns/2023/W04/RD_05-09_ČRo_Brno_-_Sun_29_01_2023_2_14561296_20230130001249.xml
2023/01/31 15:07:35 Document minified from 57691 lines to 26500 lines, ratio: 45.934372%
2023/01/31 15:07:35 Zipping: tmp/RD_05-09_ČRo_Brno_Sunday_W04_2023_01_29.zip
2023/01/31 15:07:35 Validating source file: /mnt/cro.cz/Rundowns/2023/W04/RD_05-09_ČRo_Brno_-_Sun_29_01_2023_2_14561296_20230130001249.xml
2023/01/31 15:07:35 Validating destination file: tmp/RD_05-09_ČRo_Brno_Sunday_W04_2023_01_29.xml
2023/01/31 15:07:35 Minifying PASSED! 70/498
2023/01/31 15:07:43 Minifying: /mnt/cro.cz/Rundowns/2023/W04/RD_05-09_ČRo_Brno_-_Thu_26_01_2023_2_14525751_20230127001315.xml
2023/01/31 15:07:43 Document minified from 63116 lines to 29251 lines, ratio: 46.344826%
2023/01/31 15:07:43 Zipping: tmp/RD_05-09_ČRo_Brno_Thursday_W04_2023_01_26.zip
2023/01/31 15:07:43 Validating source file: /mnt/cro.cz/Rundowns/2023/W04/RD_05-09_ČRo_Brno_-_Thu_26_01_2023_2_14525751_20230127001315.xml
2023/01/31 15:07:43 Validating destination file: tmp/RD_05-09_ČRo_Brno_Thursday_W04_2023_01_26.xml
2023/01/31 15:07:43 Minifying PASSED! 71/498
```

## Dependencies & Build

Program requires libxml2-dev package to compile. Debian install:

```shell
apt-get install libxml2-dev pkg-config
go mod tidy
go build
```

## Errors

There is a validation process of both input and output files. It can occasionally produce an error. Is such case resulting file will be marked as `_MALFORMED` in filename.

TODO: better memory handling, fixed by [b20445b](https://github.com/czech-radio/openmedia-reduce/commit/b20445b429d019a6392fb6738ea79c188a8878a7)

## Developement guide and discussion

### Which golang version to lock?
*2023-11-05*

-   arch_actual: go version go1.21.3 linux/amd64
-   alpine_3.18 (./deploy/dockerfile_devel.yml): go version go1.20.10 linux/amd64

### Logging
[](https://betterstack.com/community/guides/logging/best-golang-logging-libraries/)

-   slog: new bulit-in logging in Go 1.21 (chosen one)
-   zerolog: fastest
-   zap: fast yet flexible

### Commandline interface

-   using cobra package without viper

### Rundown files

#### Element object structure

```
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

#### Object structure

```
OM_OBJECT "Radio Rundown"
 OM_OBJECT "Hourly Rundown"
  OM_OBJECT "Sub Rundown"
   OM_OBJECT "Radio Story"
    OM_OBJECT "Contact item" [Optional]
    OM_OBJECT "Audio clip"   [Optional]
```

### Testing

#### Additional testing files

':/GŘ/Strategický rozvoj/Analytická sekce/_Archiv/Projekty/OpenMedia/04_03_2020'

#### XML rundown validation

xmllint --schema schema.xsd`
The xml does not validate when schema.xsd imports another xsd for common ojects.
