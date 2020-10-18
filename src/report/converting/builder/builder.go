package builder

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/VladLeb13/report-maker-lib/src/datalib"
)

func getValue(val interface{}) (out []map[string]string) {
	types := reflect.TypeOf(val)
	values := reflect.ValueOf(val)
	for i := 0; i < types.NumField(); i++ {
		key := types.Field(i).Tag.Get("name")
		v := values.Field(i)
		var value string
		//fmt.Println(name, reflect.ValueOf(value), "type of:", value.Type())
		switch v.Kind() {
		case reflect.String:
			value = fmt.Sprint(reflect.ValueOf(v))
		case reflect.Uint16:
			value = strconv.FormatUint(uint64(v.Interface().(uint16)), 10)
		case reflect.Uint64:
			value = strconv.FormatUint(uint64(v.Interface().(uint64)), 10)
		case reflect.Uint32:
			value = strconv.FormatUint(uint64(v.Interface().(uint32)), 10)
		case reflect.Struct:
			switch v.Type().String() {
			case "datalib.BIOS":

				tmp := getValue(v.Interface().(datalib.BIOS))
				for _, el := range tmp {
					out = append(out, el)
				}

				continue
			default:
				log.Println("other struct detected", v.Type())
			}
		case reflect.Slice:
			if !v.IsZero() {
				line := fmt.Sprint(v)
				value = line[1 : len(line)-1]
			} else {
				value = " - "
			}
		case reflect.Bool:
			line := v.Interface().(bool)
			if line {
				value = "🗹"
			} else {
				value = "☐"
			}
		default:
			fmt.Println("не строка", v, "type of:", v.Type())
		}
		m := make(map[string]string)
		m[key] = value
		out = append(out, m)
	}
	return
}
