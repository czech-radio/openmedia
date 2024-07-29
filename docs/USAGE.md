


## Command: root
### 1. root: help
Command input

running from source:

```
go run main.go -h
```

running compiled:

```
./openmedia -h
```

#### Command output

```
-h, -help
	display this help and exit

-U, -Usage=false
	Print usage manual

-H, -generalHelp=false
	Get help on all subcommands

-V, -version=false
	Print version of program.

-v, -verbose=0
	Level of verbosity. Lower the number the more verbose is the output.
	[6 4 2 0 -2 -4]

-logt, -logType=json
	Type of logs formating.
	[json plain]

-dr, -dryRun=false
	Dry run, useful for tests. Avoid any pernament changes to filesystem or any expensive tasks

-dc, -debugConfig=false
	Debug/print flag values

-logts, -logTest=false
	Print test logs


```


### 2. root: print version
Command input

running from source:

```
go run main.go -V
```

running compiled:

```
./openmedia -V
```

#### Command output

```
{ProgramName: Version:1.0.1 GitTag: GitCommit: BuildTime:}
```


### 3. root: print config
Command input

running from source:

```
go run main.go -dc
```

running compiled:

```
./openmedia -dc
```

#### Command output

```
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
```


### 4. root: print version and config
Command input

running from source:

```
go run main.go -V -dc
```

running compiled:

```
./openmedia -V -dc
```

#### Command output

```
{ProgramName: Version:1.0.1 GitTag: GitCommit: BuildTime:}
Root config: &{Usage:false GeneralHelp:false Version:true Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
```


### 5. root: test log [err]
Command input

running from source:

```
go run main.go -logt=plain -logts -v=6
```

running compiled:

```
./openmedia -logt=plain -logts -v=6
```

#### Command output

```
time=2024-07-29T23:21:26.201+02:00 level=ERROR source=sloger.go:53 msg="test error"
```


### 6. root: test log [err,warn]
Command input

running from source:

```
go run main.go -logt=plain -logts -v=4
```

running compiled:

```
./openmedia -logt=plain -logts -v=4
```

#### Command output

```
time=2024-07-29T23:21:26.587+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-29T23:21:26.587+02:00 level=ERROR source=sloger.go:53 msg="test error"
```


### 7. root: test log [err,warn,info]
Command input

running from source:

```
go run main.go -logt=plain -logts -v=0
```

running compiled:

```
./openmedia -logt=plain -logts -v=0
```

#### Command output

```
time=2024-07-29T23:21:27.005+02:00 level=INFO source=sloger.go:51 msg="test info"
time=2024-07-29T23:21:27.005+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-29T23:21:27.005+02:00 level=ERROR source=sloger.go:53 msg="test error"
```


### 8. root: terst log [err,warn,info,debug]
Command input

running from source:

```
go run main.go -logt=plain -logts -v=-4
```

running compiled:

```
./openmedia -logt=plain -logts -v=-4
```

#### Command output

```
time=2024-07-29T23:21:27.397+02:00 level=DEBUG source=config_mapany.go:253 msg="subcommand added" subname=archive
time=2024-07-29T23:21:27.397+02:00 level=DEBUG source=config_mapany.go:253 msg="subcommand added" subname=extractArchive
time=2024-07-29T23:21:27.397+02:00 level=DEBUG source=config_mapany.go:253 msg="subcommand added" subname=extractFile
time=2024-07-29T23:21:27.397+02:00 level=DEBUG source=sloger.go:50 msg="test debug"
time=2024-07-29T23:21:27.397+02:00 level=INFO source=sloger.go:51 msg="test info"
time=2024-07-29T23:21:27.397+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-29T23:21:27.397+02:00 level=ERROR source=sloger.go:53 msg="test error"
```


### 9. root: test log json
Command input

running from source:

```
go run main.go -logt=json -logts -v=-4
```

running compiled:

```
./openmedia -logt=json -logts -v=-4
```

#### Command output

```
{"time":"2024-07-29T23:21:27.775764533+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":253},"msg":"subcommand added","subname":"archive"}
{"time":"2024-07-29T23:21:27.775858855+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":253},"msg":"subcommand added","subname":"extractArchive"}
{"time":"2024-07-29T23:21:27.775870837+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":253},"msg":"subcommand added","subname":"extractFile"}
{"time":"2024-07-29T23:21:27.77587999+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":50},"msg":"test debug"}
{"time":"2024-07-29T23:21:27.775888784+02:00","level":"INFO","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":51},"msg":"test info"}
{"time":"2024-07-29T23:21:27.775896257+02:00","level":"WARN","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":52},"msg":"test warn"}
{"time":"2024-07-29T23:21:27.775903969+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":53},"msg":"test error"}
```

### Run summary

--- PASS: TestRunCommand_Root (3.55s)
    --- PASS: TestRunCommand_Root/1 (0.36s)
    --- PASS: TestRunCommand_Root/2 (0.39s)
    --- PASS: TestRunCommand_Root/3 (0.41s)
    --- PASS: TestRunCommand_Root/4 (0.41s)
    --- PASS: TestRunCommand_Root/5 (0.37s)
    --- PASS: TestRunCommand_Root/6 (0.39s)
    --- PASS: TestRunCommand_Root/7 (0.42s)
    --- PASS: TestRunCommand_Root/8 (0.39s)
    --- PASS: TestRunCommand_Root/9 (0.38s)
