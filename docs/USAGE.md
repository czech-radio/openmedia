


## Command: root
### 1. root: help
Command input:

running from source:
```
go run main.go -h
```

running compiled:
```
./openmedia -h
```

#### Command output:
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
Command input:

running from source:
```
go run main.go -V
```

running compiled:
```
./openmedia -V
```

#### Command output:
```
{ProgramName: Version:1.0.1 GitTag: GitCommit: BuildTime:}
```

### 3. root: print config
Command input:

running from source:
```
go run main.go -dc
```

running compiled:
```
./openmedia -dc
```

#### Command output:
```
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
```

### 4. root: print version and config
Command input:

running from source:
```
go run main.go -V -dc
```

running compiled:
```
./openmedia -V -dc
```

#### Command output:
```
{ProgramName: Version:1.0.1 GitTag: GitCommit: BuildTime:}
Root config: &{Usage:false GeneralHelp:false Version:true Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
```

### 5. root: test log [err]
Command input:

running from source:
```
go run main.go -logt=plain -logts -v=6
```

running compiled:
```
./openmedia -logt=plain -logts -v=6
```

#### Command output:
```
time=2024-07-08T09:49:10.683+02:00 level=ERROR source=sloger.go:53 msg="test error"
```

### 6. root: test log [err,warn]
Command input:

running from source:
```
go run main.go -logt=plain -logts -v=4
```

running compiled:
```
./openmedia -logt=plain -logts -v=4
```

#### Command output:
```
time=2024-07-08T09:49:11.191+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-08T09:49:11.191+02:00 level=ERROR source=sloger.go:53 msg="test error"
```

### 7. root: test log [err,warn,info]
Command input:

running from source:
```
go run main.go -logt=plain -logts -v=0
```

running compiled:
```
./openmedia -logt=plain -logts -v=0
```

#### Command output:
```
time=2024-07-08T09:49:11.597+02:00 level=INFO source=sloger.go:51 msg="test info"
time=2024-07-08T09:49:11.597+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-08T09:49:11.597+02:00 level=ERROR source=sloger.go:53 msg="test error"
```

### 8. root: terst log [err,warn,info,debug]
Command input:

running from source:
```
go run main.go -logt=plain -logts -v=-4
```

running compiled:
```
./openmedia -logt=plain -logts -v=-4
```

#### Command output:
```
time=2024-07-08T09:49:12.145+02:00 level=DEBUG source=config_mapany.go:349 msg="subcommand added" subname=archive
time=2024-07-08T09:49:12.145+02:00 level=DEBUG source=config_mapany.go:349 msg="subcommand added" subname=extractArchive
time=2024-07-08T09:49:12.145+02:00 level=DEBUG source=sloger.go:50 msg="test debug"
time=2024-07-08T09:49:12.145+02:00 level=INFO source=sloger.go:51 msg="test info"
time=2024-07-08T09:49:12.145+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-07-08T09:49:12.145+02:00 level=ERROR source=sloger.go:53 msg="test error"
```

### 9. root: test log json
Command input:

running from source:
```
go run main.go -logt=json -logts -v=-4
```

running compiled:
```
./openmedia -logt=json -logts -v=-4
```

#### Command output:
```
{"time":"2024-07-08T09:49:12.611250656+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":349},"msg":"subcommand added","subname":"archive"}
{"time":"2024-07-08T09:49:12.6113316+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":349},"msg":"subcommand added","subname":"extractArchive"}
{"time":"2024-07-08T09:49:12.611348339+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":50},"msg":"test debug"}
{"time":"2024-07-08T09:49:12.611367054+02:00","level":"INFO","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":51},"msg":"test info"}
{"time":"2024-07-08T09:49:12.611376491+02:00","level":"WARN","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":52},"msg":"test warn"}
{"time":"2024-07-08T09:49:12.611384803+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":53},"msg":"test error"}
```
### Run summary

