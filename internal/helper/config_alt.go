package helper

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

type FlagOption struct {
	FlagDescription
	AllovedValues []any
	FuncMatch     func(any) error
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

type CommandConfig struct {
	Opts    []FlagOption
	OptsMap map[string][5]interface{}
}

type RootCfg struct {
	Version bool
	Verbose int
}

var CommandRoot = CommandConfig{
	Opts: []FlagOption{
		{FlagDescription{"version", "V", "false", "bool", "version of program"},
			nil, nil},
		{FlagDescription{"verbose", "v", "2", "int", "level of verbosity"},
			[]any{1, 2, 3, 4}, nil},
	},
}

func (opt FlagDescription) Error(err error) {
	if err == nil {
		return
	}
	errMsgFormat := "cannot parse flag %s as type %s, err %v"
	errMsg := fmt.Errorf(errMsgFormat, opt.LongFlag, opt.Type, err)
	panic(errMsg)
}

func (cc *CommandConfig) DeclareFlags() {
	cc.OptsMap = make(map[string][5]interface{})
	for i := range cc.Opts {
		res := cc.Opts[i].DeclareFlag()
		name := cc.Opts[i].LongFlag
		cc.OptsMap[name] = res
	}
	flag.Parse()
}

func CheckValAlloved(flagName string, inp any, alloved interface{}) {
	var match bool
	for _, i := range alloved.([]interface{}) {
		if inp == i {
			match = true
			break
		}
	}
	if !match {
		err := fmt.Errorf("flag '-%s=%v' not alloved, alloved values: %v",
			flagName, inp, alloved)
		panic(err)
	}
}

func (opt *FlagOption) DeclareFlag() [5]interface{} {
	var def, long, short interface{}
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
	default:
		err := fmt.Errorf("unknow flag type")
		opt.Error(err)
	}
	return [5]interface{}{def, long, short, "", opt.AllovedValues}
}

func (cc *CommandConfig) ParseFlags(iface interface{}) error {
	vof := reflect.ValueOf(iface)
	if vof.Kind() != reflect.Ptr || vof.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("Invalid input: not a pointer to a struct")
	}
	vofe := vof.Elem()
	n := vofe.NumField()
	for i := 0; i < n; i++ {
		field := vofe.Type().Field(i)
		optName := FirstLetterToLowercase(field.Name)
		vals, ok := cc.OptsMap[optName]
		if !ok {
			continue
		}
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
			CheckValAlloved(optName, res, alloved)
			vofe.Field(i).SetInt(int64(res))
		default:
			panic("flag type not implemented")
		}
	}
	return nil
}

// GetBoolValuePriority return value according to priority. Priority is given in desceding. Last value is default value.
func GetBoolValuePriority(vals ...bool) bool {
	count := len(vals) - 1
	res := vals[count]
	for i := count - 1; i >= 0; i-- {
		res = XOR(res, vals[i])
	}
	return res
}

func GetIntValuePriority(
	vals ...int) int {
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