PASS
ok  	github/czech-radio/openmedia/cmd	3.573s



## Command: archive
### 1. archive: help
Command input

running from source:

```
go run main.go archive -h -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/1
```

running compiled:

```
./openmedia archive -h -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/1
```

#### Command output

```
-h, -help
	display this help and exit

-sdir, -SourceDirectory=.
	Source directory of rundown files.

-odir, -OutputDirectory=.
	Destination directory for archived rundwon files

-ct, -CompressionType=zip
	Type of file compression
	[zip]

-ifnc, -InvalidFilenameContinue=true
	Continue even though unknown filename encountered

-ifc, -InvalidFileContinue=true
	Continue even though unprocessable file encountered

-ifr, -InvalidFileRename=false
	Rename invalid files in source folder.

-pfr, -ProcessedFileRename=false
	Rename original rundown files after they are processed/archived: add "proccesed" prefix to source filename

-pfd, -ProcessedFileDelete=false
	Delete original rundown files after they are processed/archived.

-pfia, -PreserveFoldersInArchive=false
	Preserve source folder structure in archive

-R, -RecurseSourceDirectory=false
	Recurse source directory


```


### 2. archive: debug config
Command input

running from source:

```
go run main.go -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/2
```

running compiled:

```
./openmedia -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/2
```

#### Command output

```
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
Archive config: {SourceDirectory:/tmp/openmedia2824439708/cmd/SRC/cmd/archive OutputDirectory:/tmp/openmedia2824439708/cmd/DST/cmd/archive/2 CompressionType:zip InvalidFilenameContinue:true InvalidFileContinue:false InvalidFileRename:false ProcessedFileRename:false ProcessedFileDelete:false PreserveFoldersInArchive:false RecurseSourceDirectory:true}
```


### 3. archive: dry run
Command input

running from source:

```
go run main.go -dr -v=0 archive -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/3
```

running compiled:

```
./openmedia -dr -v=0 archive -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/3
```

#### Command output

```
{"time":"2024-07-29T23:21:29.407708478+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":62},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2824439708/cmd/DST/cmd/archive/3","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-07-29T23:21:29.40783634+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":65},"msg":"dry run activated","output_path":"/tmp/openmedia_archive2209640947"}
{"time":"2024-07-29T23:21:29.407923042+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-07-29T23:21:29.411689586+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-07-29T23:21:29.412246555+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-29T23:21:29.415364507+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-29T23:21:29.416163228+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-29T23:21:29.416288782+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-07-29T23:21:29.42212233+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-29T23:21:29.424086418+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":296},"msg":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-29T23:21:29.987854426+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-29T23:21:30.113355154+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:30.622917214+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia2824439708/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia2824439708/cmd/SRC/cmd/archive/hello_invalid2.txt"]}
{"time":"2024-07-29T23:21:30.623011182+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}

result: archive_create
{"time":"2024-07-29T23:21:30.623032015+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-29T23:21:30.623049227+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.072","MinifyRatio":"1.000","original":13646328,"compressed":987136,"minified":13646328,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia_archive2209640947"}
{"time":"2024-07-29T23:21:30.623069055+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.028","MinifyRatio":"0.325","original":13646328,"compressed":385024,"minified":4441033,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia_archive2209640947"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```


### 4. archive: exit on src file error or filename error
Command input

running from source:

```
go run main.go -v=0 archive -ifc -ifnc -R -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/4
```

running compiled:

```
./openmedia -v=0 archive -ifc -ifnc -R -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/4
```

#### Command output

```
{"time":"2024-07-29T23:21:31.001826434+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":62},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2824439708/cmd/DST/cmd/archive/4","CompressionType":"zip","InvalidFilenameContinue":false,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-29T23:21:31.002041261+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-29T23:21:31.002077116+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/helper.ErrorsCodeMap.ExitWithCode","file":"errors.go","line":112},"msg":"filenames_validation: {\"file does not have xml extension\":[\"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/hello_invalid.txt\",\"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/hello_invalid2.txt\"],\"filename is not valid OpenMedia file\":[\"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/ohter_files/hello_invalid.xml\",\"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/ohter_files/hello_invalid2.xml\"]}"}
exit status 12
```


### 5. archive: run exit on src file error
Command input

running from source:

```
go run main.go -v=0 archive -ifc -R -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/5
```

running compiled:

```
./openmedia -v=0 archive -ifc -R -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/5
```

#### Command output

```
{"time":"2024-07-29T23:21:31.367421255+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":62},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2824439708/cmd/DST/cmd/archive/5","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-29T23:21:31.367661832+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-29T23:21:31.370984991+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/1/1/0"}
{"time":"2024-07-29T23:21:31.372489686+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-29T23:21:31.375767583+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-29T23:21:31.376073666+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-29T23:21:31.376339595+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/2/2/0"}
{"time":"2024-07-29T23:21:31.382756121+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-29T23:21:31.384565904+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/helper.ErrorsCodeMap.ExitWithCode","file":"errors.go","line":112},"msg":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
exit status 12
```