--- PASS: TestRunCommand_Root (4.00s)
    --- PASS: TestRunCommand_Root/1 (0.45s)
    --- PASS: TestRunCommand_Root/2 (0.44s)
    --- PASS: TestRunCommand_Root/3 (0.41s)
    --- PASS: TestRunCommand_Root/4 (0.39s)
    --- PASS: TestRunCommand_Root/5 (0.37s)
    --- PASS: TestRunCommand_Root/6 (0.51s)
    --- PASS: TestRunCommand_Root/7 (0.41s)
    --- PASS: TestRunCommand_Root/8 (0.55s)
    --- PASS: TestRunCommand_Root/9 (0.47s)
PASS
ok  	github/czech-radio/openmedia/cmd	4.012s



## Command: archive
### 1. archive: help
Command input:

running from source:
```
go run main.go archive -h -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/1
```

running compiled:
```
./openmedia archive -h -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/1
```

#### Command output:
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
Command input:

running from source:
```
go run main.go -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/2
```

running compiled:
```
./openmedia -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/2
```

#### Command output:
```
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
Archive config: {SourceDirectory:/tmp/openmedia1694863704/cmd/SRC/cmd/archive OutputDirectory:/tmp/openmedia1694863704/cmd/DST/cmd/archive/2 CompressionType:zip InvalidFilenameContinue:true InvalidFileContinue:false InvalidFileRename:false ProcessedFileRename:false ProcessedFileDelete:false PreserveFoldersInArchive:false RecurseSourceDirectory:true}
```

### 3. archive: dry run
Command input:

running from source:
```
go run main.go -dr -v=0 archive -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/3
```

running compiled:
```
./openmedia -dr -v=0 archive -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/3
```

#### Command output:
```
{"time":"2024-07-08T09:49:14.272697717+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia1694863704/cmd/DST/cmd/archive/3","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-07-08T09:49:14.27283887+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":66},"msg":"dry run activated","output_path":"/tmp/openmedia_archive1408910581"}
{"time":"2024-07-08T09:49:14.27293352+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-07-08T09:49:14.279873467+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-07-08T09:49:14.280587493+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T09:49:14.287203683+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T09:49:14.287374076+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T09:49:14.287508402+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-07-08T09:49:14.294742965+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T09:49:14.298651432+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":315},"msg":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-08T09:49:14.907765764+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-08T09:49:15.031017533+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:15.534991027+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia1694863704/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia1694863704/cmd/SRC/cmd/archive/hello_invalid2.txt"]}
{"time":"2024-07-08T09:49:15.535062518+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}

result: archive_create
{"time":"2024-07-08T09:49:15.53507408+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-08T09:49:15.535083465+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.072","MinifyRatio":"1.000","original":13646328,"compressed":987136,"minified":13646328,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia_archive1408910581"}
{"time":"2024-07-08T09:49:15.535095533+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.028","MinifyRatio":"0.325","original":13646328,"compressed":385024,"minified":4441033,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia_archive1408910581"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```

### 4. archive: exit on src file error or filename error
Command input:

running from source:
```
go run main.go -v=0 archive -ifc -ifnc -R -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/4
```

running compiled:
```
./openmedia -v=0 archive -ifc -ifnc -R -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/4
```

#### Command output:
```
{"time":"2024-07-08T09:49:16.012905471+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia1694863704/cmd/DST/cmd/archive/4","CompressionType":"zip","InvalidFilenameContinue":false,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-08T09:49:16.01327425+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-08T09:49:16.013318824+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/helper.ErrorsCodeMap.ExitWithCode","file":"errors.go","line":112},"msg":"filenames_validation: {\"file does not have xml extension\":[\"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/hello_invalid.txt\",\"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/hello_invalid2.txt\"],\"filename is not valid OpenMedia file\":[\"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/ohter_files/hello_invalid.xml\",\"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/ohter_files/hello_invalid2.xml\"]}"}
exit status 12
```

### 5. archive: run exit on src file error
Command input:

running from source:
```
go run main.go -v=0 archive -ifc -R -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/5
```

