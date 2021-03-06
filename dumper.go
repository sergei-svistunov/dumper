package dumper

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"sort"
)

func Dump(w io.Writer, variable interface{}) {
	(&dumper{
		ptrs: make(map[uintptr]bool),
	}).printReflect(w, reflect.ValueOf(variable))
}

// DumpToString dumps interface to string.
func DumpToString(i interface{}) string {
	buf := new(bytes.Buffer)
	Dump(buf, i)
	return buf.String()
}

type reflectValues []reflect.Value

func (v reflectValues) Len() int {
	return len(v)
}

// Not all reflectValues can be sorted, but for tests it's enaugh
func (v reflectValues) Less(i, j int) bool {
	if v[i].Kind() != v[j].Kind() {
		return false
	}

	switch v[i].Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v[i].Int() < v[j].Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v[i].Uint() < v[j].Uint()
	case reflect.Float32, reflect.Float64:
		return v[i].Float() < v[j].Float()
	case reflect.Ptr:
		return v[i].UnsafeAddr() < v[j].UnsafeAddr()
	case reflect.String:
		return v[i].String() < v[j].String()
	}

	return false
}

func (v reflectValues) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

type dumper struct {
	ptrs map[uintptr]bool
}

func (d *dumper) printReflect(w io.Writer, reflectValue reflect.Value) {
	switch reflectValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(w, "%s(%d)", reflectValue.Type().Name(), reflectValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(w, "%s(%d)", reflectValue.Type().Name(), reflectValue.Uint())
	case reflect.Bool:
		fmt.Fprintf(w, "%s(%t)", reflectValue.Type().Name(), reflectValue.Bool())
	case reflect.String:
		fmt.Fprintf(w, "%s(%q)", reflectValue.Type().Name(), reflectValue.String())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(w, "%s(%f)", reflectValue.Type().Name(), reflectValue.Float())
	case reflect.Ptr:
		if reflectValue.IsNil() {
			fmt.Fprintf(w, "%s(nil)", reflectValue.Type().Name())
		} else {
			addr := reflectValue.Pointer()
			if _, exists := d.ptrs[addr]; exists {
				fmt.Fprintf(w, "%s(&(0x%x))", reflectValue.Elem().Type().Name(), addr)
			} else {
				d.ptrs[addr] = true
				fmt.Fprintf(w, "&(0x%x)", addr)
				d.printReflect(w, reflectValue.Elem())
			}
		}
	case reflect.Invalid:
		fmt.Fprintf(w, "<INVALID>")
	case reflect.Interface:
		d.printReflect(w, reflectValue.Elem())
	case reflect.Struct:
		fmt.Fprintf(w, "%s{", reflectValue.Type().Name())
		for i := 0; i < reflectValue.NumField(); i++ {
			if i != 0 {
				fmt.Fprint(w, ",")
			}
			field := reflectValue.Field(i)

			fmt.Fprintf(w, "%s:", reflectValue.Type().Field(i).Name)
			d.printReflect(w, field)
		}
		fmt.Fprint(w, "}")
	case reflect.Array, reflect.Slice:
		if reflectValue.Type().Name() == "" {
			fmt.Fprintf(w, "%s{", reflectValue.Type().String())
		} else {
			fmt.Fprintf(w, "%s/*%s*/{", reflectValue.Type().String(), reflectValue.Type().Kind().String())
		}
		for i := 0; i < reflectValue.Len(); i++ {
			if i != 0 {
				fmt.Fprint(w, ",")
			}
			d.printReflect(w, reflectValue.Index(i))
		}
		fmt.Fprint(w, "}")
	case reflect.Map:
		if reflectValue.Type().Name() == "" {
			fmt.Fprintf(w, "%s{", reflectValue.Type().String())
		} else {
			fmt.Fprintf(w, "%s/*%s*/{", reflectValue.Type().String(), reflectValue.Type().Kind().String())
		}
		keys := reflectValues(reflectValue.MapKeys())
		sort.Sort(keys)
		for i, key := range keys {
			if i != 0 {
				fmt.Fprint(w, ",")
			}
			d.printReflect(w, key)
			fmt.Fprint(w, ":")
			d.printReflect(w, reflectValue.MapIndex(key))
		}
		fmt.Fprint(w, "}")
	default:
		fmt.Fprintf(w, "%s(<???>)", reflectValue.Kind().String())
	}
}
