package helper

import (
	"flag"
	"fmt"
	"log/slog"
	"reflect"
	"strconv"
	"time"
)

// TODO: print usage for all commands at once using run main.go and all subcmds with help
// TODO: implement validate flag value by function (check alloved values)

type RootCfg struct {
	Version     bool
	Verbose     int
	DryRun      bool
	LogType     string
	DebugConfig bool
}

var CommandRoot = CommandConfig{
	Opts: []FlagOption{
		{FlagDescription{"version", "V", "false", "bool",
			"Version of program."},
			nil, nil},

		{FlagDescription{"verbose", "v", "0", "int",
			"Level of verbosity."},
			[]any{-4, -3, -2, -1, 0, 1, 2, 3, 4}, nil},

		{FlagDescription{"dryRun", "dr", "false", "bool",
			"Dry run, useful for tests."},
			nil, nil},

		{FlagDescription{"logType", "logt", "json", "string",
			"Type of logs formating."},
			[]any{"json", "plain"}, nil},

		{FlagDescription{"debugConfig", "dc", "false", "bool",
			"Debug/print flag values"},
			[]any{"json", "plain"}, nil},
	},
}

type CommandConfig struct {
	OptsMap map[string][5]interface{}
	Opts    []FlagOption
	Subs    Subcommands
	Values  interface{}
}

type FlagOption struct {
	FlagDescription
	AllovedValues []any
	FuncMatch     func(any) bool `json:"-"`
}

type FlagDescription struct {
	LongFlag   string
	ShortFlag  string
	Default    string
	Type       string
	Descripton string
}

type OptsDec struct {
	Long, Short, Default interface{}
	Alloved              interface{}
}

type Subcommands map[string]func()

func (cc *CommandConfig) Init() {
	cc.DeclareFlags()
	rcfg := &RootCfg{}
	flag.Parse()
	err := cc.ParseFlags(rcfg)
	if err != nil {
		panic(err)
	}
	SetLogLevel(strconv.Itoa(rcfg.Verbose), rcfg.LogType)
	cc.Values = rcfg
	if flag.NArg() < 1 {
		fmt.Println("version to do")
		// VersionInfoPrint()
		return
	}
	slog.Info("root config", "config", cc.Values)
}

func (cc *CommandConfig) RunRoot() {
	subCmdName := flag.Arg(0)
	if subCmdName == "" {
		return
	}
	subCmdFunc, ok := cc.Subs[subCmdName]
	if !ok {
		panic(fmt.Errorf("unknown subcommand: %s", subCmdName))
	}
	subCmdFunc()
}

func (cc *CommandConfig) RunSub(intf interface{}) {
	subcmd := flag.Arg(0)
	slog.Info("subcommand called", "subcommand", subcmd)
	FlagsUsage = fmt.Sprintf("subcommand: %s\n", subcmd)
	cc.DeclareFlags()
	err := flag.CommandLine.Parse(flag.Args()[1:])
	if err != nil {
		panic(err)
	}
	err = cc.ParseFlags(intf)
	if err != nil {
		panic(err)
	}
}

func (cc *CommandConfig) AddSub(subName string, subF func()) {
	if cc.Subs == nil {
		cc.Subs = make(Subcommands)
	}
	cc.Subs[subName] = subF
}

func (cc *CommandConfig) AddOption(
	long, short, defValue, typeValue, descr string,
	alloved []any, funcM func(any) bool) {
	opt := FlagOption{
		FlagDescription: FlagDescription{
			long, short, defValue, typeValue, descr,
		}, AllovedValues: alloved, FuncMatch: funcM}
	cc.Opts = append(cc.Opts, opt)
}

func (opt FlagDescription) Error(err error) {
	if err == nil {
		return
	}
	errMsgFormat := "cannot parse flag %s as type %s, err %v"
	errMsg := fmt.Errorf(errMsgFormat, opt.LongFlag, opt.Type, err)
	panic(errMsg)
}

func CheckAllovedValues(flagName string, inp any, alloved interface{}) {
	var match bool
	if alloved == nil {
		return
	}
	switch t := alloved.(type) {
	case []interface{}:
		if len(t) == 0 {
			return
		}
		for _, i := range alloved.([]interface{}) {
			if inp == i {
				match = true
				break
			}
		}
	case func(any) bool:
	default:
		err := fmt.Errorf("unknow type of alloved definition: %v", t)
		panic(err)
	}
	if !match {
		err := fmt.Errorf("flag '-%s=%v' not alloved, alloved values: %v",
			flagName, inp, alloved)
		panic(err)
	}
}

func (opt *FlagOption) DeclareUsage() {
	fd := opt.FlagDescription
	if opt.AllovedValues == nil {
		format := "-%s, -%s\n\t%s\n\n"
		FlagsUsage += fmt.Sprintf(format, fd.ShortFlag, fd.LongFlag, fd.Descripton)
	} else {
		format := "-%s, -%s\n\t%s\n\t%v\n\n"
		FlagsUsage += fmt.Sprintf(format, fd.ShortFlag, fd.LongFlag, fd.Descripton, opt.AllovedValues)
	}
}