### 6. archive: do not exit on any file error
Command input

running from source:

```
go run main.go -v=0 archive -R -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/6
```

running compiled:

```
./openmedia -v=0 archive -R -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/6
```

#### Command output

```
{"time":"2024-07-29T23:21:31.77909693+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":62},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2824439708/cmd/DST/cmd/archive/6","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-29T23:21:31.779345984+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-29T23:21:31.788696494+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-29T23:21:31.789714885+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/1/1/0"}
{"time":"2024-07-29T23:21:31.79494884+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-29T23:21:31.795112228+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/2/2/0"}
{"time":"2024-07-29T23:21:31.796977679+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-29T23:21:31.807200648+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-29T23:21:31.809420257+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":296},"msg":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-29T23:21:32.394458047+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/4/3/1"}
{"time":"2024-07-29T23:21:32.397128447+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/5/4/1"}
{"time":"2024-07-29T23:21:32.398281806+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.087","MinifyRatio":"1.000","original":47218,"compressed":4096,"minified":47218,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/ohter_files/CT_Krpаlek__Lukач_2_977869_20231212143738.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Sakísuw,_Frmšú_Tuesday_W50_2023_12_12T143738.xml"}
{"time":"2024-07-29T23:21:32.403148072+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.127","original":47218,"compressed":0,"minified":6007,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/ohter_files/CT_Krpаlek__Lukач_2_977869_20231212143738.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Sakísuw,_Frmšú_Tuesday_W50_2023_12_12T143738.xml"}
{"time":"2024-07-29T23:21:32.54596973+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:33.156259765+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.077","MinifyRatio":"1.000","original":11960034,"compressed":921600,"minified":11960034,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/ohter_files/RD_05-09_ČRo_Sever_-_Wed__04_03_2020_2_1607196_20200304234759.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:33.300521482+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:33.300676241+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/6/5/1"}
{"time":"2024-07-29T23:21:33.628303041+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.033","MinifyRatio":"0.323","original":11960034,"compressed":393216,"minified":3865086,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/ohter_files/RD_05-09_ČRo_Sever_-_Wed__04_03_2020_2_1607196_20200304234759.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia2824439708/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia2824439708/cmd/SRC/cmd/archive/hello_invalid2.txt"],"filename is not valid OpenMedia file":["/tmp/openmedia2824439708/cmd/SRC/cmd/archive/ohter_files/hello_invalid.xml","/tmp/openmedia2824439708/cmd/SRC/cmd/archive/ohter_files/hello_invalid2.xml"]}
{"time":"2024-07-29T23:21:33.628392831+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}

result: archive_create
{"time":"2024-07-29T23:21:33.628413801+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/6/5/1"}
{"time":"2024-07-29T23:21:33.628430979+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.075","MinifyRatio":"1.000","original":25653580,"compressed":1912832,"minified":25653580,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia2824439708/cmd/DST/cmd/archive/6"}
{"time":"2024-07-29T23:21:33.628450868+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.030","MinifyRatio":"0.324","original":25653580,"compressed":778240,"minified":8312126,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia2824439708/cmd/DST/cmd/archive/6"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```


### 7. archive: do not recurse the source folder
Command input

running from source:

```
go run main.go -v=0 archive -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/7
```

running compiled:

```
./openmedia -v=0 archive -sdir=/tmp/openmedia2824439708/cmd/SRC/cmd/archive -odir=/tmp/openmedia2824439708/cmd/DST/cmd/archive/7
```

#### Command output

```
{"time":"2024-07-29T23:21:34.036968959+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":62},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2824439708/cmd/DST/cmd/archive/7","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-07-29T23:21:34.037151529+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-07-29T23:21:34.040480905+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-07-29T23:21:34.041216393+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-29T23:21:34.045076844+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-29T23:21:34.045393564+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-07-29T23:21:34.045953679+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-29T23:21:34.051997202+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-29T23:21:34.053845135+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":296},"msg":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-29T23:21:34.654830961+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-29T23:21:34.889415929+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:35.568544946+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia2824439708/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia2824439708/cmd/SRC/cmd/archive/hello_invalid2.txt"]}
{"time":"2024-07-29T23:21:35.568619847+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}

result: archive_create
{"time":"2024-07-29T23:21:35.568637792+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-29T23:21:35.568650031+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.072","MinifyRatio":"1.000","original":13646328,"compressed":987136,"minified":13646328,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia2824439708/cmd/DST/cmd/archive/7"}
{"time":"2024-07-29T23:21:35.568663829+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.028","MinifyRatio":"0.325","original":13646328,"compressed":385024,"minified":4441033,"filePathSource":"/tmp/openmedia2824439708/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia2824439708/cmd/DST/cmd/archive/7"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia2824439708/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```

### Run summary