running compiled:
```
./openmedia -v=0 archive -ifc -R -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/5
```

#### Command output:
```
{"time":"2024-07-08T09:49:16.417937345+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia1694863704/cmd/DST/cmd/archive/5","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-08T09:49:16.418188631+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-08T09:49:16.422426338+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T09:49:16.42270813+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/1/1/0"}
{"time":"2024-07-08T09:49:16.428142387+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T09:49:16.42812945+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T09:49:16.428398587+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/2/2/0"}
{"time":"2024-07-08T09:49:16.435956134+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T09:49:16.443647397+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/helper.ErrorsCodeMap.ExitWithCode","file":"errors.go","line":112},"msg":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
exit status 12
```

### 6. archive: do not exit on any file error
Command input:

running from source:
```
go run main.go -v=0 archive -R -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/6
```

running compiled:
```
./openmedia -v=0 archive -R -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/6
```

#### Command output:
```
{"time":"2024-07-08T09:49:16.986166355+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia1694863704/cmd/DST/cmd/archive/6","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-07-08T09:49:16.986407847+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}
{"time":"2024-07-08T09:49:16.990262893+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/1/1/0"}
{"time":"2024-07-08T09:49:16.991074369+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T09:49:16.995203236+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T09:49:16.995338614+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/2/2/0"}
{"time":"2024-07-08T09:49:16.995332248+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T09:49:17.001805761+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T09:49:17.003684987+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":315},"msg":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-08T09:49:17.642833892+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/4/3/1"}
{"time":"2024-07-08T09:49:17.648103894+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/5/4/1"}
{"time":"2024-07-08T09:49:17.655840799+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.087","MinifyRatio":"1.000","original":47218,"compressed":4096,"minified":47218,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/ohter_files/CT_Krpаlek__Lukач_2_977869_20231212143738.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Sakísuw,_Frmšú_Tuesday_W50_2023_12_12T143738.xml"}
{"time":"2024-07-08T09:49:17.656290498+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.127","original":47218,"compressed":0,"minified":6007,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/ohter_files/CT_Krpаlek__Lukач_2_977869_20231212143738.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Sakísuw,_Frmšú_Tuesday_W50_2023_12_12T143738.xml"}
{"time":"2024-07-08T09:49:17.842419239+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:18.565538688+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.077","MinifyRatio":"1.000","original":11960034,"compressed":921600,"minified":11960034,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/ohter_files/RD_05-09_ČRo_Sever_-_Wed__04_03_2020_2_1607196_20200304234759.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:18.691593734+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:18.691735242+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/6/5/1"}
{"time":"2024-07-08T09:49:19.214094751+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.033","MinifyRatio":"0.323","original":11960034,"compressed":393216,"minified":3865086,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/ohter_files/RD_05-09_ČRo_Sever_-_Wed__04_03_2020_2_1607196_20200304234759.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia1694863704/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia1694863704/cmd/SRC/cmd/archive/hello_invalid2.txt"],"filename is not valid OpenMedia file":["/tmp/openmedia1694863704/cmd/SRC/cmd/archive/ohter_files/hello_invalid.xml","/tmp/openmedia1694863704/cmd/SRC/cmd/archive/ohter_files/hello_invalid2.xml"]}
{"time":"2024-07-08T09:49:19.214205229+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"10/10/6/4"}

result: archive_create
{"time":"2024-07-08T09:49:19.214228394+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"6/6/5/1"}
{"time":"2024-07-08T09:49:19.214246155+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.075","MinifyRatio":"1.000","original":25653580,"compressed":1912832,"minified":25653580,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia1694863704/cmd/DST/cmd/archive/6"}
{"time":"2024-07-08T09:49:19.214275211+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.030","MinifyRatio":"0.324","original":25653580,"compressed":778240,"minified":8312126,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia1694863704/cmd/DST/cmd/archive/6"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```

### 7. archive: do not recurse the source folder
Command input:

running from source:
```
go run main.go -v=0 archive -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/7
```

