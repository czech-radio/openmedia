


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
time=2024-07-28T23:20:41.008+02:00 level=ERROR source=sloger.go:53 msg="test error"
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
time=2024-07-28T23:20:41.536+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-28T23:20:41.536+02:00 level=ERROR source=sloger.go:53 msg="test error"
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
time=2024-07-28T23:20:41.981+02:00 level=INFO source=sloger.go:51 msg="test info"
time=2024-07-28T23:20:41.981+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-28T23:20:41.981+02:00 level=ERROR source=sloger.go:53 msg="test error"
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
time=2024-07-28T23:20:42.415+02:00 level=DEBUG source=config_mapany.go:253 msg="subcommand added" subname=archive
time=2024-07-28T23:20:42.415+02:00 level=DEBUG source=config_mapany.go:253 msg="subcommand added" subname=extractArchive
time=2024-07-28T23:20:42.415+02:00 level=DEBUG source=config_mapany.go:253 msg="subcommand added" subname=extractFile
time=2024-07-28T23:20:42.415+02:00 level=DEBUG source=sloger.go:50 msg="test debug"
time=2024-07-28T23:20:42.415+02:00 level=INFO source=sloger.go:51 msg="test info"
time=2024-07-28T23:20:42.415+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-28T23:20:42.415+02:00 level=ERROR source=sloger.go:53 msg="test error"
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
{"time":"2024-07-28T23:20:42.845497748+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":253},"msg":"subcommand added","subname":"archive"}
{"time":"2024-07-28T23:20:42.845582146+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":253},"msg":"subcommand added","subname":"extractArchive"}
{"time":"2024-07-28T23:20:42.845593739+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":253},"msg":"subcommand added","subname":"extractFile"}
{"time":"2024-07-28T23:20:42.845605111+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":50},"msg":"test debug"}
{"time":"2024-07-28T23:20:42.845616842+02:00","level":"INFO","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":51},"msg":"test info"}
{"time":"2024-07-28T23:20:42.845626852+02:00","level":"WARN","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":52},"msg":"test warn"}
{"time":"2024-07-28T23:20:42.845643016+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":53},"msg":"test error"}
```

### Run summary

--- PASS: TestRunCommand_Root (4.21s)
    --- PASS: TestRunCommand_Root/1 (0.47s)
    --- PASS: TestRunCommand_Root/2 (0.43s)
    --- PASS: TestRunCommand_Root/3 (0.43s)
    --- PASS: TestRunCommand_Root/4 (0.50s)
    --- PASS: TestRunCommand_Root/5 (0.52s)
    --- PASS: TestRunCommand_Root/6 (0.53s)
    --- PASS: TestRunCommand_Root/7 (0.45s)
    --- PASS: TestRunCommand_Root/8 (0.43s)
    --- PASS: TestRunCommand_Root/9 (0.43s)
PASS
ok  	github/czech-radio/openmedia/cmd	4.222s



## Command: archive
### 1. archive: help
Command input

running from source:

```
go run main.go archive -h -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/1
```

running compiled:

```
./openmedia archive -h -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/1
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
go run main.go -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/2
```

running compiled:

```
./openmedia -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/2
```

#### Command output

```
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
Archive config: {SourceDirectory:/tmp/openmedia1761589961/cmd/SRC/cmd/archive OutputDirectory:/tmp/openmedia1761589961/cmd/DST/cmd/archive/2 CompressionType:zip InvalidFilenameContinue:true InvalidFileContinue:false InvalidFileRename:false ProcessedFileRename:false ProcessedFileDelete:false PreserveFoldersInArchive:false RecurseSourceDirectory:true}
```


### 3. archive: dry run
Command input

running from source:

```
go run main.go -dr -v=0 archive -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/3
```

running compiled:

```
./openmedia -dr -v=0 archive -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/3
```

#### Command output

```
{"time":"2024-07-28T23:20:44.592031103+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":62},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia1761589961/cmd/DST/cmd/archive/3","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-07-28T23:20:44.59215785+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":65},"msg":"dry run activated","output_path":"/tmp/openmedia_archive2788001706"}
{"time":"2024-07-28T23:20:44.592251122+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-07-28T23:20:44.595397622+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-28T23:20:44.595646524+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-07-28T23:20:44.600852812+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-07-28T23:20:44.600481766+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-28T23:20:44.601901982+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-28T23:20:44.606735392+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-28T23:20:44.618312009+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":296},"msg":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-28T23:20:45.21873821+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-28T23:20:45.338416435+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:45.974813267+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia1761589961/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia1761589961/cmd/SRC/cmd/archive/hello_invalid2.txt"]}
{"time":"2024-07-28T23:20:45.974889665+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}

result: archive_create
{"time":"2024-07-28T23:20:45.974904321+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-28T23:20:45.974916407+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.072","MinifyRatio":"1.000","original":13646328,"compressed":987136,"minified":13646328,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia_archive2788001706"}
{"time":"2024-07-28T23:20:45.974930035+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.028","MinifyRatio":"0.325","original":13646328,"compressed":385024,"minified":4441033,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia_archive2788001706"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```


### 4. archive: exit on src file error or filename error
Command input

running from source:

```
go run main.go -v=0 archive -ifc -ifnc -R -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/4
```

running compiled:

```
./openmedia -v=0 archive -ifc -ifnc -R -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/4
```

#### Command output

```
{"time":"2024-07-28T23:20:46.374977543+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":62},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia1761589961/cmd/DST/cmd/archive/4","CompressionType":"zip","InvalidFilenameContinue":false,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-28T23:20:46.37519537+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-28T23:20:46.375230167+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/helper.ErrorsCodeMap.ExitWithCode","file":"errors.go","line":112},"msg":"filenames_validation: {\"file does not have xml extension\":[\"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/hello_invalid.txt\",\"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/hello_invalid2.txt\"],\"filename is not valid OpenMedia file\":[\"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/ohter_files/hello_invalid.xml\",\"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/ohter_files/hello_invalid2.xml\"]}"}
exit status 12
```


### 5. archive: run exit on src file error
Command input

running from source:

```
go run main.go -v=0 archive -ifc -R -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/5
```

running compiled:

```
./openmedia -v=0 archive -ifc -R -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/5
```

#### Command output

```
{"time":"2024-07-28T23:20:46.817548109+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":62},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia1761589961/cmd/DST/cmd/archive/5","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-28T23:20:46.817759433+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-28T23:20:46.820924282+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/1/1/0"}
{"time":"2024-07-28T23:20:46.821016512+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-28T23:20:46.824956867+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-28T23:20:46.825316637+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/2/2/0"}
{"time":"2024-07-28T23:20:46.827003724+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-28T23:20:46.832728722+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-28T23:20:46.842819283+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/helper.ErrorsCodeMap.ExitWithCode","file":"errors.go","line":112},"msg":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
exit status 12
```


### 6. archive: do not exit on any file error
Command input

running from source:

```
go run main.go -v=0 archive -R -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/6
```

running compiled:

```
./openmedia -v=0 archive -R -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/6
```

#### Command output

```
{"time":"2024-07-28T23:20:47.289305544+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":62},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia1761589961/cmd/DST/cmd/archive/6","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-28T23:20:47.289540671+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-28T23:20:47.292343552+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/1/1/0"}
{"time":"2024-07-28T23:20:47.292932835+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-28T23:20:47.298068616+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-28T23:20:47.298215249+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/2/2/0"}
{"time":"2024-07-28T23:20:47.300120799+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-28T23:20:47.305658764+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-28T23:20:47.318672733+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":296},"msg":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-28T23:20:47.990029735+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/4/3/1"}
{"time":"2024-07-28T23:20:47.997124978+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/5/4/1"}
{"time":"2024-07-28T23:20:48.017646403+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.127","original":47218,"compressed":0,"minified":6007,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/ohter_files/CT_Krpаlek__Lukач_2_977869_20231212143738.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Sakísuw,_Frmšú_Tuesday_W50_2023_12_12T143738.xml"}
{"time":"2024-07-28T23:20:48.019350014+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.087","MinifyRatio":"1.000","original":47218,"compressed":4096,"minified":47218,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/ohter_files/CT_Krpаlek__Lukач_2_977869_20231212143738.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Sakísuw,_Frmšú_Tuesday_W50_2023_12_12T143738.xml"}
{"time":"2024-07-28T23:20:48.273750252+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:49.032339429+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.077","MinifyRatio":"1.000","original":11960034,"compressed":921600,"minified":11960034,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/ohter_files/RD_05-09_ČRo_Sever_-_Wed__04_03_2020_2_1607196_20200304234759.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:49.176035787+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:49.176193694+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/6/5/1"}
{"time":"2024-07-28T23:20:49.569172047+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.033","MinifyRatio":"0.323","original":11960034,"compressed":393216,"minified":3865086,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/ohter_files/RD_05-09_ČRo_Sever_-_Wed__04_03_2020_2_1607196_20200304234759.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia1761589961/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia1761589961/cmd/SRC/cmd/archive/hello_invalid2.txt"],"filename is not valid OpenMedia file":["/tmp/openmedia1761589961/cmd/SRC/cmd/archive/ohter_files/hello_invalid.xml","/tmp/openmedia1761589961/cmd/SRC/cmd/archive/ohter_files/hello_invalid2.xml"]}
{"time":"2024-07-28T23:20:49.569248165+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}

result: archive_create
{"time":"2024-07-28T23:20:49.569265631+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/6/5/1"}
{"time":"2024-07-28T23:20:49.56927931+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.075","MinifyRatio":"1.000","original":25653580,"compressed":1912832,"minified":25653580,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia1761589961/cmd/DST/cmd/archive/6"}
{"time":"2024-07-28T23:20:49.56929318+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.030","MinifyRatio":"0.324","original":25653580,"compressed":778240,"minified":8312126,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia1761589961/cmd/DST/cmd/archive/6"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```


### 7. archive: do not recurse the source folder
Command input

running from source:

```
go run main.go -v=0 archive -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/7
```

running compiled:

```
./openmedia -v=0 archive -sdir=/tmp/openmedia1761589961/cmd/SRC/cmd/archive -odir=/tmp/openmedia1761589961/cmd/DST/cmd/archive/7
```

#### Command output

```
{"time":"2024-07-28T23:20:49.987139882+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":62},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia1761589961/cmd/DST/cmd/archive/7","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-07-28T23:20:49.987348012+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-07-28T23:20:49.991071443+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-28T23:20:49.991261823+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-07-28T23:20:49.995328422+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-28T23:20:49.996118879+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-28T23:20:49.996275483+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-07-28T23:20:50.002375924+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-28T23:20:50.011729864+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":296},"msg":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-28T23:20:50.593710538+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-28T23:20:50.716298458+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:51.37464454+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia1761589961/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia1761589961/cmd/SRC/cmd/archive/hello_invalid2.txt"]}
{"time":"2024-07-28T23:20:51.37474885+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}

result: archive_create
{"time":"2024-07-28T23:20:51.374763453+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-28T23:20:51.374775976+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.072","MinifyRatio":"1.000","original":13646328,"compressed":987136,"minified":13646328,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia1761589961/cmd/DST/cmd/archive/7"}
{"time":"2024-07-28T23:20:51.374789975+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":455},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.028","MinifyRatio":"0.325","original":13646328,"compressed":385024,"minified":4441033,"filePathSource":"/tmp/openmedia1761589961/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia1761589961/cmd/DST/cmd/archive/7"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia1761589961/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```

### Run summary

--- PASS: TestRunCommand_Archive (8.05s)
    --- PASS: TestRunCommand_Archive/1 (0.39s)
    --- PASS: TestRunCommand_Archive/2 (0.46s)
    --- PASS: TestRunCommand_Archive/3 (1.79s)
    --- PASS: TestRunCommand_Archive/4 (0.39s)
    --- PASS: TestRunCommand_Archive/5 (0.47s)
    --- PASS: TestRunCommand_Archive/6 (2.74s)
    --- PASS: TestRunCommand_Archive/7 (1.80s)
PASS
ok  	github/czech-radio/openmedia/cmd	8.059s



## Command: extractArchive
### 1. extractArchive: help
Command input

running from source:

```
go run main.go extractArchive -h -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/2
```

running compiled:

```
./openmedia extractArchive -h -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/2
```

#### Command output

```
-h, -help
	display this help and exit

-odir, -OutputDirectory=
	Output file path for extracted data.

-ofn, -OutputFileName=
	Output file path for extracted data.

-csvD, -CSVdelim=	
	csv column field delimiter
	[	 ;]

-excode, -ExtractorsCode=production_all
	Name of extractor which specifies the parts of xml to be extracted

-frns, -FilterRadioNames=
	Filter data corresponding to radio names

-fdf, -FilterDateFrom=2024-07-15 00:00:00 +0200 CEST
	Filter rundowns from date. Format of the date is given in form 'YYYY-mm-ddTHH:mm:ss' e.g. 2024, 2024-02-01 or 2024-02-01T10. The precission of date given is arbitrary.

-fdt, -FilterDateTo=2024-07-22 00:00:00 +0200 CEST
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
go run main.go extractArchive -ofn=production -excode=production_all -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/3
```

running compiled:

```
./openmedia extractArchive -ofn=production -excode=production_all -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/3
```

#### Command output

```
{ExtractorsCode:production_all FilterRadioNames:map[] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName: AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/3 OutputFileName:production CSVdelim:	}
{FilterFileName: FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-28T23:20:52.763627509+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-28T23:20:52.763734038+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":218},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-28T23:20:52.763751658+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":1}
{"time":"2024-07-28T23:20:52.763765976+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":2}
{"time":"2024-07-28T23:20:52.763881242+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-28T23:20:52.763903144+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:53.156069177+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":166}
{"time":"2024-07-28T23:20:53.156109987+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:53.475029358+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":147}
{"time":"2024-07-28T23:20:53.475182449+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-28T23:20:53.475205744+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-28T23:20:53.476502543+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-28T23:20:53.476516953+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-28T23:20:53.476607071+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":312}
{"time":"2024-07-28T23:20:53.4766269+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":312}
{"time":"2024-07-28T23:20:53.480342789+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/3/production_production_all_base_wh.csv","bytesCount":1100}
{"time":"2024-07-28T23:20:53.480427892+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/3/production_production_all_base_wh.csv","bytesCount":154684}
{"time":"2024-07-28T23:20:53.480462658+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/3/production_production_all_base_woh.csv","bytesCount":468}
{"time":"2024-07-28T23:20:53.480545678+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/3/production_production_all_base_woh.csv","bytesCount":154684}
{"time":"2024-07-28T23:20:53.480560934+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDatasetNew","file":"outputs.go","line":51},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-28T23:20:53.48057782+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDatasetNew","file":"outputs.go","line":74},"msg":"filter file not specified"}
```


### 3. extractArchive: extract all contacts from minified rundowns
Command input

running from source:

```
go run main.go extractArchive -ofn=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/4
```

running compiled:

```
./openmedia extractArchive -ofn=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/4
```

#### Command output

```
{ExtractorsCode:production_contacts FilterRadioNames:map[] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName: AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/4 OutputFileName:production CSVdelim:	}
{FilterFileName: FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-28T23:20:54.121195335+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-28T23:20:54.12128852+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":218},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-28T23:20:54.121303765+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":1}
{"time":"2024-07-28T23:20:54.121316901+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":2}
{"time":"2024-07-28T23:20:54.121419203+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-28T23:20:54.121442897+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:54.464350946+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":166}
{"time":"2024-07-28T23:20:54.464397532+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:54.845236451+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":147}
{"time":"2024-07-28T23:20:54.845393004+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-28T23:20:54.845417963+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-28T23:20:54.846753712+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-28T23:20:54.846774106+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-28T23:20:54.846864549+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-28T23:20:54.846876691+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-28T23:20:54.846945409+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-28T23:20:54.84695719+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-28T23:20:54.847543436+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/4/production_production_contacts_base_wh.csv","bytesCount":1026}
{"time":"2024-07-28T23:20:54.847594822+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/4/production_production_contacts_base_wh.csv","bytesCount":29762}
{"time":"2024-07-28T23:20:54.847635873+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/4/production_production_contacts_base_woh.csv","bytesCount":434}
{"time":"2024-07-28T23:20:54.84768352+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/4/production_production_contacts_base_woh.csv","bytesCount":29762}
{"time":"2024-07-28T23:20:54.847697796+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDatasetNew","file":"outputs.go","line":51},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-28T23:20:54.847711143+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDatasetNew","file":"outputs.go","line":74},"msg":"filter file not specified"}
```


### 4. extractArchive: extract all story parts from minified rundowns, extract only specified radios
Command input

running from source:

```
go run main.go extractArchive -ofn=production -excode=production_all -frns=Olomouc,Plus -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/5
```

running compiled:

```
./openmedia extractArchive -ofn=production -excode=production_all -frns=Olomouc,Plus -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/5
```

#### Command output

```
{ExtractorsCode:production_all FilterRadioNames:map[Olomouc:true Plus:true] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName: AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/5 OutputFileName:production CSVdelim:	}
{FilterFileName: FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-28T23:20:55.331622852+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
FUCE ŘRf_Tcunyso 2
FUCE ŠKs_Wfbvy 2
{"time":"2024-07-28T23:20:55.331740364+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":218},"msg":"filenames in all packages","count":0,"matched":true}
{"time":"2024-07-28T23:20:55.331758243+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":0}
{"time":"2024-07-28T23:20:55.331768925+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":0}
{"time":"2024-07-28T23:20:55.331907618+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-28T23:20:55.331978177+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-28T23:20:55.331990795+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-28T23:20:55.33200065+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-28T23:20:55.332009111+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-28T23:20:55.332018036+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-28T23:20:55.332026589+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-28T23:20:55.332110831+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/5/production_production_all_base_wh.csv","bytesCount":1100}
{"time":"2024-07-28T23:20:55.332143374+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/5/production_production_all_base_wh.csv","bytesCount":0}
{"time":"2024-07-28T23:20:55.332180646+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/5/production_production_all_base_woh.csv","bytesCount":468}
{"time":"2024-07-28T23:20:55.332202826+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/5/production_production_all_base_woh.csv","bytesCount":0}
{"time":"2024-07-28T23:20:55.332214319+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDatasetNew","file":"outputs.go","line":51},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-28T23:20:55.332226624+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDatasetNew","file":"outputs.go","line":74},"msg":"filter file not specified"}
```


### 5. extractArchive: extract all story parts from minified rundowns and validate
Command input

running from source:

```
go run main.go extractArchive -ofn=production -excode=production_all -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6
```

running compiled:

```
./openmedia extractArchive -ofn=production -excode=production_all -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6
```

#### Command output

```
{ExtractorsCode:production_all FilterRadioNames:map[] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName:../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6 OutputFileName:production CSVdelim:	}
{FilterFileName: FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-28T23:20:55.73513892+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-28T23:20:55.735238245+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":218},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-28T23:20:55.735254008+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":1}
{"time":"2024-07-28T23:20:55.735269588+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":2}
{"time":"2024-07-28T23:20:55.735374613+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-28T23:20:55.735396481+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:56.12667911+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":166}
{"time":"2024-07-28T23:20:56.126756563+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:56.499495584+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":147}
{"time":"2024-07-28T23:20:56.499703891+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-28T23:20:56.499736257+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-28T23:20:56.501704534+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-28T23:20:56.501738382+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-28T23:20:56.501887989+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":312}
{"time":"2024-07-28T23:20:56.501914742+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":312}
{"time":"2024-07-28T23:20:56.507276355+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6/production_production_all_base_wh.csv","bytesCount":1100}
{"time":"2024-07-28T23:20:56.507397605+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6/production_production_all_base_wh.csv","bytesCount":154684}
{"time":"2024-07-28T23:20:56.507441391+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6/production_production_all_base_woh.csv","bytesCount":468}
{"time":"2024-07-28T23:20:56.507522067+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6/production_production_all_base_woh.csv","bytesCount":154684}
{"time":"2024-07-28T23:20:56.594187706+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6/production_production_all_base_validated_wh.csv","bytesCount":1100}
{"time":"2024-07-28T23:20:56.594290605+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6/production_production_all_base_validated_wh.csv","bytesCount":146747}
{"time":"2024-07-28T23:20:56.594331053+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6/production_production_all_base_validated_woh.csv","bytesCount":468}
{"time":"2024-07-28T23:20:56.594392393+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6/production_production_all_base_validated_woh.csv","bytesCount":146747}
{"time":"2024-07-28T23:20:56.594575257+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":332},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/6/production_base_validated_log.csv","bytesCount":4948}
{"time":"2024-07-28T23:20:56.594588316+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDatasetNew","file":"outputs.go","line":74},"msg":"filter file not specified"}
```


### 6. extractArchive: extract all story parts from minified rundowns, validate and use filter oposition
Command input

running from source:

```
go run main.go extractArchive -ofn=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7
```

running compiled:

```
./openmedia extractArchive -ofn=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7
```

#### Command output

```
{ExtractorsCode:production_contacts FilterRadioNames:map[] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName:../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7 OutputFileName:production CSVdelim:	}
{FilterFileName:../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-28T23:20:57.055527864+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-28T23:20:57.055626156+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":218},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-28T23:20:57.055642324+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":1}
{"time":"2024-07-28T23:20:57.055651947+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":2}
{"time":"2024-07-28T23:20:57.055775439+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-28T23:20:57.055803331+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:57.370738478+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":166}
{"time":"2024-07-28T23:20:57.370774044+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:57.641549772+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":147}
{"time":"2024-07-28T23:20:57.641681325+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-28T23:20:57.641697858+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-28T23:20:57.642984534+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-28T23:20:57.642998108+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-28T23:20:57.64308901+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-28T23:20:57.64309902+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-28T23:20:57.643157663+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-28T23:20:57.643166329+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-28T23:20:57.643781473+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_wh.csv","bytesCount":1026}
{"time":"2024-07-28T23:20:57.643817841+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_wh.csv","bytesCount":29762}
{"time":"2024-07-28T23:20:57.643846366+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_woh.csv","bytesCount":434}
{"time":"2024-07-28T23:20:57.643884549+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_woh.csv","bytesCount":29762}
{"time":"2024-07-28T23:20:57.705941416+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_wh.csv","bytesCount":1026}
{"time":"2024-07-28T23:20:57.706008182+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_wh.csv","bytesCount":24855}
{"time":"2024-07-28T23:20:57.70605296+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_woh.csv","bytesCount":434}
{"time":"2024-07-28T23:20:57.706094252+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_woh.csv","bytesCount":24855}
{"time":"2024-07-28T23:20:57.706231602+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":332},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_base_validated_log.csv","bytesCount":4256}
{"time":"2024-07-28T23:20:57.878353024+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":1187}
{"time":"2024-07-28T23:20:57.878420609+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":27275}
{"time":"2024-07-28T23:20:57.878465917+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":492}
{"time":"2024-07-28T23:20:57.878511032+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":27275}
```


### 7. extractArchive: extract all story parts from minified rundowns, validate and use filter eurovolby
Command input

running from source:

```
go run main.go extractArchive -ofn=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8
```

running compiled:

```
./openmedia extractArchive -ofn=production -excode=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx -sdir=/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8
```

#### Command output

```
{ExtractorsCode:production_contacts FilterRadioNames:map[] FilterDateFrom:2020-03-04 00:00:00 +0100 CET FilterDateTo:2020-03-05 00:00:00 +0100 CET DateRange:[2020-03-04 00:00:00.000000001 +0000 UTC 2020-03-05 00:00:00 +0000 UTC] FilterIsoWeeks:map[] FilterMonths:map[] FilterWeekDays:map[] ValidatorFileName:../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx AddRecordNumbers:false ComputeUniqueRows:false}
{SourceDirectory:/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive SourceDirectoryType:MINIFIED.zip SourceFilePath: SourceCharEncoding: OutputDirectory:/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8 OutputFileName:production CSVdelim:	}
{FilterFileName:../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx FilterSheetName:data ColumnHeaderRow:0 RowHeaderColumn:0}
{"time":"2024-07-28T23:20:58.317314652+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":97},"msg":"packages matched","count":1}
{"time":"2024-07-28T23:20:58.317400631+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":218},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-28T23:20:58.317415964+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":108},"msg":"packages matched","count":1}
{"time":"2024-07-28T23:20:58.317428916+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":109},"msg":"files inside packages matched","count":2}
{"time":"2024-07-28T23:20:58.317550069+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":119},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-28T23:20:58.317568352+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:58.640441148+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":166}
{"time":"2024-07-28T23:20:58.640479609+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":134},"msg":"proccessing package","package":"/tmp/openmedia850730381/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-28T23:20:58.932813953+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":144},"msg":"extracted lines","count":147}
{"time":"2024-07-28T23:20:58.933090678+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-28T23:20:58.933128181+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-28T23:20:58.935681863+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-28T23:20:58.935724199+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-28T23:20:58.935943287+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-28T23:20:58.935965437+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-28T23:20:58.936098344+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-28T23:20:58.936118051+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-28T23:20:58.940595928+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_wh.csv","bytesCount":1026}
{"time":"2024-07-28T23:20:58.940702183+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_wh.csv","bytesCount":29762}
{"time":"2024-07-28T23:20:58.940758145+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_woh.csv","bytesCount":434}
{"time":"2024-07-28T23:20:58.94080304+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_woh.csv","bytesCount":29762}
{"time":"2024-07-28T23:20:59.032657511+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_wh.csv","bytesCount":1026}
{"time":"2024-07-28T23:20:59.032709089+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_wh.csv","bytesCount":24855}
{"time":"2024-07-28T23:20:59.032741263+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_woh.csv","bytesCount":434}
{"time":"2024-07-28T23:20:59.032769028+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_woh.csv","bytesCount":24855}
{"time":"2024-07-28T23:20:59.032893251+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":332},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_base_validated_log.csv","bytesCount":4256}
{"time":"2024-07-28T23:20:59.133382212+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":1154}
{"time":"2024-07-28T23:20:59.133437622+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":27011}
{"time":"2024-07-28T23:20:59.133475456+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":487}
{"time":"2024-07-28T23:20:59.133510237+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia850730381/cmd/DST/cmd/extractArchive/8/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":27011}
```

### Run summary

--- PASS: TestRunCommand_ExtractArchive (7.33s)
    --- PASS: TestRunCommand_ExtractArchive/1 (0.42s)
    --- PASS: TestRunCommand_ExtractArchive/2 (1.26s)
    --- PASS: TestRunCommand_ExtractArchive/3 (1.36s)
    --- PASS: TestRunCommand_ExtractArchive/4 (0.48s)
    --- PASS: TestRunCommand_ExtractArchive/5 (1.27s)
    --- PASS: TestRunCommand_ExtractArchive/6 (1.28s)
    --- PASS: TestRunCommand_ExtractArchive/7 (1.25s)
PASS
ok  	github/czech-radio/openmedia/cmd	7.335s
