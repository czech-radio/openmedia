# openmedia-minify


_Remove unnecessary or empty fields from openmedia files._

[![build](https://github.com/czech-radio/openmedia-minify/actions/workflows/main.yml/badge.svg)](https://github.com/czech-radio/openmedia-minify/actions/workflows/main.yml)


## How it works?

Openmedia minify operates on Rundown files. It strips down unnecessary or empty fields and produces light version of an original file.
There are two flags needed to run the program `-i` for input folder and `-o` for output folder. Whole command would look like this:

```bash
openmedia-minifiy -i /path/to/source/folder -o /path/to/destination
```

When program runs it creates two files in output folder for each `RD_*.xml` file in input folder. Two files are:
- minified (stripped down) XML version of original in `UTF-8` format. It is named: `RD_00-05_Plus_Friday_W02_2023_01_13.xml`
- zipped XML original in `UTF-16` named similarly (inside of zip archive is the original file with its original name): `RD_00-05_Plus_Friday_W02_2023_01_13.zip`

The date and name-day is derived from XML timetag and it is the same as given folder on `ANNOVA` disk. Input folder remains unchanged.

## Logging

When everything works well you should see something similar in console output:

```
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

```
apt-get install libxml2-dev pkg-config
go mod tidy
go build
```

## Errors

There is a validation process of both input and output files. It can occasionally produce an error. Is such case resulting file will be marked as `_MALFORMED` in filename.

Process can be quite memory hungry. It can make use of lot of ram in host computer, occassioanlly even crash it.