running compiled:
```
./openmedia -v=0 archive -sdir=/tmp/openmedia1694863704/cmd/SRC/cmd/archive -odir=/tmp/openmedia1694863704/cmd/DST/cmd/archive/7
```

#### Command output:
```
{"time":"2024-07-08T09:49:19.679471359+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia1694863704/cmd/DST/cmd/archive/7","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-07-08T09:49:19.679699224+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-07-08T09:49:19.683362787+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T09:49:19.684544508+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-07-08T09:49:19.689748222+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T09:49:19.689982116+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-07-08T09:49:19.693295704+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-07-08T09:49:19.697234118+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-07-08T09:49:19.702317694+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":315},"msg":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-07-08T09:49:20.277192118+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-08T09:49:20.422486356+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:20.974510571+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia1694863704/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia1694863704/cmd/SRC/cmd/archive/hello_invalid2.txt"]}
{"time":"2024-07-08T09:49:20.974598821+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}

result: archive_create
{"time":"2024-07-08T09:49:20.974614398+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-07-08T09:49:20.974626064+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.072","MinifyRatio":"1.000","original":13646328,"compressed":987136,"minified":13646328,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia1694863704/cmd/DST/cmd/archive/7"}
{"time":"2024-07-08T09:49:20.974639972+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.028","MinifyRatio":"0.325","original":13646328,"compressed":385024,"minified":4441033,"filePathSource":"/tmp/openmedia1694863704/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia1694863704/cmd/DST/cmd/archive/7"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia1694863704/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}
```
### Run summary

--- PASS: TestRunCommand_Archive (7.94s)
    --- PASS: TestRunCommand_Archive/1 (0.40s)
    --- PASS: TestRunCommand_Archive/2 (0.41s)
    --- PASS: TestRunCommand_Archive/3 (1.68s)
    --- PASS: TestRunCommand_Archive/4 (0.48s)
    --- PASS: TestRunCommand_Archive/5 (0.43s)
    --- PASS: TestRunCommand_Archive/6 (2.77s)
    --- PASS: TestRunCommand_Archive/7 (1.76s)
PASS
ok  	github/czech-radio/openmedia/cmd	7.952s



## Command: extractArchive
### 1. extractArchive: help
Command input:

running from source:
```
go run main.go extractArchive -h -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/1
```

running compiled:
```
./openmedia extractArchive -h -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/1
```

#### Command output:
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
Command input:

running from source:
```
go run main.go extractArchive -ofname=production -exsn=production_all -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/2
```

running compiled:
```
./openmedia extractArchive -ofname=production -exsn=production_all -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/2
```

#### Command output:
```
{"time":"2024-07-08T09:49:22.177702627+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":null,"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip|Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","38","5082"],"PartPrefixCode":11,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/2","OutputFileName":"production","ExtractorsName":"production_all","ExtractorsCode":"production_all","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"","FilterRecords":false,"ValidatorFileName":""}}
{"time":"2024-07-08T09:49:22.178138822+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-08T09:49:22.178159341+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":1}
{"time":"2024-07-08T09:49:22.178169348+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":2}
{"time":"2024-07-08T09:49:22.178276833+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T09:49:22.178299289+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:22.486648504+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":166}
{"time":"2024-07-08T09:49:22.486692082+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:22.765613257+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":147}
{"time":"2024-07-08T09:49:22.765774526+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T09:49:22.765792279+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T09:49:22.767144985+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T09:49:22.767158652+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T09:49:22.767260221+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":312}
{"time":"2024-07-08T09:49:22.767295028+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":312}
{"time":"2024-07-08T09:49:22.771975116+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/2/production_production_all_base_wh.csv","bytesCount":1125}
{"time":"2024-07-08T09:49:22.772065329+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/2/production_production_all_base_wh.csv","bytesCount":157961}
{"time":"2024-07-08T09:49:22.772109174+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/2/production_production_all_base_woh.csv","bytesCount":478}
{"time":"2024-07-08T09:49:22.77219075+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/2/production_production_all_base_woh.csv","bytesCount":157961}
{"time":"2024-07-08T09:49:22.77220706+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDataset","file":"outputs.go","line":27},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-08T09:49:22.772216441+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":49},"msg":"filter file not specified"}
```