func (cc *CommandConfig) DeclareFlags() {
	slog.Debug("declaring flags")
	cc.OptsMap = make(map[string][5]interface{})
	for i := range cc.Opts {
		slog.Debug("declaring flag", "opt", cc.Opts[i])
		res := cc.Opts[i].DeclareFlag()
		name := cc.Opts[i].LongFlag
		name = FirstLetterToUppercase(name)
		cc.OptsMap[name] = res
	}
	flag.Usage = Usage
}

func (opt *FlagOption) DeclareFlag() [5]interface{} {
	var def, long, short interface{}
	opt.DeclareUsage()
	switch opt.FlagDescription.Type {
	case "bool":
		b, err := strconv.ParseBool(opt.Default)
		opt.Error(err)
		def = &b
		long = flag.Bool(opt.LongFlag, b, opt.Descripton)
		short = flag.Bool(opt.ShortFlag, b, opt.Descripton)
	case "int":
		b, err := strconv.Atoi(opt.Default)
		opt.Error(err)
		def = &b
		long = flag.Int(opt.LongFlag, b, opt.Descripton)
		short = flag.Int(opt.ShortFlag, b, opt.Descripton)
	case "string":
		def = opt.Default
		long = flag.String(opt.LongFlag, opt.Default, opt.Descripton)
		short = flag.String(opt.ShortFlag, opt.Default, opt.Descripton)
	case "date":
		def = opt.Default
		long = flag.String(opt.LongFlag, opt.Default, opt.Descripton)
		short = flag.String(opt.ShortFlag, opt.Default, opt.Descripton)
	default:
		err := fmt.Errorf("unknow flag type: %s", opt.FlagDescription.Type)
		opt.Error(err)
	}
	return [5]interface{}{def, long, short, "", opt.AllovedValues}
}

func (cc *CommandConfig) ParseFlags(iface interface{}) error {
	slog.Info("parsing flags")
	vof := reflect.ValueOf(iface)
	if vof.Kind() != reflect.Ptr || vof.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("Invalid input: not a pointer to a struct")
	}
	vofe := vof.Elem()
	n := vofe.NumField()
	for i := 0; i < n; i++ {
		field := vofe.Type().Field(i)
		optName := FirstLetterToUppercase(field.Name)
		vals, ok := cc.OptsMap[optName]
		if !ok {
			slog.Debug("flag not defined for struct field", "field", optName)
			continue
		}
		slog.Debug("parsing flag", "name", field.Name, "type", field.Type.Name())
		def := vals[0]
		long := vals[1]
		short := vals[2]
		alloved := vals[4]
		switch field.Type.Name() {
		case "bool":
			vals := []bool{*long.(*bool), *short.(*bool), *def.(*bool)}
			res := GetBoolValuePriority(vals...)
			vofe.Field(i).SetBool(res)
		case "int":
			vals := []int{*long.(*int), *short.(*int), *def.(*int)}
			res := GetIntValuePriority(vals...)
			CheckAllovedValues(optName, res, alloved)
			vofe.Field(i).SetInt(int64(res))
		case "string":
			vals := []string{*long.(*string), *short.(*string), def.(string)}
			res := GetStringValuePriority(vals...)
			CheckAllovedValues(optName, res, alloved)
			vofe.Field(i).SetString(res)
		case "Time":
			vals := []string{*long.(*string), *short.(*string), def.(string)}
			res := GetStringValuePriority(vals...)
			if res == "" {
				continue
			}
			date, err := ParseStringDate(res, time.Local)
			if err != nil {
				panic(err)
			}
			vofe.Field(i).Set(reflect.ValueOf(date))
		default:
			panic(fmt.Errorf("flag type not implemented: %s", field.Type.Name()))
		}
	}
	return nil
}

// OPTION GETERS
// GetBoolValuePriority return value according to priority. Priority is given in desceding. Last value is default value.
func GetBoolValuePriority(vals ...bool) bool {
	count := len(vals) - 1
	res := vals[count]
	for i := count - 1; i >= 0; i-- {
		res = XOR(res, vals[i])
	}
	return res
}

func GetIntValuePriority(vals ...int) int {
	count := len(vals) - 1
	def := vals[count]
	res := def
	for i := count - 1; i >= 0; i-- {
		if vals[i] != def {
			res = vals[i]
		}
	}
	return res
}

func GetStringValuePriority(vals ...string) string {
	count := len(vals) - 1
	def := vals[count]
	res := def
	for i := count - 1; i >= 0; i-- {
		if vals[i] != def {
			res = vals[i]
		}
	}
	return res
}
