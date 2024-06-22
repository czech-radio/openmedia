=== RUN   TestRunCommand_Root
=== RUN   TestRunCommand_Root/1

## command: root
### 1. root: help
COMMAND_INPUT:
	go run main.go -h
	openmedia -h
#### COMMAND_OUTPUT_START:
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



=== RUN   TestRunCommand_Root/2
### 2. root: print version
COMMAND_INPUT:
	go run main.go -V
	openmedia -V
#### COMMAND_OUTPUT_START:
{ProgramName: Version:1.0.0 GitTag: GitCommit: BuildTime:}

=== RUN   TestRunCommand_Root/3
### 3. root: print config
COMMAND_INPUT:
	go run main.go -dc
	openmedia -dc
#### COMMAND_OUTPUT_START:
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}

=== RUN   TestRunCommand_Root/4
### 4. root: print version and config
COMMAND_INPUT:
	go run main.go -V -dc
	openmedia -V -dc
#### COMMAND_OUTPUT_START:
{ProgramName: Version:1.0.0 GitTag: GitCommit: BuildTime:}
Root config: &{Usage:false GeneralHelp:false Version:true Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}

=== RUN   TestRunCommand_Root/5
### 5. root: test log [err]
COMMAND_INPUT:
	go run main.go -logt=plain -logts -v=6
	openmedia -logt=plain -logts -v=6
#### COMMAND_OUTPUT_START:
time=2024-06-22T14:26:46.122+02:00 level=ERROR source=sloger.go:53 msg="test error"

=== RUN   TestRunCommand_Root/6
### 6. root: test log [err,warn]
COMMAND_INPUT:
	go run main.go -logt=plain -logts -v=4
	openmedia -logt=plain -logts -v=4
#### COMMAND_OUTPUT_START:
time=2024-06-22T14:26:46.598+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-06-22T14:26:46.598+02:00 level=ERROR source=sloger.go:53 msg="test error"

=== RUN   TestRunCommand_Root/7
### 7. root: test log [err,warn,info]
COMMAND_INPUT:
	go run main.go -logt=plain -logts -v=0
	openmedia -logt=plain -logts -v=0
#### COMMAND_OUTPUT_START:
time=2024-06-22T14:26:47.015+02:00 level=INFO source=sloger.go:51 msg="test info"
time=2024-06-22T14:26:47.015+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-06-22T14:26:47.015+02:00 level=ERROR source=sloger.go:53 msg="test error"

=== RUN   TestRunCommand_Root/8
### 8. root: terst log [err,warn,info,debug]
COMMAND_INPUT:
	go run main.go -logt=plain -logts -v=-4
	openmedia -logt=plain -logts -v=-4
#### COMMAND_OUTPUT_START:
time=2024-06-22T14:26:47.414+02:00 level=DEBUG source=config_mapany.go:309 msg="subcommand added" subname=archive
time=2024-06-22T14:26:47.414+02:00 level=DEBUG source=config_mapany.go:309 msg="subcommand added" subname=extractArchive
time=2024-06-22T14:26:47.414+02:00 level=DEBUG source=sloger.go:50 msg="test debug"
time=2024-06-22T14:26:47.414+02:00 level=INFO source=sloger.go:51 msg="test info"
time=2024-06-22T14:26:47.414+02:00 level=WARN source=sloger.go:52 msg="test warn"
time=2024-06-22T14:26:47.414+02:00 level=ERROR source=sloger.go:53 msg="test error"

=== RUN   TestRunCommand_Root/9
### 9. root: test log json
COMMAND_INPUT:
	go run main.go -logt=json -logts -v=-4
	openmedia -logt=json -logts -v=-4
#### COMMAND_OUTPUT_START:
{"time":"2024-06-22T14:26:47.798906545+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":309},"msg":"subcommand added","subname":"archive"}
{"time":"2024-06-22T14:26:47.798975356+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/configure.(*CommanderConfig).AddSub","file":"config_mapany.go","line":309},"msg":"subcommand added","subname":"extractArchive"}
{"time":"2024-06-22T14:26:47.79899156+02:00","level":"DEBUG","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":50},"msg":"test debug"}
{"time":"2024-06-22T14:26:47.799003518+02:00","level":"INFO","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":51},"msg":"test info"}
{"time":"2024-06-22T14:26:47.799013487+02:00","level":"WARN","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":52},"msg":"test warn"}
{"time":"2024-06-22T14:26:47.799024613+02:00","level":"ERROR","source":{"function":"github.com/triopium/go_utils/pkg/logging.LoggingOutputTest","file":"sloger.go","line":53},"msg":"test error"}

### Run summary
--- PASS: TestRunCommand_Root (3.91s)
    --- PASS: TestRunCommand_Root/1 (0.52s)
    --- PASS: TestRunCommand_Root/2 (0.41s)
    --- PASS: TestRunCommand_Root/3 (0.47s)
    --- PASS: TestRunCommand_Root/4 (0.40s)
    --- PASS: TestRunCommand_Root/5 (0.42s)
    --- PASS: TestRunCommand_Root/6 (0.48s)
    --- PASS: TestRunCommand_Root/7 (0.42s)
    --- PASS: TestRunCommand_Root/8 (0.40s)
    --- PASS: TestRunCommand_Root/9 (0.38s)