### 3. extractArchive: extract all contacts from minified rundowns
Command input:

running from source:
```
go run main.go extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/3
```

running compiled:
```
./openmedia extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/3
```

#### Command output:
```
{"time":"2024-07-08T09:49:23.218518289+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":null,"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/3","OutputFileName":"production","ExtractorsName":"production_contacts","ExtractorsCode":"production_contacts","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"","FilterRecords":false,"ValidatorFileName":""}}
{"time":"2024-07-08T09:49:23.21891811+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-08T09:49:23.218935171+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":1}
{"time":"2024-07-08T09:49:23.218947915+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":2}
{"time":"2024-07-08T09:49:23.21904831+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T09:49:23.219070183+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:23.504932742+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":166}
{"time":"2024-07-08T09:49:23.504976501+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:23.788807714+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":147}
{"time":"2024-07-08T09:49:23.788959657+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T09:49:23.788976803+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T09:49:23.790344591+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T09:49:23.790361353+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T09:49:23.790453454+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-08T09:49:23.790465127+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-08T09:49:23.791003691+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-08T09:49:23.791017074+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-08T09:49:23.791639994+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/3/production_production_contacts_base_wh.csv","bytesCount":1051}
{"time":"2024-07-08T09:49:23.791685173+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/3/production_production_contacts_base_wh.csv","bytesCount":30295}
{"time":"2024-07-08T09:49:23.791725797+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/3/production_production_contacts_base_woh.csv","bytesCount":444}
{"time":"2024-07-08T09:49:23.791771874+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/3/production_production_contacts_base_woh.csv","bytesCount":30295}
{"time":"2024-07-08T09:49:23.791786077+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDataset","file":"outputs.go","line":27},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-08T09:49:23.791801261+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":49},"msg":"filter file not specified"}
```

### 4. extractArchive: extract all story parts from minified rundowns, extract only specified radios
Command input:

running from source:
```
go run main.go extractArchive -ofname=production -exsn=production_all -frn=Olomouc -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/4
```

running compiled:
```
./openmedia extractArchive -ofname=production -exsn=production_all -frn=Olomouc -fdf=2020-03-04 -fdt=2020-03-05 -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/4
```

#### Command output:
```
{"time":"2024-07-08T09:49:24.20713445+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":{"Olomouc":true},"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip|Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","38","5082"],"PartPrefixCode":11,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/4","OutputFileName":"production","ExtractorsName":"production_all","ExtractorsCode":"production_all","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"Olomouc","FilterRecords":false,"ValidatorFileName":""}}
{"time":"2024-07-08T09:49:24.207633868+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":0,"matched":true}
{"time":"2024-07-08T09:49:24.207656557+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":0}
{"time":"2024-07-08T09:49:24.207697357+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":0}
{"time":"2024-07-08T09:49:24.207851043+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T09:49:24.207929992+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-08T09:49:24.207941418+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-08T09:49:24.2079524+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-08T09:49:24.207961261+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-08T09:49:24.207969695+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":0,"filtered":0}
{"time":"2024-07-08T09:49:24.207978711+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":0}
{"time":"2024-07-08T09:49:24.20834449+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/4/production_production_all_base_wh.csv","bytesCount":1125}
{"time":"2024-07-08T09:49:24.208378013+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/4/production_production_all_base_wh.csv","bytesCount":0}
{"time":"2024-07-08T09:49:24.208419386+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/4/production_production_all_base_woh.csv","bytesCount":478}
{"time":"2024-07-08T09:49:24.208444737+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/4/production_production_all_base_woh.csv","bytesCount":0}
{"time":"2024-07-08T09:49:24.208459643+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputValidatedDataset","file":"outputs.go","line":27},"msg":"validation_warning","msg":"validation receipe file not specified"}
{"time":"2024-07-08T09:49:24.208469748+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":49},"msg":"filter file not specified"}
```