--- PASS: TestRunCommand_Archive (7.32s)
    --- PASS: TestRunCommand_Archive/1 (0.39s)
    --- PASS: TestRunCommand_Archive/2 (0.38s)
    --- PASS: TestRunCommand_Archive/3 (1.59s)
    --- PASS: TestRunCommand_Archive/4 (0.38s)
    --- PASS: TestRunCommand_Archive/5 (0.38s)
    --- PASS: TestRunCommand_Archive/6 (2.25s)
    --- PASS: TestRunCommand_Archive/7 (1.94s)
PASS
ok  	github/czech-radio/openmedia/cmd	7.334s



## Command: extractArchive
### 1. extractArchive: help
Command input

running from source:

```
go run main.go extractArchive -h -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/2
```

running compiled:

```
./openmedia extractArchive -h -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/2
```

#### Command output

```
-h, -help
	display this help and exit

-odir, -OutputDirectory=
	Output file path for extracted data.

-ofname, -OutputFileName=
	Output file path for extracted data.

-csvD, -CSVdelim=	
	csv column field delimiter
	[	 ;]

-excode, -ExtractorsCode=production_all
	Name of extractor which specifies the parts of xml to be extracted

-frns, -FilterRadioNames=
	Filter data corresponding to radio names

-fdf, -FilterDateFrom=2024-07-22 00:00:00 +0200 CEST
	Filter rundowns from date. Format of the date is given in form 'YYYY-mm-ddTHH:mm:ss' e.g. 2024, 2024-02-01 or 2024-02-01T10. The precission of date given is arbitrary.

-fdt, -FilterDateTo=2024-07-29 00:00:00 +0200 CEST
	Filter rundowns to date

-fisow, -FilterIsoWeeks=
	Filter data corresponding to specified ISO weeks

-fmonths, -FilterMonths=
	Filter data corresponding to specified months

-fwdays, -FilterWeekDays=
	Filter data corresponding to specified weekdays

-arn, -AddRecordNumbers=false
	Add record numbers columns and dependent columns

-valfn, -ValidatorFileName=
	xlsx file containing validation receipe

-frfn, -FilterFileName=
	Special filters filename. The filter filename specifies how the file is parsed and how it is used

-frsn, -FilterSheetName=data
	Special filters sheetname

-sdir, -SourceDirectory=
	Source rundown folder archive.

-sdirType, -SourceDirectoryType=MINIFIED.zip
	type of source directory where the rundowns resides


```


### 2. extractArchive: extract all story parts from minified rundowns
Command input

running from source:

```
go run main.go extractArchive -ofname=production -excode=production_all -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/3
```

running compiled:

```
./openmedia extractArchive -ofname=production -excode=production_all -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/3
```

#### Command output

```
{ExtractorsCode:production_all FilterRadioNames:map[] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName: AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/3 OutputFileName:production CSVdelim:	}
{FilterFileName: FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-29T23:21:37.334048068+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:37.334166316+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-29T23:21:37.334182886+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:37.334192312+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":2}
{"time":"2024-07-29T23:21:37.334300148+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-29T23:21:37.334320643+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:37.663627107+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":166}
{"time":"2024-07-29T23:21:37.663667345+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:38.026999467+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":147}
{"time":"2024-07-29T23:21:38.02724743+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-29T23:21:38.027274705+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-29T23:21:38.029853004+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-29T23:21:38.029901941+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-29T23:21:38.030148478+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":312}
{"time":"2024-07-29T23:21:38.030180768+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":312}
{"time":"2024-07-29T23:21:38.042015423+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/3/production_production_all_base_wh.csv","bytesCount":1100}
{"time":"2024-07-29T23:21:38.042131728+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/3/production_production_all_base_wh.csv","bytesCount":154684}
{"time":"2024-07-29T23:21:38.042182198+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/3/production_production_all_base_woh.csv","bytesCount":468}
{"time":"2024-07-29T23:21:38.042261744+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/3/production_production_all_base_woh.csv","bytesCount":154684}
{"time":"2024-07-29T23:21:38.042278184+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDataset","file":"outputs.go","line":51},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-29T23:21:38.042297791+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":75},"msg":"filter file not specified"}
```


### 3. extractArchive: extract all contacts from minified rundowns
Command input

running from source:

```
go run main.go extractArchive -ofname=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/4
```

running compiled:

```
./openmedia extractArchive -ofname=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/4
```

#### Command output

```
{ExtractorsCode:production_contacts FilterRadioNames:map[] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName: AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/4 OutputFileName:production CSVdelim:	}
{FilterFileName: FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-29T23:21:38.885956247+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:38.886120735+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-29T23:21:38.886281899+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:38.886333269+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":2}
{"time":"2024-07-29T23:21:38.886587452+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-29T23:21:38.886679138+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:39.701503385+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":166}
{"time":"2024-07-29T23:21:39.701554618+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:40.24308042+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":147}
{"time":"2024-07-29T23:21:40.243243505+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-29T23:21:40.243260247+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-29T23:21:40.244619665+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-29T23:21:40.24463641+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-29T23:21:40.244727971+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-29T23:21:40.244740434+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-29T23:21:40.244822602+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-29T23:21:40.244835437+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-29T23:21:40.245459375+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/4/production_production_contacts_base_wh.csv","bytesCount":1026}
{"time":"2024-07-29T23:21:40.245496354+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/4/production_production_contacts_base_wh.csv","bytesCount":29762}
{"time":"2024-07-29T23:21:40.245528789+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/4/production_production_contacts_base_woh.csv","bytesCount":434}
{"time":"2024-07-29T23:21:40.245566827+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/4/production_production_contacts_base_woh.csv","bytesCount":29762}
{"time":"2024-07-29T23:21:40.245579006+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDataset","file":"outputs.go","line":51},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-29T23:21:40.245588198+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":75},"msg":"filter file not specified"}
```