PASS
ok  	github/czech-radio/openmedia/cmd	3.923s
=== RUN   TestRunCommand_Archive
=== RUN   TestRunCommand_Archive/0

## command: archive
### 1. archive: help
COMMAND_INPUT:
	go run main.go archive -h -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/1
	openmedia archive -h -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/1
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



=== RUN   TestRunCommand_Archive/1
### 2. archive: debug config
COMMAND_INPUT:
	go run main.go -dc -v=0 archive -ifc -R -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/2
	openmedia -dc -v=0 archive -ifc -R -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/2
#### COMMAND_OUTPUT_START:
Root config: &{Usage:false GeneralHelp:false Version:false Verbose:0 DryRun:false LogType:json LogTest:false DebugConfig:true}
Archive config: {SourceDirectory:/tmp/openmedia906999227/cmd/SRC/cmd/archive OutputDirectory:/tmp/openmedia906999227/cmd/DST/cmd/archive/2 CompressionType:zip InvalidFilenameContinue:true InvalidFileContinue:false InvalidFileRename:false ProcessedFileRename:false ProcessedFileDelete:false PreserveFoldersInArchive:false RecurseSourceDirectory:true}

=== RUN   TestRunCommand_Archive/2
### 3. archive: run exit on src file error
COMMAND_INPUT:
	go run main.go -v=0 archive -ifc -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/3
	openmedia -v=0 archive -ifc -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/3
#### COMMAND_OUTPUT_START:
{"time":"2024-06-22T14:26:49.471810024+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia906999227/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia906999227/cmd/DST/cmd/archive/3","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-06-22T14:26:49.472014488+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":340},"msg":"Archive filenames validation result","valid_ratio":"0/1"}

=== RUN   TestRunCommand_Archive/3
### 4. archive: run do not exit on src file error
COMMAND_INPUT:
	go run main.go -v=0 archive -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/4
	openmedia -v=0 archive -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/4
#### COMMAND_OUTPUT_START:
{"time":"2024-06-22T14:26:49.941988196+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia906999227/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia906999227/cmd/DST/cmd/archive/4","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-06-22T14:26:49.942202197+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":340},"msg":"Archive filenames validation result","valid_ratio":"0/1"}

=== RUN   TestRunCommand_Archive/4
### 5. archive: dry run, no exit on src file error
COMMAND_INPUT:
	go run main.go -dr -v=0 archive -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/5
	openmedia -dr -v=0 archive -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/5
#### COMMAND_OUTPUT_START:
{"time":"2024-06-22T14:26:50.471140183+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia906999227/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia906999227/cmd/DST/cmd/archive/5","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":true,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":false}}
{"time":"2024-06-22T14:26:50.471270137+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":66},"msg":"dry run activated","output_path":"/tmp/openmedia_archive3516975605"}
{"time":"2024-06-22T14:26:50.471366449+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":340},"msg":"Archive filenames validation result","valid_ratio":"0/1"}

=== RUN   TestRunCommand_Archive/5
### 6. archive: recurse source directory
COMMAND_INPUT:
	go run main.go -v=0 archive -ifc -R -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/6
	openmedia -v=0 archive -ifc -R -sdir=/tmp/openmedia906999227/cmd/SRC/cmd/archive -odir=/tmp/openmedia906999227/cmd/DST/cmd/archive/6
#### COMMAND_OUTPUT_START:
{"time":"2024-06-22T14:26:50.94314165+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/cmd.GlobalConfig.RunCommandArchive","file":"archive.go","line":63},"msg":"effective config","config":{"SourceDirectory":"/tmp/openmedia906999227/cmd/SRC/cmd/archive","OutputDirectory":"/tmp/openmedia906999227/cmd/DST/cmd/archive/6","CompressionType":"zip","InvalidFilenameContinue":true,"InvalidFileContinue":false,"InvalidFileRename":false,"ProcessedFileRename":false,"ProcessedFileDelete":false,"PreserveFoldersInArchive":false,"RecurseSourceDirectory":true}}
{"time":"2024-06-22T14:26:50.943361272+02:00","level":"INFO","source":{"function":"github/czech-radio/openmedia/internal/archive.(*Archive).Folder","file":"archive.go","line":340},"msg":"Archive filenames validation result","valid_ratio":"1/1"}

### Run summary
--- PASS: TestRunCommand_Archive (2.73s)
    --- PASS: TestRunCommand_Archive/0 (0.47s)
    --- PASS: TestRunCommand_Archive/1 (0.40s)
    --- PASS: TestRunCommand_Archive/2 (0.38s)
    --- PASS: TestRunCommand_Archive/3 (0.47s)
    --- PASS: TestRunCommand_Archive/4 (0.53s)
    --- PASS: TestRunCommand_Archive/5 (0.47s)
PASS
ok  	github/czech-radio/openmedia/cmd	2.737s