### 5. extractArchive: extract all story parts from minified rundowns and validate
Command input:

running from source:
```
go run main.go extractArchive -ofname=production -exsn=production_all -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5
```

running compiled:
```
./openmedia extractArchive -ofname=production -exsn=production_all -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5
```

#### Command output:
```
{"time":"2024-07-08T09:49:24.740653676+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":null,"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip|Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Audioclip","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","38","5082"],"PartPrefixCode":11,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5","OutputFileName":"production","ExtractorsName":"production_all","ExtractorsCode":"production_all","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"","FilterRecords":false,"ValidatorFileName":"../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx"}}
{"time":"2024-07-08T09:49:24.741098789+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-08T09:49:24.741115594+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":1}
{"time":"2024-07-08T09:49:24.741126261+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":2}
{"time":"2024-07-08T09:49:24.741233336+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T09:49:24.741256162+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:25.102952385+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":166}
{"time":"2024-07-08T09:49:25.102995323+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:25.382838338+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":147}
{"time":"2024-07-08T09:49:25.382988428+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T09:49:25.38300042+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T09:49:25.384308151+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T09:49:25.384320535+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T09:49:25.384409781+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":312}
{"time":"2024-07-08T09:49:25.384428306+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":312}
{"time":"2024-07-08T09:49:25.388567699+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5/production_production_all_base_wh.csv","bytesCount":1125}
{"time":"2024-07-08T09:49:25.388651632+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5/production_production_all_base_wh.csv","bytesCount":157961}
{"time":"2024-07-08T09:49:25.388681843+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5/production_production_all_base_woh.csv","bytesCount":478}
{"time":"2024-07-08T09:49:25.38874267+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5/production_production_all_base_woh.csv","bytesCount":157961}
{"time":"2024-07-08T09:49:25.448776155+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5/production_production_all_base_validated_wh.csv","bytesCount":1125}
{"time":"2024-07-08T09:49:25.448908853+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5/production_production_all_base_validated_wh.csv","bytesCount":150024}
{"time":"2024-07-08T09:49:25.448952945+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5/production_production_all_base_validated_woh.csv","bytesCount":478}
{"time":"2024-07-08T09:49:25.449045373+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5/production_production_all_base_validated_woh.csv","bytesCount":150024}
{"time":"2024-07-08T09:49:25.449399374+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":332},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/5/production_base_validated_log.csv","bytesCount":4948}
{"time":"2024-07-08T09:49:25.449416724+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).OutputFilteredDataset","file":"outputs.go","line":49},"msg":"filter file not specified"}
```

### 6. extractArchive: extract all story parts from minified rundowns, validate and use filter oposition
Command input:

running from source:
```
go run main.go extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6
```

running compiled:
```
./openmedia extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6
```