### 4. extractArchive: extract all story parts from minified rundowns, extract only specified radios
Command input

running from source:

```
go run main.go extractArchive -ofname=production -excode=production_all -frns=Olomouc,Plus -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/5
```

running compiled:

```
./openmedia extractArchive -ofname=production -excode=production_all -frns=Olomouc,Plus -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/5
```

#### Command output

```
{ExtractorsCode:production_all FilterRadioNames:map[Olomouc:true Plus:true] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName: AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/5 OutputFileName:production CSVdelim:	}
{FilterFileName: FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-29T23:21:40.614008853+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:40.614103124+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":0,"matched":true}
{"time":"2024-07-29T23:21:40.614117919+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":0}
{"time":"2024-07-29T23:21:40.614127264+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":0}
{"time":"2024-07-29T23:21:40.614249368+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-29T23:21:40.614315005+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-29T23:21:40.614326552+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-29T23:21:40.614336387+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-29T23:21:40.614345018+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-29T23:21:40.614352847+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-29T23:21:40.61436101+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-29T23:21:40.614444731+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/5/production_production_all_base_wh.csv","bytesCount":1100}
{"time":"2024-07-29T23:21:40.614476843+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/5/production_production_all_base_wh.csv","bytesCount":0}
{"time":"2024-07-29T23:21:40.614507675+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/5/production_production_all_base_woh.csv","bytesCount":468}
{"time":"2024-07-29T23:21:40.614532245+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/5/production_production_all_base_woh.csv","bytesCount":0}
{"time":"2024-07-29T23:21:40.614546926+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDataset","file":"outputs.go","line":51},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-29T23:21:40.614557169+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":75},"msg":"filter file not specified"}
```


### 5. extractArchive: extract all story parts from minified rundowns and validate
Command input

running from source:

```
go run main.go extractArchive -ofname=production -excode=production_all -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6
```

running compiled:

```
./openmedia extractArchive -ofname=production -excode=production_all -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6
```

#### Command output

```
{ExtractorsCode:production_all FilterRadioNames:map[] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName:../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6 OutputFileName:production CSVdelim:	}
{FilterFileName: FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-29T23:21:41.033846547+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:41.033938905+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-29T23:21:41.033955004+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:41.03396991+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":2}
{"time":"2024-07-29T23:21:41.034081086+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-29T23:21:41.034104277+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:41.414611256+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":166}
{"time":"2024-07-29T23:21:41.414660807+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:41.789072809+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":147}
{"time":"2024-07-29T23:21:41.789289957+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-29T23:21:41.789308663+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-29T23:21:41.791897661+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-29T23:21:41.791951734+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-29T23:21:41.792200697+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":312}
{"time":"2024-07-29T23:21:41.792235705+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":312}
{"time":"2024-07-29T23:21:41.803499537+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6/production_production_all_base_wh.csv","bytesCount":1100}
{"time":"2024-07-29T23:21:41.803700921+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6/production_production_all_base_wh.csv","bytesCount":154684}
{"time":"2024-07-29T23:21:41.803803119+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6/production_production_all_base_woh.csv","bytesCount":468}
{"time":"2024-07-29T23:21:41.804246531+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6/production_production_all_base_woh.csv","bytesCount":154684}
{"time":"2024-07-29T23:21:41.896122919+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6/production_production_all_base_validated_wh.csv","bytesCount":1100}
{"time":"2024-07-29T23:21:41.896227537+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6/production_production_all_base_validated_wh.csv","bytesCount":146747}
{"time":"2024-07-29T23:21:41.896272705+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6/production_production_all_base_validated_woh.csv","bytesCount":468}
{"time":"2024-07-29T23:21:41.896342353+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6/production_production_all_base_validated_woh.csv","bytesCount":146747}
{"time":"2024-07-29T23:21:41.89650666+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":331},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/6/production_base_validated_log.csv","bytesCount":4948}
{"time":"2024-07-29T23:21:41.896522374+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":75},"msg":"filter file not specified"}
```


### 6. extractArchive: extract all story parts from minified rundowns, validate and use filter oposition
Command input

running from source:

```
go run main.go extractArchive -ofname=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7
```

running compiled:

```
./openmedia extractArchive -ofname=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7
```

#### Command output

