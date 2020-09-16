package csv

import (
	"fmt"
	"reflect"

	"logger"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

var Comma string = ";"
var Booltrue string = "wahr"
var Boolfalse string = "falsch"

func FormatCsv(v interface{}) (s string) {
	val := reflect.Indirect(reflect.ValueOf(v))
	return FormatCsvReflect(val)
}

func FormatCsvReflect(v reflect.Value) (s string) {
	p := message.NewPrinter(language.German)
	switch v.Kind() {
	case reflect.String:
		s = fmt.Sprintf("\"%s\"", v)
	case reflect.Float32, reflect.Float64:
		s = p.Sprint(number.Decimal(v.Float(), number.MaxFractionDigits(2)))
	case reflect.Int:
		s = p.Sprint(number.Decimal(v.Int(), number.MaxFractionDigits(0)))
	case reflect.Bool:
		if v.Bool() {
			s = Booltrue
		} else {
			s = Boolfalse
		}
	default:
		logger.Warn("Unknown type: '%T'\n", v)
	}
	return s
}

func GetCsvName(val reflect.StructField) string {
	if tag := val.Tag.Get("csv"); tag != "" {
		return tag
	} else {
		return val.Name
	}
}

func GenerericToCsv(a interface{}) string {
	var result string
	val := reflect.Indirect(reflect.ValueOf(a))
	for i := 0; i < val.Type().NumField(); i++ {
		result += Comma + FormatCsvReflect(val.Field(i))
	}
	return result
}

func GenericDescribe(a interface{}) string {
	result := ""
	val := reflect.Indirect(reflect.ValueOf(a))
	for i := 0; i < val.Type().NumField(); i++ {
		if i > 0 {
			result += Comma
		}
		result += FormatCsv(GetCsvName(val.Type().Field(i)))
	}
	return result
}

func GenericStructToString(a interface{}) string {
	var result string
	val := reflect.Indirect(reflect.ValueOf(a))
	for i := 0; i < val.Type().NumField(); i++ {
		result += fmt.Sprintf("%s: %s\n", GetCsvName(val.Type().Field(i)), FormatCsvReflect(val.Field(i)))
	}
	return result
}

func GenericSliceToString(fields []interface{}) (result string) {
	for _, f := range fields[:len(fields)-1] {
		result += FormatCsv(f) + Comma
	}
	result += FormatCsv(fields[len(fields)-1])
	return result
}