#### Command output:
```
{"time":"2024-07-08T09:49:25.833087937+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":null,"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6","OutputFileName":"production","ExtractorsName":"production_contacts","ExtractorsCode":"production_contacts","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"","FilterRecords":false,"ValidatorFileName":"../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx"}}
{"time":"2024-07-08T09:49:25.833511658+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-08T09:49:25.833530405+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":1}
{"time":"2024-07-08T09:49:25.833543569+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":2}
{"time":"2024-07-08T09:49:25.833656471+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T09:49:25.83367822+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:26.121108257+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":166}
{"time":"2024-07-08T09:49:26.121152901+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:26.384638401+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":147}
{"time":"2024-07-08T09:49:26.384775165+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T09:49:26.384787161+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T09:49:26.386126194+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T09:49:26.386139188+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T09:49:26.386219699+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-08T09:49:26.386228693+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-08T09:49:26.386759201+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-08T09:49:26.386768535+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-08T09:49:26.387362185+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_wh.csv","bytesCount":1051}
{"time":"2024-07-08T09:49:26.387398205+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_wh.csv","bytesCount":30295}
{"time":"2024-07-08T09:49:26.3874258+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_woh.csv","bytesCount":444}
{"time":"2024-07-08T09:49:26.387454681+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_woh.csv","bytesCount":30295}
{"time":"2024-07-08T09:49:26.446029293+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_wh.csv","bytesCount":1051}
{"time":"2024-07-08T09:49:26.446083924+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_wh.csv","bytesCount":25388}
{"time":"2024-07-08T09:49:26.446114439+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_woh.csv","bytesCount":444}
{"time":"2024-07-08T09:49:26.44614066+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_woh.csv","bytesCount":25388}
{"time":"2024-07-08T09:49:26.446262774+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":332},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_base_validated_log.csv","bytesCount":4256}
{"time":"2024-07-08T09:49:26.617168888+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":1212}
{"time":"2024-07-08T09:49:26.617224987+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":27808}
{"time":"2024-07-08T09:49:26.617256223+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":502}
{"time":"2024-07-08T09:49:26.617282551+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/6/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":27808}
```

### 7. extractArchive: extract all story parts from minified rundowns, validate and use filter eurovolby
Command input:

running from source:
```
go run main.go extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7
```

running compiled:
```
./openmedia extractArchive -ofname=production -exsn=production_contacts -fdf=2020-03-04 -fdt=2020-03-05 -valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx -frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx -sdir=/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive -odir=/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7
```

