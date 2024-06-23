
## Command: root
### 1. root: help
Command input:
	go run main.go -h
	openmedia -h
#### Command output:
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



### 2. root: print version
Command input:
	go run main.go -V
	openmedia -V
#### Command output:
{ProgramName: Version:1.0.0 GitTag: GitCommit: BuildTime:}

### 3. root: print config
Command input:
	go run main.go -dc
	openmedia -dc
#### Command output:
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}

### 4. root: print version and config
Command input:
	go run main.go -V -dc
	openmedia -V -dc
#### Command output:
{ProgramName: Version:1.0.0 GitTag: GitCommit: BuildTime:}
Root config: &{Usage:false GeneralHelp:false Version:true Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}

### Run summary
--- PASS: TestRunCommand_Root (1.87s)
    --- PASS: TestRunCommand_Root/1 (0.48s)
    --- PASS: TestRunCommand_Root/2 (0.47s)
    --- PASS: TestRunCommand_Root/3 (0.51s)
    --- PASS: TestRunCommand_Root/4 (0.39s)
PASS
ok  	github/czech-radio/openmedia/cmd	1.874s

## Command: archive
### 1. archive: help
Command input:
	go run main.go archive -h -sdir=/tmp/openmedia2985602263/cmd/SRC/cmd/archive -odir=/tmp/openmedia2985602263/cmd/DST/cmd/archive/1
	openmedia archive -h -sdir=/tmp/openmedia2985602263/cmd/SRC/cmd/archive -odir=/tmp/openmedia2985602263/cmd/DST/cmd/archive/1
#### Command output:
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



### 2. archive: debug config
Command input:
	go run main.go -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia2985602263/cmd/SRC/cmd/archive -odir=/tmp/openmedia2985602263/cmd/DST/cmd/archive/2
	openmedia -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia2985602263/cmd/SRC/cmd/archive -odir=/tmp/openmedia2985602263/cmd/DST/cmd/archive/2
#### Command output:
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
Archive config: {SourceDirectory:/tmp/openmedia2985602263/cmd/SRC/cmd/archive OutputDirectory:/tmp/openmedia2985602263/cmd/DST/cmd/archive/2 CompressionType:zip InvalidFilenameContinue:true InvalidFileContinue:false InvalidFileRename:false ProcessedFileRename:false ProcessedFileDelete:false PreserveFoldersInArchive:false RecurseSourceDirectory:true}

### 3. archive: run exit on src file error
Command input:
	go run main.go -v=0 archive -ifc -sdir=/tmp/openmedia2985602263/cmd/SRC/cmd/archive -odir=/tmp/openmedia2985602263/cmd/DST/cmd/archive/3
	openmedia -v=0 archive -ifc -sdir=/tmp/openmedia2985602263/cmd/SRC/cmd/archive -odir=/tmp/openmedia2985602263/cmd/DST/cmd/archive/3
#### Command output:
{"time":"2024-06-23T13:50:15.335037809+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2985602263/cmd/DST/cmd/archive/3","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-06-23T13:50:15.335229216+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-06-23T13:50:15.338317253+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-06-23T13:50:15.338402352+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-06-23T13:50:15.342193761+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-06-23T13:50:15.342699986+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-06-23T13:50:15.342369238+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-06-23T13:50:15.350853042+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-06-23T13:50:15.350880142+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/helper.ErrorsCodeMap.ExitWithCode","file":"errors.go","line":112},"msg":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
exit status 12

### 4. archive: run do not exit on src file error
Command input:
	go run main.go -v=0 archive -sdir=/tmp/openmedia2985602263/cmd/SRC/cmd/archive -odir=/tmp/openmedia2985602263/cmd/DST/cmd/archive/4
	openmedia -v=0 archive -sdir=/tmp/openmedia2985602263/cmd/SRC/cmd/archive -odir=/tmp/openmedia2985602263/cmd/DST/cmd/archive/4
