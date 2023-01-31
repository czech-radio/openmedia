# openmedia-minify

Remove unnecessary or empty field from openmedia files.

## How it works?

Openmedia minify operates on Rundown files. It strips down unnecessary or empty fields and produces light version of an original file.
There are two flags needed to run the program `-i` for input folder and `-o` for output folder. Whole command would look like this:

```bash
openmedia-minifiy -i /path/to/source/folder -o /path/to/destination
```

When program runs it creates two files for each `RD_*.xml` file in output folder. Two files are:
- 1. minified version of an original XML in `UTF-8` format named: `RD_00-05_Plus_Friday_W02_2023_01_13.xml`
- 2. zipped original in `UTF-16` named same as 1. (inside of zip archive is an original file with original name): `RD_00-05_Plus_Friday_W02_2023_01_13.zip`

The date and name-day is derived from XML timetag and it is the same as given folder on `ANNOVA` disk. Input folder remains unchanged.

## Dependencies

Program requres libxml2-dev package to compile. Debian install:

```
apt-get install libxml2-dev
```

## Errors

There is a validation process of both input and output files. It can occasionally produces error. The file will mark CORRUPT in filname (TODO).