#### Command output:
```
{"time":"2024-07-08T09:49:27.039311684+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.ParseConfigOptions","file":"extract_archive.go","line":88},"msg":"effective subcommand config","config":{"RadioNames":null,"DateRange":["2020-03-04T00:00:00.000000001Z","2020-03-05T00:00:00Z"],"IsoWeeks":null,"Months":null,"WeekDays":null,"Extractors":[{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":["C-index"],"PartPrefixCode":19,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["5081","1000","8"],"PartPrefixCode":1,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Radio Rundown/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":0,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Hourly Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1000"],"PartPrefixCode":2,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":3,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Sub Rundown","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","1004","1005","1002","321"],"PartPrefixCode":4,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":5,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/*Radio Story","ObjectAttrsNames":["ObjectID"],"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["8","12","1004","1005","1002","321","5079","16","5016","5","6","5070","5071","5072"],"PartPrefixCode":6,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/<OM_RECORD>","ObjectAttrsNames":["RecordID","RecordID"],"FieldsPath":"","FieldIDs":["5001","8","38","5082","421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":7,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":["TemplateName","ObjectID"],"FieldsPath":"","FieldIDs":null,"PartPrefixCode":15,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":1,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"/Contact Item|Contact Bin","ObjectAttrsNames":null,"FieldsPath":"/OM_HEADER/OM_FIELD","FieldIDs":["421","422","5015","424","423","5016","5088","5087","5068"],"PartPrefixCode":13,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":true,"FieldIDsMap":null},{"OrExt":null,"ObjectPath":"","ObjectAttrsNames":null,"FieldsPath":"","FieldIDs":null,"PartPrefixCode":17,"FieldsPrefix":"","KeepInputRow":false,"ResultNodeGoUpLevels":0,"KeepWhenZeroSubnodes":false,"FieldIDsMap":null}],"ComputeUniqueRows":false,"PrintHeader":false,"CSVdelim":"\t","SourceDirectory":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive","SourceDirectoryType":"MINIFIED.zip","WorkerType":1,"OutputDirectory":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7","OutputFileName":"production","ExtractorsName":"production_contacts","ExtractorsCode":"production_contacts","AddRecordsNumbers":true,"FilterDateFrom":"2020-03-04T00:00:00+01:00","FilterDateTo":"2020-03-05T00:00:00+01:00","FilterRadioName":"","FilterRecords":false,"ValidatorFileName":"../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx"}}
{"time":"2024-07-08T09:49:27.039671231+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.PackageMap","file":"archive_package.go","line":217},"msg":"filenames in all packages","count":2,"matched":true}
{"time":"2024-07-08T09:49:27.03968839+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":128},"msg":"packages matched","count":1}
{"time":"2024-07-08T09:49:27.03969751+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderMap","file":"archive_folder.go","line":129},"msg":"files inside packages matched","count":2}
{"time":"2024-07-08T09:49:27.039798163+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":139},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip"}
{"time":"2024-07-08T09:49:27.039820331+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:27.327899925+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":166}
{"time":"2024-07-08T09:49:27.327941046+02:00","level":"WARN","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":154},"msg":"proccessing package","package":"/tmp/openmedia2409342867/cmd/SRC/cmd/extractArchive/Rundowns/2020_W10_MINIFIED.zip","file":"RD_05-09_ŠKs_Wfbvy_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-07-08T09:49:27.610157575+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*ArchiveFolder).FolderExtract","file":"archive_folder.go","line":164},"msg":"extracted lines","count":147}
{"time":"2024-07-08T09:49:27.610300286+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T09:49:27.610312691+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T09:49:27.611710831+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":313}
{"time":"2024-07-08T09:49:27.61172734+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":313}
{"time":"2024-07-08T09:49:27.611807519+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":313,"filtered":278}
{"time":"2024-07-08T09:49:27.611817372+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":278}
{"time":"2024-07-08T09:49:27.612338399+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":472},"msg":"rows count before deletion","parsed":278,"filtered":44}
{"time":"2024-07-08T09:49:27.612350167+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).DeleteNonMatchingRows","file":"filter.go","line":486},"msg":"rows count after deletion","count_after":44}
{"time":"2024-07-08T09:49:27.612943683+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_wh.csv","bytesCount":1051}
{"time":"2024-07-08T09:49:27.612981452+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_wh.csv","bytesCount":30295}
{"time":"2024-07-08T09:49:27.613010466+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_woh.csv","bytesCount":444}
{"time":"2024-07-08T09:49:27.613043972+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_woh.csv","bytesCount":30295}
{"time":"2024-07-08T09:49:27.66378833+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_wh.csv","bytesCount":1051}
{"time":"2024-07-08T09:49:27.663836732+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_wh.csv","bytesCount":25388}
{"time":"2024-07-08T09:49:27.663865255+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_woh.csv","bytesCount":444}
{"time":"2024-07-08T09:49:27.663891263+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_woh.csv","bytesCount":25388}
{"time":"2024-07-08T09:49:27.664034506+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).ValidationLogWrite","file":"validate.go","line":332},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_base_validated_log.csv","bytesCount":4256}
{"time":"2024-07-08T09:49:27.756985288+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":1179}
{"time":"2024-07-08T09:49:27.757041333+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_wh.csv","bytesCount":27544}
{"time":"2024-07-08T09:49:27.757073979+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVheaderWrite","file":"table_csv.go","line":188},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":497}
{"time":"2024-07-08T09:49:27.757109461+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/extract.(*Extractor).CSVtableWrite","file":"table_csv.go","line":170},"msg":"written bytes to file","fileName":"/tmp/openmedia2409342867/cmd/DST/cmd/extractArchive/7/production_production_contacts_base_validated_filtered_woh.csv","bytesCount":27544}
```
### Run summary

--- PASS: TestRunCommand_ExtractArchive (6.37s)
    --- PASS: TestRunCommand_ExtractArchive/1 (0.38s)
    --- PASS: TestRunCommand_ExtractArchive/2 (1.00s)
    --- PASS: TestRunCommand_ExtractArchive/3 (1.02s)
    --- PASS: TestRunCommand_ExtractArchive/4 (0.42s)
    --- PASS: TestRunCommand_ExtractArchive/5 (1.24s)
    --- PASS: TestRunCommand_ExtractArchive/6 (1.17s)
    --- PASS: TestRunCommand_ExtractArchive/7 (1.14s)
PASS
ok  	github/czech-radio/openmedia/cmd	6.372s