```
{ExtractorsCode:production_contacts FilterRadioNames:map[] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName:../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7 OutputFileName:production CSVdelim:	}
{FilterFileName:../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-29T23:21:42.281239486+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:42.281337071+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-29T23:21:42.281352315+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:42.281362091+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":2}
{"time":"2024-07-29T23:21:42.281464232+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-29T23:21:42.28148195+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:42.5774109+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":166}
{"time":"2024-07-29T23:21:42.577451726+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:42.869927859+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":147}
{"time":"2024-07-29T23:21:42.870060825+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-29T23:21:42.870074318+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-29T23:21:42.871395644+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-29T23:21:42.871413475+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-29T23:21:42.871491418+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-29T23:21:42.871500005+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-29T23:21:42.871555918+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-29T23:21:42.871564556+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-29T23:21:42.872472469+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_wh.csv","bytesCount":1026}
{"time":"2024-07-29T23:21:42.872541196+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_wh.csv","bytesCount":29762}
{"time":"2024-07-29T23:21:42.872584227+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_woh.csv","bytesCount":434}
{"time":"2024-07-29T23:21:42.872629918+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_woh.csv","bytesCount":29762}
{"time":"2024-07-29T23:21:42.938414493+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_wh.csv","bytesCount":1026}
{"time":"2024-07-29T23:21:42.938469044+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_wh.csv","bytesCount":24855}
{"time":"2024-07-29T23:21:42.938499774+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_woh.csv","bytesCount":434}
{"time":"2024-07-29T23:21:42.938526672+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_woh.csv","bytesCount":24855}
{"time":"2024-07-29T23:21:42.938663864+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":331},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_base_validated_log.csv","bytesCount":4256}
{"time":"2024-07-29T23:21:43.11143679+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":1187}
{"time":"2024-07-29T23:21:43.111495588+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":27275}
{"time":"2024-07-29T23:21:43.11153484+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":492}
{"time":"2024-07-29T23:21:43.11156902+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":27275}
```


### 7. extractArchive: extract all story parts from minified rundowns, validate and use filter eurovolby
Command input

running from source:

```
go run main.go extractArchive -ofname=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8
```

running compiled:

```
./openmedia extractArchive -ofname=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx -sdir=/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8
```

#### Command output

```
{ExtractorsCode:production_contacts FilterRadioNames:map[] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName:../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8 OutputFileName:production CSVdelim:	}
{FilterFileName:../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-29T23:21:43.519759525+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:43.519850716+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-29T23:21:43.519866273+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":1}
{"time":"2024-07-29T23:21:43.519876016+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":2}
{"time":"2024-07-29T23:21:43.519982031+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-29T23:21:43.519999909+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:43.907108289+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":166}
{"time":"2024-07-29T23:21:43.907224141+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia2241653920/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-29T23:21:44.264711776+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":147}
{"time":"2024-07-29T23:21:44.264919436+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-29T23:21:44.264948118+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-29T23:21:44.267744873+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-29T23:21:44.267782766+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-29T23:21:44.267951817+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-29T23:21:44.267974314+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-29T23:21:44.26809937+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-29T23:21:44.268118894+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-29T23:21:44.26912753+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_wh.csv","bytesCount":1026}
{"time":"2024-07-29T23:21:44.269179529+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_wh.csv","bytesCount":29762}
{"time":"2024-07-29T23:21:44.269219898+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_woh.csv","bytesCount":434}
{"time":"2024-07-29T23:21:44.269261255+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_woh.csv","bytesCount":29762}
{"time":"2024-07-29T23:21:44.337038024+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_wh.csv","bytesCount":1026}
{"time":"2024-07-29T23:21:44.33709101+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_wh.csv","bytesCount":24855}
{"time":"2024-07-29T23:21:44.337121791+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_woh.csv","bytesCount":434}
{"time":"2024-07-29T23:21:44.337147387+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_woh.csv","bytesCount":24855}
{"time":"2024-07-29T23:21:44.337272291+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":331},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_base_validated_log.csv","bytesCount":4256}
{"time":"2024-07-29T23:21:44.430638274+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":1154}
{"time":"2024-07-29T23:21:44.430693091+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":27011}
{"time":"2024-07-29T23:21:44.430728363+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":487}
{"time":"2024-07-29T23:21:44.430767918+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2241653920/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":27011}
```

### Run summary

--- PASS: TestRunCommand_ExtractArchive (8.32s)
    --- PASS: TestRunCommand_ExtractArchive/1 (0.67s)
    --- PASS: TestRunCommand_ExtractArchive/2 (1.28s)
    --- PASS: TestRunCommand_ExtractArchive/3 (2.19s)
    --- PASS: TestRunCommand_ExtractArchive/4 (0.37s)
    --- PASS: TestRunCommand_ExtractArchive/5 (1.28s)
    --- PASS: TestRunCommand_ExtractArchive/6 (1.21s)
    --- PASS: TestRunCommand_ExtractArchive/7 (1.32s)
PASS
ok  	github/czech-radio/openmedia/cmd	8.331s



## Command: extractFile
### 1. extractFile: help
Command input

running from source:

```
go run main.go extractFile -h
```

running compiled:

```
./openmedia extractFile -h
```

#### Command output

