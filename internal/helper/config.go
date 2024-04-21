package helper

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var flagsUsage = "Usage:\n"

// Usage called when help command invoked
func Usage() {
	fmt.Println(flagsUsage)
}

type FlagsMap map[string]map[string]interface{}

func SetField(rv reflect.Value, value any) error {
	if !rv.IsValid() {
		return fmt.Errorf("not a valid field")
	}

	// Check if the field is settable
	if !rv.CanSet() {
		return fmt.Errorf("cannot set value for field")
	}
	rv.Set(reflect.ValueOf(value))
	return nil
}

func GetStringValueByPriority(
	longFlagValue, shortFlagValue, envValue, defaultValue string) string {
	res := defaultValue
	if longFlagValue != defaultValue {
		res = longFlagValue
	}
	if shortFlagValue != defaultValue {
		res = shortFlagValue
	}
	return res
}

func GetBoolValueByPriority(
	longFlagValue, shortFlagValue, envValue, defaultValue bool) bool {
	res := XOR(defaultValue, shortFlagValue)
	res = XOR(res, longFlagValue)
	return res
}

func GetIntValueByPriority(
	longFlagValue, shortFlagValue, envValue, defaultValue int) int {
	res := defaultValue
	if longFlagValue != defaultValue {
		res = longFlagValue
	}
	if shortFlagValue != defaultValue {
		res = shortFlagValue
	}
	return res
}

func DeclareFlags(config interface{}) (FlagsMap, error) {
	flags := make(FlagsMap)
	v := reflect.ValueOf(config)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return flags, fmt.Errorf("Invalid input: not a pointer to a struct")
	}

	velem := v.Elem()
	n := velem.NumField()
	for i := 0; i < n; i++ {
		field := velem.Type().Field(i)
		tagValue := field.Tag.Get("cmd")
		if tagValue == "" {
			continue
		}
		fieldType := field.Type
		cmdOpts := strings.Split(tagValue, "; ")
		flagsUsage += fmt.Sprintf("-%s, --%s\n\t%s\n", cmdOpts[1], cmdOpts[0], cmdOpts[3])
		flagMap := make(map[string]interface{})
		flags[field.Name] = flagMap
		switch fieldType.Kind() {
		case reflect.String:
			flagMap["long"] = flag.String(cmdOpts[0], cmdOpts[2], cmdOpts[3])
			flagMap["short"] = flag.String(cmdOpts[1], cmdOpts[2], cmdOpts[3])
			flagMap["default"] = cmdOpts[2]
			flagMap["field"] = velem.FieldByName(field.Name)
		case reflect.Bool:
			b, err := strconv.ParseBool(cmdOpts[2])
			if err != nil {
				return flags, err
			}
			flagMap["long"] = flag.Bool(cmdOpts[0], false, cmdOpts[3])
			flagMap["short"] = flag.Bool(cmdOpts[1], false, cmdOpts[3])
			flagMap["default"] = b
			flagMap["field"] = velem.FieldByName(field.Name)
		case reflect.Int:
			v, err := strconv.Atoi(cmdOpts[2])
			if err != nil {
				return flags, err
			}
			flagMap["long"] = flag.Int(cmdOpts[0], v, cmdOpts[3])
			flagMap["short"] = flag.Int(cmdOpts[1], v, cmdOpts[3])
			flagMap["default"] = v
			flagMap["field"] = velem.FieldByName(field.Name)
		}
		flags[field.Name] = flagMap
	}
	flag.Usage = Usage
	return flags, nil
}

func (fm FlagsMap) ParseFlags() error {
	for _, k := range fm {
		rfv := k["field"].(reflect.Value)
		short := k["short"]
		long := k["long"]
		def := k["default"]
		switch rfv.Kind() {
		case reflect.String:
			val := GetStringValueByPriority(
				*long.(*string), *short.(*string), "", def.(string))
			err := SetField(rfv, val)
			if err != nil {
				return err
			}
		case reflect.Bool:
			val := GetBoolValueByPriority(
				*long.(*bool), *short.(*bool), false, def.(bool))
			err := SetField(rfv, val)
			if err != nil {
				return err
			}
		case reflect.Int:
			val := GetIntValueByPriority(
				*long.(*int), *short.(*int), 0, def.(int))
			err := SetField(rfv, val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SetupRootFlags(config interface{}) {
	flags, err := DeclareFlags(config)
	if err != nil {
		panic(err)
	}
	flag.Parse()
	err = flags.ParseFlags()
	if err != nil {
		panic(err)
	}
}

func SetupSubFlags(config interface{}) {
	flags, err := DeclareFlags(config)
	if err != nil {
		panic(err)
	}
	err = flag.CommandLine.Parse(flag.Args()[1:])
	if err != nil {
		panic(err)
	}
	err = flags.ParseFlags()
	if err != nil {
		panic(err)
	}
}

// CopyFields copy struct fields values from one struct to another struct
func CopyFields(a interface{}, b interface{}) {
	va := reflect.ValueOf(a)
	va_elem := va.Elem()
	n := va_elem.NumField()
	for i := 0; i < n; i++ {
		field := va_elem.Type().Field(i)
		fieldValue := va_elem.Field(i).Interface()
		vb := reflect.ValueOf(b)
		vbfield := vb.Elem().FieldByName(field.Name)
		if !vbfield.IsValid() {
			continue
		}
		err := SetField(vbfield, fieldValue)
		if err != nil {
			panic(err)
		}
	}
}
