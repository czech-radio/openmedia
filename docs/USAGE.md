


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
time=2024-07-08T11:00:14.051+02:00 level=ERROR source=sloger.go:53 msg="test error"
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
time=2024-07-08T11:00:14.447+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-08T11:00:14.447+02:00 level=ERROR source=sloger.go:53 msg="test error"
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
time=2024-07-08T11:00:14.886+02:00 level=INFO source=sloger.go:51 msg="test info"
time=2024-07-08T11:00:14.886+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-08T11:00:14.886+02:00 level=ERROR source=sloger.go:53 msg="test error"
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
time=2024-07-08T11:00:15.320+02:00 level=DEBUG source=config_mapany.go:349 msg="subcommand added" subname=archive
time=2024-07-08T11:00:15.320+02:00 level=DEBUG source=config_mapany.go:349 msg="subcommand added" subname=extractArchive
time=2024-07-08T11:00:15.320+02:00 level=DEBUG source=sloger.go:50 msg="test debug"
time=2024-07-08T11:00:15.320+02:00 level=INFO source=sloger.go:51 msg="test info"
time=2024-07-08T11:00:15.320+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-08T11:00:15.320+02:00 level=ERROR source=sloger.go:53 msg="test error"
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
{"time":"2024-07-08T11:00:15.725165968+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":349},"msg":"subcommand added","subname":"archive"}
{"time":"2024-07-08T11:00:15.725244337+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":349},"msg":"subcommand added","subname":"extractArchive"}
{"time":"2024-07-08T11:00:15.725260988+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":50},"msg":"test debug"}
{"time":"2024-07-08T11:00:15.725276151+02:00","level":"INFO","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":51},"msg":"test info"}
{"time":"2024-07-08T11:00:15.725286663+02:00","level":"WARN","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":52},"msg":"test warn"}
{"time":"2024-07-08T11:00:15.725297294+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":53},"msg":"test error"}
```

### Run summary

--- PASS: TestRunCommand_Root (3.90s)
    --- PASS: TestRunCommand_Root/1 (0.44s)
    --- PASS: TestRunCommand_Root/2 (0.45s)
    --- PASS: TestRunCommand_Root/3 (0.46s)
    --- PASS: TestRunCommand_Root/4 (0.47s)
    --- PASS: TestRunCommand_Root/5 (0.40s)
    --- PASS: TestRunCommand_Root/6 (0.40s)
    --- PASS: TestRunCommand_Root/7 (0.44s)
    --- PASS: TestRunCommand_Root/8 (0.44s)
    --- PASS: TestRunCommand_Root/9 (0.40s)
PASS
ok  	github/czech-radio/openmedia/cmd	3.913s



## Command: archive
### 1. archive: help
Command input

running from source:

```
go run main.go archive -h -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/1
```

running compiled:

```
./openmedia archive -h -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/1
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
go run main.go -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/2
```

running compiled:

```
./openmedia -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/2
```

#### Command output

```
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
Archive config: {SourceDirectory:/tmp/openmedia2411149295/cmd/SRC/cmd/archive OutputDirectory:/tmp/openmedia2411149295/cmd/DST/cmd/archive/2 CompressionType:zip InvalidFilenameContinue:true InvalidFileContinue:false InvalidFileRename:false ProcessedFileRename:false ProcessedFileDelete:false PreserveFoldersInArchive:false RecurseSourceDirectory:true}
```


### 3. archive: dry run
Command input

running from source:

```
go run main.go -dr -v=0 archive -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/3
```

running compiled:

```
./openmedia -dr -v=0 archive -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/3
```

#### Command output

```
{"time":"2024-07-08T11:00:17.444916992+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2411149295/cmd/DST/cmd/archive/3","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-07-08T11:00:17.445067809+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":66},"msg":"dry run activated","output_path":"/tmp/openmedia_archive1272023116"}
{"time":"2024-07-08T11:00:17.445161769+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-07-08T11:00:17.448502316+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-07-08T11:00:17.449130505+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T11:00:17.452229048+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T11:00:17.452761353+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-07-08T11:00:17.452731991+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T11:00:17.459171198+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T11:00:17.461132702+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":315},"msg":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-08T11:00:17.998381722+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-08T11:00:18.15840167+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:18.671231473+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia2411149295/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia2411149295/cmd/SRC/cmd/archive/hello_invalid2.txt"]}
{"time":"2024-07-08T11:00:18.671309522+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}

result: archive_create
{"time":"2024-07-08T11:00:18.671323013+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-08T11:00:18.671332273+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.072","MinifyRatio":"1.000","original":13646328,"compressed":987136,"minified":13646328,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia_archive1272023116"}
{"time":"2024-07-08T11:00:18.671341675+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.028","MinifyRatio":"0.325","original":13646328,"compressed":385024,"minified":4441033,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia_archive1272023116"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```


### 4. archive: exit on src file error or filename error
Command input

running from source:

```
go run main.go -v=0 archive -ifc -ifnc -R -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/4
```

running compiled:

```
./openmedia -v=0 archive -ifc -ifnc -R -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/4
```

#### Command output

```
{"time":"2024-07-08T11:00:19.090007369+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2411149295/cmd/DST/cmd/archive/4","CompressionType":"zip","InvalidFilenameContinue":false,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-08T11:00:19.090234704+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-08T11:00:19.090267545+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/helper.ErrorsCodeMap.ExitWithCode","file":"errors.go","line":112},"msg":"filenames_validation: {\"file does not have xml extension\":[\"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/hello_invalid.txt\",\"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/hello_invalid2.txt\"],\"filename is not valid OpenMedia file\":[\"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/ohter_files/hello_invalid.xml\",\"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/ohter_files/hello_invalid2.xml\"]}"}
exit status 12
```


### 5. archive: run exit on src file error
Command input

running from source:

```
go run main.go -v=0 archive -ifc -R -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/5
```

running compiled:

```
./openmedia -v=0 archive -ifc -R -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/5
```

#### Command output

```
{"time":"2024-07-08T11:00:19.527261655+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2411149295/cmd/DST/cmd/archive/5","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-08T11:00:19.527493734+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-08T11:00:19.530106145+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/1/1/0"}
{"time":"2024-07-08T11:00:19.531134857+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T11:00:19.535351439+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/2/2/0"}
{"time":"2024-07-08T11:00:19.535357077+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T11:00:19.535378845+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T11:00:19.541044055+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T11:00:19.544714213+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/helper.ErrorsCodeMap.ExitWithCode","file":"errors.go","line":112},"msg":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
exit status 12
```


### 6. archive: do not exit on any file error
Command input

running from source:

```
go run main.go -v=0 archive -R -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/6
```

running compiled:

```
./openmedia -v=0 archive -R -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/6
```

#### Command output

```
{"time":"2024-07-08T11:00:19.915052305+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2411149295/cmd/DST/cmd/archive/6","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-08T11:00:19.915259912+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-08T11:00:19.919381282+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T11:00:19.9193899+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/1/1/0"}
{"time":"2024-07-08T11:00:19.923514681+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T11:00:19.924014452+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/2/2/0"}
{"time":"2024-07-08T11:00:19.92412582+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T11:00:19.928704656+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T11:00:19.932429763+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":315},"msg":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-08T11:00:20.440238871+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/4/3/1"}
{"time":"2024-07-08T11:00:20.44284669+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/5/4/1"}
{"time":"2024-07-08T11:00:20.444113543+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.087","MinifyRatio":"1.000","original":47218,"compressed":4096,"minified":47218,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/ohter_files/CT_Krpаlek__Lukач_2_977869_20231212143738.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Sakísuw,_Frmšú_Tuesday_W50_2023_12_12T143738.xml"}
{"time":"2024-07-08T11:00:20.448967265+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.127","original":47218,"compressed":0,"minified":6007,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/ohter_files/CT_Krpаlek__Lukач_2_977869_20231212143738.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Sakísuw,_Frmšú_Tuesday_W50_2023_12_12T143738.xml"}
{"time":"2024-07-08T11:00:20.606279327+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:21.128058618+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.077","MinifyRatio":"1.000","original":11960034,"compressed":921600,"minified":11960034,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/ohter_files/RD_05-09_ČRo_Sever_-_Wed__04_03_2020_2_1607196_20200304234759.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:21.397906282+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:21.398065682+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/6/5/1"}
{"time":"2024-07-08T11:00:21.713251087+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.033","MinifyRatio":"0.323","original":11960034,"compressed":393216,"minified":3865086,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/ohter_files/RD_05-09_ČRo_Sever_-_Wed__04_03_2020_2_1607196_20200304234759.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia2411149295/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia2411149295/cmd/SRC/cmd/archive/hello_invalid2.txt"],"filename is not valid OpenMedia file":["/tmp/openmedia2411149295/cmd/SRC/cmd/archive/ohter_files/hello_invalid.xml","/tmp/openmedia2411149295/cmd/SRC/cmd/archive/ohter_files/hello_invalid2.xml"]}
{"time":"2024-07-08T11:00:21.713339914+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}

result: archive_create
{"time":"2024-07-08T11:00:21.713359215+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/6/5/1"}
{"time":"2024-07-08T11:00:21.713373395+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.075","MinifyRatio":"1.000","original":25653580,"compressed":1912832,"minified":25653580,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia2411149295/cmd/DST/cmd/archive/6"}
{"time":"2024-07-08T11:00:21.713393671+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.030","MinifyRatio":"0.324","original":25653580,"compressed":778240,"minified":8312126,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia2411149295/cmd/DST/cmd/archive/6"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```


### 7. archive: do not recurse the source folder
Command input

running from source:

```
go run main.go -v=0 archive -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/7
```

running compiled:

```
./openmedia -v=0 archive -sdir=/tmp/openmedia2411149295/cmd/SRC/cmd/archive -odir=/tmp/openmedia2411149295/cmd/DST/cmd/archive/7
```

#### Command output

```
{"time":"2024-07-08T11:00:22.085576881+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2411149295/cmd/DST/cmd/archive/7","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-07-08T11:00:22.085769541+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-07-08T11:00:22.089170362+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T11:00:22.089330661+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-07-08T11:00:22.093533938+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T11:00:22.093765358+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-07-08T11:00:22.094441243+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T11:00:22.101159448+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T11:00:22.101909984+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":315},"msg":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-08T11:00:22.616582933+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-08T11:00:22.751642659+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:23.265544119+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia2411149295/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia2411149295/cmd/SRC/cmd/archive/hello_invalid2.txt"]}
{"time":"2024-07-08T11:00:23.265626787+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}

result: archive_create
{"time":"2024-07-08T11:00:23.26564354+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-08T11:00:23.265655537+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.072","MinifyRatio":"1.000","original":13646328,"compressed":987136,"minified":13646328,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia2411149295/cmd/DST/cmd/archive/7"}
{"time":"2024-07-08T11:00:23.265669366+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.028","MinifyRatio":"0.325","original":13646328,"compressed":385024,"minified":4441033,"filePathSource":"/tmp/openmedia2411149295/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia2411149295/cmd/DST/cmd/archive/7"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia2411149295/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```

### Run summary

--- PASS: TestRunCommand_Archive (7.13s)
    --- PASS: TestRunCommand_Archive/1 (0.41s)
    --- PASS: TestRunCommand_Archive/2 (0.42s)
    --- PASS: TestRunCommand_Archive/3 (1.68s)
    --- PASS: TestRunCommand_Archive/4 (0.42s)
    --- PASS: TestRunCommand_Archive/5 (0.45s)
    --- PASS: TestRunCommand_Archive/6 (2.17s)
    --- PASS: TestRunCommand_Archive/7 (1.55s)
PASS
ok  	github/czech-radio/openmedia/cmd	7.141s



## Command: extractArchive
### 1. extractArchive: help
Command input

running from source:

```
go run main.go extractArchive -h -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/1
```

running compiled:

```
./openmedia extractArchive -h -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/1
```

#### Command output

```
-h, -help
	display this help and exit

-sdir, -SourceDirectory=/mnt/remote/cro/export-avo/Rundowns
	Source directory of rundown files.

-sdirType, -SourceDirectoryType=MINIFIED.zip
	type of source directory where the rundowns resides

-odir, -OutputDirectory=/tmp/test/
	Destination directory or file

-ofname, -OutputFileName=
	Output file name.

-exsn, -ExtractorsName=production_all
	Name of extractor which specifies the parts of xml to be extracted

-fdf, -FilterDateFrom=
	Filter rundowns from date. Format of the date is given in form 'YYYY-mm-ddTHH:mm:ss' e.g. 2024, 2024-02-01 or 2024-02-01T10. The precission of date given is arbitrary.

-fdt, -FilterDateTo=
	Filter rundowns to date

-frn, -FilterRadioName=
	Filter radio names

-csvD, -CSVdelim=	
	csv column field delimiter
	[	 ;]

-arn, -AddRecordsNumbers=false
	Add record numbers columns and dependent columns

-frfn, -FilterFileName=
	Special filters filename. The filter filename specifies how the file is parsed and how it is used

-frsn, -FilterSheetName=data
	Special filters sheetname

-valfn, -ValidatorFileName=
	xlsx file containing validation receipe


```


### 2. extractArchive: extract all story parts from minified rundowns
Command input

running from source:

```
go run main.go extractArchive -ofname=production -exsn=production_all -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/2
```

running compiled:

```
./openmedia extractArchive -ofname=production -exsn=production_all -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/2
```

#### Command output

```
{"time":"2024-07-08T11:00:24.515357055+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":null,"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip|Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","38","5082"],"PartPrefixCode":11,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/2","OutputFileName":"production","ExtractorsName":"production_all","ExtractorsCode":"production_all","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"","FilterRecords":false,"ValidatorFileName":""}}
{"time":"2024-07-08T11:00:24.515787156+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-08T11:00:24.515805595+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":1}
{"time":"2024-07-08T11:00:24.515818552+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":2}
{"time":"2024-07-08T11:00:24.515928207+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T11:00:24.515951129+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:24.866432077+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":166}
{"time":"2024-07-08T11:00:24.866470532+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:25.143171568+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":147}
{"time":"2024-07-08T11:00:25.143339056+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T11:00:25.143357602+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T11:00:25.144678685+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T11:00:25.144696384+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T11:00:25.144797973+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":312}
{"time":"2024-07-08T11:00:25.144811266+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":312}
{"time":"2024-07-08T11:00:25.148987854+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/2/production_production_all_base_wh.csv","bytesCount":1125}
{"time":"2024-07-08T11:00:25.149084305+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/2/production_production_all_base_wh.csv","bytesCount":157961}
{"time":"2024-07-08T11:00:25.149128926+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/2/production_production_all_base_woh.csv","bytesCount":478}
{"time":"2024-07-08T11:00:25.149203053+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/2/production_production_all_base_woh.csv","bytesCount":157961}
{"time":"2024-07-08T11:00:25.149218065+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDataset","file":"outputs.go","line":27},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-08T11:00:25.149233621+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":49},"msg":"filter file not specified"}
```


### 3. extractArchive: extract all contacts from minified rundowns
Command input

running from source:

```
go run main.go extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/3
```

running compiled:

```
./openmedia extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/3
```

#### Command output

```
{"time":"2024-07-08T11:00:25.631065708+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":null,"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/3","OutputFileName":"production","ExtractorsName":"production_contacts","ExtractorsCode":"production_contacts","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"","FilterRecords":false,"ValidatorFileName":""}}
{"time":"2024-07-08T11:00:25.631363964+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-08T11:00:25.631379098+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":1}
{"time":"2024-07-08T11:00:25.63138551+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":2}
{"time":"2024-07-08T11:00:25.631470555+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T11:00:25.631488176+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:25.919148136+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":166}
{"time":"2024-07-08T11:00:25.919188732+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:26.210989806+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":147}
{"time":"2024-07-08T11:00:26.211169026+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T11:00:26.211186253+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T11:00:26.212590763+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T11:00:26.21261171+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T11:00:26.212724864+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-08T11:00:26.21273748+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-08T11:00:26.213333729+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-08T11:00:26.213348527+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-08T11:00:26.214104242+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/3/production_production_contacts_base_wh.csv","bytesCount":1051}
{"time":"2024-07-08T11:00:26.214157067+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/3/production_production_contacts_base_wh.csv","bytesCount":30295}
{"time":"2024-07-08T11:00:26.214201292+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/3/production_production_contacts_base_woh.csv","bytesCount":444}
{"time":"2024-07-08T11:00:26.214248755+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/3/production_production_contacts_base_woh.csv","bytesCount":30295}
{"time":"2024-07-08T11:00:26.214264064+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDataset","file":"outputs.go","line":27},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-08T11:00:26.2142793+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":49},"msg":"filter file not specified"}
```


### 4. extractArchive: extract all story parts from minified rundowns, extract only specified radios
Command input

running from source:

```
go run main.go extractArchive -ofname=production -exsn=production_all -frn=Olomouc -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/4
```

running compiled:

```
./openmedia extractArchive -ofname=production -exsn=production_all -frn=Olomouc -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/4
```

#### Command output

```
{"time":"2024-07-08T11:00:26.646682906+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":{"Olomouc":true},"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip|Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","38","5082"],"PartPrefixCode":11,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/4","OutputFileName":"production","ExtractorsName":"production_all","ExtractorsCode":"production_all","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"Olomouc","FilterRecords":false,"ValidatorFileName":""}}
{"time":"2024-07-08T11:00:26.647055862+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":0,"matched":true}
{"time":"2024-07-08T11:00:26.647070178+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":0}
{"time":"2024-07-08T11:00:26.6470767+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":0}
{"time":"2024-07-08T11:00:26.64714633+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T11:00:26.6471921+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-08T11:00:26.647199018+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-08T11:00:26.647205194+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-08T11:00:26.647210635+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-08T11:00:26.647215405+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-08T11:00:26.647220147+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-08T11:00:26.647412722+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/4/production_production_all_base_wh.csv","bytesCount":1125}
{"time":"2024-07-08T11:00:26.647436708+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/4/production_production_all_base_wh.csv","bytesCount":0}
{"time":"2024-07-08T11:00:26.647465145+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/4/production_production_all_base_woh.csv","bytesCount":478}
{"time":"2024-07-08T11:00:26.647495685+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/4/production_production_all_base_woh.csv","bytesCount":0}
{"time":"2024-07-08T11:00:26.647509132+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDataset","file":"outputs.go","line":27},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-08T11:00:26.647518922+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":49},"msg":"filter file not specified"}
```


### 5. extractArchive: extract all story parts from minified rundowns and validate
Command input

running from source:

```
go run main.go extractArchive -ofname=production -exsn=production_all -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5
```

running compiled:

```
./openmedia extractArchive -ofname=production -exsn=production_all -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5
```

#### Command output

```
{"time":"2024-07-08T11:00:27.030847615+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":null,"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip|Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","38","5082"],"PartPrefixCode":11,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5","OutputFileName":"production","ExtractorsName":"production_all","ExtractorsCode":"production_all","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"","FilterRecords":false,"ValidatorFileName":"../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx"}}
{"time":"2024-07-08T11:00:27.031178057+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-08T11:00:27.03119403+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":1}
{"time":"2024-07-08T11:00:27.031205834+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":2}
{"time":"2024-07-08T11:00:27.03129026+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T11:00:27.03131515+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:27.379644302+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":166}
{"time":"2024-07-08T11:00:27.379686362+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:27.657197495+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":147}
{"time":"2024-07-08T11:00:27.657360277+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T11:00:27.657383814+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T11:00:27.658681759+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T11:00:27.658697268+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T11:00:27.658799579+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":312}
{"time":"2024-07-08T11:00:27.658812975+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":312}
{"time":"2024-07-08T11:00:27.663091272+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5/production_production_all_base_wh.csv","bytesCount":1125}
{"time":"2024-07-08T11:00:27.663185045+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5/production_production_all_base_wh.csv","bytesCount":157961}
{"time":"2024-07-08T11:00:27.663227242+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5/production_production_all_base_woh.csv","bytesCount":478}
{"time":"2024-07-08T11:00:27.663303374+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5/production_production_all_base_woh.csv","bytesCount":157961}
{"time":"2024-07-08T11:00:27.725054156+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5/production_production_all_base_validated_wh.csv","bytesCount":1125}
{"time":"2024-07-08T11:00:27.725160002+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5/production_production_all_base_validated_wh.csv","bytesCount":150024}
{"time":"2024-07-08T11:00:27.725206735+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5/production_production_all_base_validated_woh.csv","bytesCount":478}
{"time":"2024-07-08T11:00:27.72527825+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5/production_production_all_base_validated_woh.csv","bytesCount":150024}
{"time":"2024-07-08T11:00:27.725509071+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":332},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/5/production_base_validated_log.csv","bytesCount":4948}
{"time":"2024-07-08T11:00:27.725527573+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":49},"msg":"filter file not specified"}
```


### 6. extractArchive: extract all story parts from minified rundowns, validate and use filter oposition
Command input

running from source:

```
go run main.go extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6
```

running compiled:

```
./openmedia extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6
```

#### Command output

```
{"time":"2024-07-08T11:00:28.159082678+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":null,"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6","OutputFileName":"production","ExtractorsName":"production_contacts","ExtractorsCode":"production_contacts","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"","FilterRecords":false,"ValidatorFileName":"../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx"}}
{"time":"2024-07-08T11:00:28.159402145+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-08T11:00:28.159416885+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":1}
{"time":"2024-07-08T11:00:28.159425229+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":2}
{"time":"2024-07-08T11:00:28.159522311+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T11:00:28.159541217+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:28.452014153+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":166}
{"time":"2024-07-08T11:00:28.452058795+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:28.717677954+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":147}
{"time":"2024-07-08T11:00:28.717813605+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T11:00:28.717825194+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T11:00:28.719129068+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T11:00:28.719142619+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T11:00:28.719222658+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-08T11:00:28.719231502+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-08T11:00:28.719759175+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-08T11:00:28.719771144+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-08T11:00:28.720391844+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_wh.csv","bytesCount":1051}
{"time":"2024-07-08T11:00:28.720445559+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_wh.csv","bytesCount":30295}
{"time":"2024-07-08T11:00:28.720477312+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_woh.csv","bytesCount":444}
{"time":"2024-07-08T11:00:28.720503867+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_woh.csv","bytesCount":30295}
{"time":"2024-07-08T11:00:28.781249265+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_wh.csv","bytesCount":1051}
{"time":"2024-07-08T11:00:28.781304429+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_wh.csv","bytesCount":25388}
{"time":"2024-07-08T11:00:28.781336274+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_woh.csv","bytesCount":444}
{"time":"2024-07-08T11:00:28.781361851+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_woh.csv","bytesCount":25388}
{"time":"2024-07-08T11:00:28.78148589+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":332},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_base_validated_log.csv","bytesCount":4256}
{"time":"2024-07-08T11:00:28.950375698+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":1212}
{"time":"2024-07-08T11:00:28.950449999+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":27808}
{"time":"2024-07-08T11:00:28.950494231+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":502}
{"time":"2024-07-08T11:00:28.950532527+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":27808}
```


### 7. extractArchive: extract all story parts from minified rundowns, validate and use filter eurovolby
Command input

running from source:

```
go run main.go extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7
```

running compiled:

```
./openmedia extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx -sdir=/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7
```

#### Command output

```
{"time":"2024-07-08T11:00:29.417114643+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":null,"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7","OutputFileName":"production","ExtractorsName":"production_contacts","ExtractorsCode":"production_contacts","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"","FilterRecords":false,"ValidatorFileName":"../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx"}}
{"time":"2024-07-08T11:00:29.417445794+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-08T11:00:29.41746066+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":1}
{"time":"2024-07-08T11:00:29.41747198+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":2}
{"time":"2024-07-08T11:00:29.417567445+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T11:00:29.417585647+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:29.705471201+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":166}
{"time":"2024-07-08T11:00:29.705517751+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia1145292749/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T11:00:29.982694572+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":147}
{"time":"2024-07-08T11:00:29.98283352+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T11:00:29.982845697+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T11:00:29.984151309+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T11:00:29.984165444+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T11:00:29.984245787+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-08T11:00:29.984253996+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-08T11:00:29.984755651+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-08T11:00:29.984765686+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-08T11:00:29.985350935+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_wh.csv","bytesCount":1051}
{"time":"2024-07-08T11:00:29.985385885+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_wh.csv","bytesCount":30295}
{"time":"2024-07-08T11:00:29.98541371+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_woh.csv","bytesCount":444}
{"time":"2024-07-08T11:00:29.985439769+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_woh.csv","bytesCount":30295}
{"time":"2024-07-08T11:00:30.042194495+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_wh.csv","bytesCount":1051}
{"time":"2024-07-08T11:00:30.04226305+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_wh.csv","bytesCount":25388}
{"time":"2024-07-08T11:00:30.042307503+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_woh.csv","bytesCount":444}
{"time":"2024-07-08T11:00:30.042344891+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_woh.csv","bytesCount":25388}
{"time":"2024-07-08T11:00:30.042479269+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":332},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_base_validated_log.csv","bytesCount":4256}
{"time":"2024-07-08T11:00:30.135445865+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":1179}
{"time":"2024-07-08T11:00:30.135502566+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":27544}
{"time":"2024-07-08T11:00:30.135531952+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":497}
{"time":"2024-07-08T11:00:30.135562259+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia1145292749/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":27544}
```

### Run summary

--- PASS: TestRunCommand_ExtractArchive (6.42s)
    --- PASS: TestRunCommand_ExtractArchive/1 (0.41s)
    --- PASS: TestRunCommand_ExtractArchive/2 (1.02s)
    --- PASS: TestRunCommand_ExtractArchive/3 (1.07s)
    --- PASS: TestRunCommand_ExtractArchive/4 (0.43s)
    --- PASS: TestRunCommand_ExtractArchive/5 (1.08s)
    --- PASS: TestRunCommand_ExtractArchive/6 (1.22s)
    --- PASS: TestRunCommand_ExtractArchive/7 (1.18s)
PASS
ok  	github/czech-radio/openmedia/cmd	6.426s