```
-h, -help
	display this help and exit

-odir, -OutputDirectory=
	Output file path for extracted data.

-ofname, -OutputFileName=
	Output file path for extracted data.

-csvD, -CSVdelim=	
	csv column field delimiter
	[	 ;]

-excode, -ExtractorsCode=production_all
	Name of extractor which specifies the parts of xml to be extracted

-frns, -FilterRadioNames=
	Filter data corresponding to radio names

-fdf, -FilterDateFrom=2024-07-22 00:00:00 +0200 CEST
	Filter rundowns from date. Format of the date is given in form 'YYYY-mm-ddTHH:mm:ss' e.g. 2024, 2024-02-01 or 2024-02-01T10. The precission of date given is arbitrary.

-fdt, -FilterDateTo=2024-07-29 00:00:00 +0200 CEST
	Filter rundowns to date

-fisow, -FilterIsoWeeks=
	Filter data corresponding to specified ISO weeks

-fmonths, -FilterMonths=
	Filter data corresponding to specified months

-fwdays, -FilterWeekDays=
	Filter data corresponding to specified weekdays

-arn, -AddRecordNumbers=false
	Add record numbers columns and dependent columns

-valfn, -ValidatorFileName=
	xlsx file containing validation receipe

-frfn, -FilterFileName=
	Special filters filename. The filter filename specifies how the file is parsed and how it is used

-frsn, -FilterSheetName=data
	Special filters sheetname

-sfp, -SourceFilePath=
	Source rundown file name.


```


### 2. extractFile: print config
Command input

running from source:

```
go run main.go extractFile -dc
```

running compiled:

```
./openmedia extractFile -dc
```

#### Command output

```
panic: flag: OutputDirectory value cannot be empty

goroutine 1 [running]:
github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).ParseFlag(0x6ae302?, {0xc000276510, 0xf}, {0x71f200?, 0xc00027e0b0?, 0x6cf8a0?}, 0x4)
	/home/jk/runs/golang/go2023/go_utils/pkg/configure/config_option_parse.go:55 +0x18b6
github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).ParseFlags(0xa6d220, {0x6c2100, 0xc00027e0b0})
	/home/jk/runs/golang/go2023/go_utils/pkg/configure/config_mapany.go:329 +0x350
github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).SubcommandOptionsParse(0xa6d220, {0x6c2100, 0xc00027e0b0})
	/home/jk/runs/golang/go2023/go_utils/pkg/configure/config_mapany.go:273 +0x15d
github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandExtractFile({{0xc00025c450, {0xa40300, 0x8, 0x8}, 0xc00025c510, {0x0, 0x0}, {0x7175c0, 0xc000106aa0}, 0xc00025c4b0, ...}})
	/home/jk/CRO/CRO_BASE/openmedia/cmd/extract_file.go:27 +0x78
github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).RunRoot(0xa6d1a0)
	/home/jk/runs/golang/go2023/go_utils/pkg/configure/config_mapany.go:249 +0x104
github/czech-radio/openmedia/cmd.RunRoot()
	/home/jk/CRO/CRO_BASE/openmedia/cmd/root.go:43 +0x265
main.main()
	/home/jk/CRO/CRO_BASE/openmedia/main.go:27 +0xf
exit status 2
```


### 3. extractFile: extract original UTF16 file
Command input

running from source:

```
go run main.go extractFile -sfp=/tmp/openmedia30922227/cmd/SRC/cmd/extractFile/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622_UTF16.xml -odir=/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4 -ofname=orig -frns=Plus,Sek -fisow=1,2,3 -fwdays=1,2,3 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx
```

running compiled:

```
./openmedia extractFile -sfp=/tmp/openmedia30922227/cmd/SRC/cmd/extractFile/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622_UTF16.xml -odir=/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4 -ofname=orig -frns=Plus,Sek -fisow=1,2,3 -fwdays=1,2,3 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx
```

#### Command output

```
{ExtractorsCode:production_all FilterRadioNames:map[Plus:true Sek:true] FilterDateFrom:2024-07-22 00:00:00 +0200 CEST FilterDateTo:2024-07-29 00:00:00 +0200 CEST DateRange:[0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC] FilterIsoWeeks:map[1:true 2:true 3:true] FilterMonths:map[] FilterWeekDays:map[1:true 2:true 3:true] ValidatorFileName:../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory: SourceDirectoryType: SourceFilePath:/tmp/openmedia30922227/cmd/SRC/cmd/extractFile/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622_UTF16.xml SourceCharEncoding: OutputDirectory:/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4 OutputFileName:orig CSVdelim:	}
{FilterFileName:../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-29T23:21:46.679529305+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":166,"filtered":166}
{"time":"2024-07-29T23:21:46.679586677+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":166}
{"time":"2024-07-29T23:21:46.680433141+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":166,"filtered":166}
{"time":"2024-07-29T23:21:46.680448901+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":166}
{"time":"2024-07-29T23:21:46.680502573+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":166,"filtered":165}
{"time":"2024-07-29T23:21:46.680519186+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":165}
{"time":"2024-07-29T23:21:46.682676984+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_wh.csv","bytesCount":1100}
{"time":"2024-07-29T23:21:46.682738532+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_wh.csv","bytesCount":82479}
{"time":"2024-07-29T23:21:46.682781802+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_woh.csv","bytesCount":468}
{"time":"2024-07-29T23:21:46.682834631+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_woh.csv","bytesCount":82479}
{"time":"2024-07-29T23:21:46.736479749+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_validated_wh.csv","bytesCount":1100}
{"time":"2024-07-29T23:21:46.736546427+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_validated_wh.csv","bytesCount":78374}
{"time":"2024-07-29T23:21:46.73657664+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_validated_woh.csv","bytesCount":468}
{"time":"2024-07-29T23:21:46.736619857+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_validated_woh.csv","bytesCount":78374}
{"time":"2024-07-29T23:21:46.736716736+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":331},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_base_validated_log.csv","bytesCount":2660}
{"time":"2024-07-29T23:21:46.829591908+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_validated_filtered_wh.csv","bytesCount":1228}
{"time":"2024-07-29T23:21:46.829670257+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_validated_filtered_wh.csv","bytesCount":86499}
{"time":"2024-07-29T23:21:46.829712363+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_validated_filtered_woh.csv","bytesCount":521}
{"time":"2024-07-29T23:21:46.829766076+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/4/orig_production_all_base_validated_filtered_woh.csv","bytesCount":86499}
```