#### Command output:
{"time":"2024-06-23T13:50:15.844235568+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2985602263/cmd/DST/cmd/archive/4","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-06-23T13:50:15.844605531+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-06-23T13:50:15.849808407+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-06-23T13:50:15.850881124+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-06-23T13:50:15.861702165+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-06-23T13:50:15.861975346+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-06-23T13:50:15.866484094+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-06-23T13:50:15.869725099+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-06-23T13:50:15.887660938+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":315},"msg":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-06-23T13:50:16.394222769+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-06-23T13:50:16.530020731+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-06-23T13:50:17.05625871+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia2985602263/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia2985602263/cmd/SRC/cmd/archive/hello_invalid2.txt"]}
{"time":"2024-06-23T13:50:17.056333043+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}

result: archive_create
{"time":"2024-06-23T13:50:17.05635359+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-06-23T13:50:17.056371149+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.072","MinifyRatio":"1.000","original":13646328,"compressed":987136,"minified":13646328,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia2985602263/cmd/DST/cmd/archive/4"}
{"time":"2024-06-23T13:50:17.056389879+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.028","MinifyRatio":"0.325","original":13646328,"compressed":385024,"minified":4441033,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia2985602263/cmd/DST/cmd/archive/4"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia2985602263/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}

### 5. archive: dry run, no exit on src file error
Command input:
	go run main.go -dr -v=0 archive -sdir=/tmp/openmedia2985602263/cmd/SRC/cmd/archive -odir=/tmp/openmedia2985602263/cmd/DST/cmd/archive/5
	openmedia -dr -v=0 archive -sdir=/tmp/openmedia2985602263/cmd/SRC/cmd/archive -odir=/tmp/openmedia2985602263/cmd/DST/cmd/archive/5
#### Command output:
{"time":"2024-06-23T13:50:17.505283124+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia2985602263/cmd/DST/cmd/archive/5","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-06-23T13:50:17.50541298+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":66},"msg":"dry run activated","output_path":"/tmp/openmedia_archive70816923"}
{"time":"2024-06-23T13:50:17.505498647+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}
{"time":"2024-06-23T13:50:17.509156279+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/1/1/0"}
{"time":"2024-06-23T13:50:17.510170131+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-06-23T13:50:17.513620263+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-06-23T13:50:17.514029434+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/2/2/0"}
{"time":"2024-06-23T13:50:17.514469591+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-06-23T13:50:17.520221176+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-06-23T13:50:17.522432584+02:00","level":"ERROR","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":315},"msg":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
{"time":"2024-06-23T13:50:18.046179293+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-06-23T13:50:18.215635715+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_ORIGINAL.zip","ArhiveRatio":"0.073","MinifyRatio":"1.000","original":13551976,"compressed":987136,"minified":13551976,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_ORIGINAL.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}
{"time":"2024-06-23T13:50:18.721088448+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"Rundowns/2020_W10_MINIFIED.zip","ArhiveRatio":"0.028","MinifyRatio":"0.327","original":13551976,"compressed":385024,"minified":4430218,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive/RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml","filePathDestination":"Rundowns/2020_W10_MINIFIED.zip/RD_12-19_ŘRf_Tcunyso_Wednesday_W10_2020_03_04.xml"}

RESULTS:

result: filenames_validation
errors: {"file does not have xml extension":["/tmp/openmedia2985602263/cmd/SRC/cmd/archive/hello_invalid.txt","/tmp/openmedia2985602263/cmd/SRC/cmd/archive/hello_invalid2.txt"]}
{"time":"2024-06-23T13:50:18.721167628+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"filenames_validation","All/Processed/Success/Failure":"6/6/4/2"}

result: archive_create
{"time":"2024-06-23T13:50:18.721188172+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.ProcessStats.LogProcessStats","file":"xml.go","line":158},"msg":"process stat","name":"archive_create","All/Processed/Success/Failure":"4/4/3/1"}
{"time":"2024-06-23T13:50:18.721207589+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_ORIGINAL","ArhiveRatio":"0.072","MinifyRatio":"1.000","original":13646328,"compressed":987136,"minified":13646328,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia_archive70816923"}
{"time":"2024-06-23T13:50:18.721225949+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":473},"msg":"GLOBAL_MINIFY","ArhiveRatio":"0.028","MinifyRatio":"0.325","original":13646328,"compressed":385024,"minified":4441033,"filePathSource":"/tmp/openmedia2985602263/cmd/SRC/cmd/archive","filePathDestination":"/tmp/openmedia_archive70816923"}
errors: {"xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil":["/tmp/openmedia2985602263/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml"]}

### Run summary
--- PASS: TestRunCommand_Archive (4.65s)
    --- PASS: TestRunCommand_Archive/1 (0.41s)
    --- PASS: TestRunCommand_Archive/2 (0.42s)
    --- PASS: TestRunCommand_Archive/3 (0.43s)
    --- PASS: TestRunCommand_Archive/4 (1.71s)
    --- PASS: TestRunCommand_Archive/5 (1.67s)
PASS
ok  	github/czech-radio/openmedia/cmd	4.655s
