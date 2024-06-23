=== RUN   TestRunCommand_Archive
=== RUN   TestRunCommand_Archive/1
## command: archive
### 1. archive: help
COMMAND_INPUT:
	go run main.go archive -h -sdir=/tmp/openmedia734137417/cmd/SRC/cmd/archive -odir=/tmp/openmedia734137417/cmd/DST/cmd/archive/1
	openmedia archive -h -sdir=/tmp/openmedia734137417/cmd/SRC/cmd/archive -odir=/tmp/openmedia734137417/cmd/DST/cmd/archive/1
#### COMMAND_OUTPUT_START:
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



=== RUN   TestRunCommand_Archive/2
### 2. archive: debug config
COMMAND_INPUT:
	go run main.go -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia734137417/cmd/SRC/cmd/archive -odir=/tmp/openmedia734137417/cmd/DST/cmd/archive/2
	openmedia -v=0 -dc archive -ifc -R -sdir=/tmp/openmedia734137417/cmd/SRC/cmd/archive -odir=/tmp/openmedia734137417/cmd/DST/cmd/archive/2
#### COMMAND_OUTPUT_START:
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
Archive config: {SourceDirectory:/tmp/openmedia734137417/cmd/SRC/cmd/archive OutputDirectory:/tmp/openmedia734137417/cmd/DST/cmd/archive/2 CompressionType:zip InvalidFilenameContinue:true InvalidFileContinue:false InvalidFileRename:false ProcessedFileRename:false ProcessedFileDelete:false PreserveFoldersInArchive:false RecurseSourceDirectory:true}

=== RUN   TestRunCommand_Archive/3
### 3. archive: run exit on src file error
COMMAND_INPUT:
	go run main.go -v=0 archive -ifc -sdir=/tmp/openmedia734137417/cmd/SRC/cmd/archive -odir=/tmp/openmedia734137417/cmd/DST/cmd/archive/3
	openmedia -v=0 archive -ifc -sdir=/tmp/openmedia734137417/cmd/SRC/cmd/archive -odir=/tmp/openmedia734137417/cmd/DST/cmd/archive/3
#### COMMAND_OUTPUT_START:
{"time":"2024-06-23T11:52:37.169232846+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia734137417/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia734137417/cmd/DST/cmd/archive/3","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-06-23T11:52:37.169404209+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":339},"msg":"Archive filenames validation result","valid_ratio":"4/6"}
{"time":"2024-06-23T11:52:37.173216975+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":504},"msg":"Contacts/2023_W49_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47182,"compressed":0,"minified":47182,"filePathSource":"/tmp/openmedia734137417/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_ORIGINAL.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-06-23T11:52:37.176040955+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":504},"msg":"Contacts/2023_W50_ORIGINAL.zip","ArhiveRatio":"0.000","MinifyRatio":"1.000","original":47170,"compressed":0,"minified":47170,"filePathSource":"/tmp/openmedia734137417/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_ORIGINAL.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-06-23T11:52:37.1802098+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":504},"msg":"Contacts/2023_W49_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.119","original":47182,"compressed":0,"minified":5596,"filePathSource":"/tmp/openmedia734137417/cmd/SRC/cmd/archive/CT_AAA-domаcб_obecnа_osoba_2_10359730_20231208144502.xml","filePathDestination":"Contacts/2023_W49_MINIFIED.zip/CT_ELV-iokíkč_mniykš_lqaht_Friday_W49_2023_12_08T144502.xml"}
{"time":"2024-06-23T11:52:37.186505077+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).WorkerLogInfo","file":"archive.go","line":504},"msg":"Contacts/2023_W50_MINIFIED.zip","ArhiveRatio":"0.000","MinifyRatio":"0.111","original":47170,"compressed":0,"minified":5219,"filePathSource":"/tmp/openmedia734137417/cmd/SRC/cmd/archive/CT_Hrab╪tovа__Jana_2_985786_20231213023749.xml","filePathDestination":"Contacts/2023_W50_MINIFIED.zip/CT_Cvmqďsbzř,_Plhh_Wednesday_W50_2023_12_13T023749.xml"}
{"time":"2024-06-23T11:52:37.18803307+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/helper.ErrorsCodeMap.ExitWithCode","file":"errors.go","line":112},"msg":"/tmp/openmedia734137417/cmd/SRC/cmd/archive/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431_bad.xml xml: encoding \"utf-16\" declared but Decoder.CharsetReader is nil"}
exit status 12

--- PASS: TestRunCommand_Archive (1.27s)
    --- PASS: TestRunCommand_Archive/1 (0.42s)
    --- PASS: TestRunCommand_Archive/2 (0.37s)
    --- PASS: TestRunCommand_Archive/3 (0.47s)
PASS
ok  	github/czech-radio/openmedia/cmd	1.283s