### 4. extractFile: extract original UT8 file
Command input

running from source:

```
go run main.go extractFile -sfp=/tmp/openmedia30922227/cmd/SRC/cmd/extractFile/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622_UTF8.xml -odir=/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5 -ofname=orig -frns=Plus,Sek -fisow=1,2,3 -fwdays=1,2,3 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx -arn
```

running compiled:

```
./openmedia extractFile -sfp=/tmp/openmedia30922227/cmd/SRC/cmd/extractFile/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622_UTF8.xml -odir=/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5 -ofname=orig -frns=Plus,Sek -fisow=1,2,3 -fwdays=1,2,3 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx -arn
```

#### Command output

```
{ExtractorsCode:production_all FilterRadioNames:map[Plus:true Sek:true] FilterDateFrom:2024-07-22 00:00:00 +0200 CEST FilterDateTo:2024-07-29 00:00:00 +0200 CEST DateRange:[0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC] FilterIsoWeeks:map[1:true 2:true 3:true] FilterMonths:map[] FilterWeekDays:map[1:true 2:true 3:true] ValidatorFileName:../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx AddRecordNumbers:true ComputeUniqueRows:false}
{SourceDirectory: SourceDirectoryType: SourceFilePath:/tmp/openmedia30922227/cmd/SRC/cmd/extractFile/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622_UTF8.xml SourceCharEncoding: OutputDirectory:/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5 OutputFileName:orig CSVdelim:	}
{FilterFileName:../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
FIX
{"time":"2024-07-29T23:21:47.728360381+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":166,"filtered":166}
{"time":"2024-07-29T23:21:47.728413808+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":166}
{"time":"2024-07-29T23:21:47.729168907+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":166,"filtered":166}
{"time":"2024-07-29T23:21:47.729183728+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":166}
{"time":"2024-07-29T23:21:47.729236324+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":166,"filtered":165}
{"time":"2024-07-29T23:21:47.729248568+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":165}
{"time":"2024-07-29T23:21:47.731940846+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_wh.csv","bytesCount":1125}
{"time":"2024-07-29T23:21:47.732003459+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_wh.csv","bytesCount":84094}
{"time":"2024-07-29T23:21:47.732040045+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_woh.csv","bytesCount":478}
{"time":"2024-07-29T23:21:47.732102544+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_woh.csv","bytesCount":84094}
{"time":"2024-07-29T23:21:47.831416935+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_validated_wh.csv","bytesCount":1125}
{"time":"2024-07-29T23:21:47.83149879+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_validated_wh.csv","bytesCount":79989}
{"time":"2024-07-29T23:21:47.83155409+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_validated_woh.csv","bytesCount":478}
{"time":"2024-07-29T23:21:47.831617506+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_validated_woh.csv","bytesCount":79989}
{"time":"2024-07-29T23:21:47.831798241+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":331},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_base_validated_log.csv","bytesCount":2660}
{"time":"2024-07-29T23:21:47.918645323+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_validated_filtered_wh.csv","bytesCount":1253}
{"time":"2024-07-29T23:21:47.918718406+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_validated_filtered_wh.csv","bytesCount":88114}
{"time":"2024-07-29T23:21:47.91876087+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_validated_filtered_woh.csv","bytesCount":531}
{"time":"2024-07-29T23:21:47.91881374+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia30922227/cmd/DST/cmd/extractFile/5/orig_production_all_base_validated_filtered_woh.csv","bytesCount":88114}
```

--- PASS: TestRunCommand_ExtractFile (3.09s)
    --- PASS: TestRunCommand_ExtractFile/1 (0.38s)
    --- PASS: TestRunCommand_ExtractFile/2 (0.40s)
    --- PASS: TestRunCommand_ExtractFile/3 (1.21s)
    --- PASS: TestRunCommand_ExtractFile/4 (1.09s)
PASS
ok  	github/czech-radio/openmedia/cmd	3.099s
