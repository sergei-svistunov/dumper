package dumper_test

import (
	"bytes"
	"reflect"
	"regexp"
	"testing"

	"github.com/sergei-svistunov/dumper"
)

var addrRe, _ = regexp.Compile("\\&\\(0x[0-9a-f]+\\)")

type testStruct struct {
	f1 int
	f2 string
	f3 *testStruct
	f4 interface{}
}

type ArrayAlias [3]int
type SliceAlias []int
type MapAlias map[int]int

func sPtr(s string) *string { return &s }

func check(t *testing.T, v interface{}, expected, description string) {
	buf := &bytes.Buffer{}
	dumper.Dump(buf, v)

	gotten := addrRe.ReplaceAllString(buf.String(), "&(0xADDR)")

	if gotten != expected {
		t.Errorf("%s: expected \"%s\", gotten \"%s\"", description, expected, gotten)
	}
}

func checkBeautify(t *testing.T, v interface{}, expected *dumper.BeautifyNode, description string) {
	buf := &bytes.Buffer{}

	dumper.Dump(buf, v)

	gotten, err := dumper.GetBeautifyTree(buf.Bytes())
	if err != nil {
		t.Errorf("%s: parse error: \"%s\"", description, err.Error())
		return
	}

	fixPtrs(gotten)

	if !reflect.DeepEqual(expected, gotten) {
		t.Errorf("%s: expected \n%s\ngotten \n%s\n", description, expected, gotten)
	}
}

func fixPtrs(n *dumper.BeautifyNode) {
	if n == nil {
		return
	}

	if n.Ptr != nil {
		n.Ptr = sPtr("0xADDR")
	}

	for _, sn := range n.StructValues {
		fixPtrs(sn.Value)
	}
}

func TestDumpInt(t *testing.T) {
	check(t, int(123), "int(123)", "Int")
	check(t, int8(123), "int8(123)", "Int8")
	check(t, int16(123), "int16(123)", "Int16")
	check(t, int32(123), "int32(123)", "Int32")
	check(t, int64(123), "int64(123)", "Int64")

	checkBeautify(t, int(-123), &dumper.BeautifyNode{Type: "int", Value: sPtr("-123")}, "Int beautify")
}

func TestDumpUint(t *testing.T) {
	check(t, uint(123), "uint(123)", "Uint")
	check(t, uint8(123), "uint8(123)", "Uint8")
	check(t, uint16(123), "uint16(123)", "Uint16")
	check(t, uint32(123), "uint32(123)", "Uint32")
	check(t, uint64(123), "uint64(123)", "Uint64")

	checkBeautify(t, uint(123), &dumper.BeautifyNode{Type: "uint", Value: sPtr("123")}, "Uint beautify")
}

func TestDumpBool(t *testing.T) {
	check(t, true, "bool(true)", "Bool true")
	check(t, false, "bool(false)", "Bool false")

	checkBeautify(t, true, &dumper.BeautifyNode{Type: "bool", Value: sPtr("true")}, "Bool true beautify")
	checkBeautify(t, false, &dumper.BeautifyNode{Type: "bool", Value: sPtr("false")}, "Bool false beautify")
}

func TestDumpFloat(t *testing.T) {
	check(t, float32(10.5), "float32(10.500000)", "Float32")
	check(t, float64(10.5), "float64(10.500000)", "Float64")

	checkBeautify(t, float32(10.5), &dumper.BeautifyNode{Type: "float32", Value: sPtr("10.500000")}, "Float32 beautify")
	checkBeautify(t, float64(-10.5), &dumper.BeautifyNode{Type: "float64", Value: sPtr("-10.500000")}, "Float64 negative beautify")
}

func TestDumpString(t *testing.T) {
	check(t, "Test\n,\t,\"", `string("Test\n,\t,\"")`, "String")

	checkBeautify(t, "Test\n,\t,\"", &dumper.BeautifyNode{Type: "string", Value: sPtr(`"Test\n,\t,\""`)}, "String beautify")
}

func TestDumpPtr(t *testing.T) {
	intVar := int(123)
	check(t, &intVar, "&(0xADDR)int(123)", "Ptr to int")

	checkBeautify(t, &intVar, &dumper.BeautifyNode{Ptr: sPtr("0xADDR"), Type: "int", Value: sPtr("123")}, "Ptr to int beautify")

	intVarPtr := &intVar
	check(t, &intVarPtr, "&(0xADDR)&(0xADDR)int(123)", "Ptr to ptr to int")

	checkBeautify(t, &intVarPtr, &dumper.BeautifyNode{Ptr: sPtr("0xADDR"), Type: "int", Value: sPtr("123")}, "Ptr to ptr to int beautify")
}

func TestDumpInterface(t *testing.T) {
	var interfaceVar interface{}

	check(t, interfaceVar, "<INVALID>", "Empty interface")
	checkBeautify(t, interfaceVar, &dumper.BeautifyNode{Type: "<INVALID>"}, "Empty interface beautify")

	check(t, &interfaceVar, "&(0xADDR)<INVALID>", "Ptr to empty interface")
	checkBeautify(t, &interfaceVar, &dumper.BeautifyNode{Ptr: sPtr("0xADDR"), Type: "<INVALID>"}, "Ptr to empty interface beautify")

	interfaceVar = int(123)

	check(t, interfaceVar, "int(123)", "Int in interface")
	checkBeautify(t, interfaceVar, &dumper.BeautifyNode{Type: "int", Value: sPtr("123")}, "Int in interface beautify")

	check(t, &interfaceVar, "&(0xADDR)int(123)", "Ptr to Int in interface")
	checkBeautify(t, &interfaceVar, &dumper.BeautifyNode{Ptr: sPtr("0xADDR"), Type: "int", Value: sPtr("123")}, "Ptr to Int in interface beautify")
}

func TestDumpStruct(t *testing.T) {
	s := testStruct{
		f1: 1,
		f3: &testStruct{
			f1: 2,
			f3: &testStruct{
				f1: 3,
			},
		},
		f4: 100,
	}

	check(t, s, `testStruct{f1:int(1),f2:string(""),f3:&(0xADDR)testStruct{f1:int(2),f2:string(""),f3:&(0xADDR)testStruct{f1:int(3),f2:string(""),f3:(nil),f4:<INVALID>},f4:<INVALID>},f4:int(100)}`, "Struct")
	checkBeautify(t, s, &dumper.BeautifyNode{
		Type: "testStruct",
		StructValues: []*dumper.StructKV{
			&dumper.StructKV{"f1", &dumper.BeautifyNode{Type: "int", Value: sPtr("1")}},
			&dumper.StructKV{"f2", &dumper.BeautifyNode{Type: "string", Value: sPtr(`""`)}},
			&dumper.StructKV{"f3", &dumper.BeautifyNode{
				Ptr:  sPtr("0xADDR"),
				Type: "testStruct",
				StructValues: []*dumper.StructKV{
					&dumper.StructKV{"f1", &dumper.BeautifyNode{Type: "int", Value: sPtr("2")}},
					&dumper.StructKV{"f2", &dumper.BeautifyNode{Type: "string", Value: sPtr(`""`)}},
					&dumper.StructKV{"f3", &dumper.BeautifyNode{
						Ptr:  sPtr("0xADDR"),
						Type: "testStruct",
						StructValues: []*dumper.StructKV{
							&dumper.StructKV{"f1", &dumper.BeautifyNode{Type: "int", Value: sPtr("3")}},
							&dumper.StructKV{"f2", &dumper.BeautifyNode{Type: "string", Value: sPtr(`""`)}},
							&dumper.StructKV{"f3", nil},
							&dumper.StructKV{"f4", &dumper.BeautifyNode{Type: "<INVALID>"}},
						},
					},
					},
					&dumper.StructKV{"f4", &dumper.BeautifyNode{Type: "<INVALID>"}},
				},
			},
			},
			&dumper.StructKV{"f4", &dumper.BeautifyNode{Type: "int", Value: sPtr("100")}},
		},
	}, "Struct beautify")

	s.f3.f3.f3 = &s
	check(t, s, `testStruct{f1:int(1),f2:string(""),f3:&(0xADDR)testStruct{f1:int(2),f2:string(""),f3:&(0xADDR)testStruct{f1:int(3),f2:string(""),f3:&(0xADDR)testStruct{f1:int(1),f2:string(""),f3:testStruct(&(0xADDR)),f4:int(100)},f4:<INVALID>},f4:<INVALID>},f4:int(100)}`, "Struct with cycle")
	checkBeautify(t, &s, &dumper.BeautifyNode{
		Ptr:  sPtr("0xADDR"),
		Type: "testStruct",
		StructValues: []*dumper.StructKV{
			&dumper.StructKV{"f1", &dumper.BeautifyNode{Type: "int", Value: sPtr("1")}},
			&dumper.StructKV{"f2", &dumper.BeautifyNode{Type: "string", Value: sPtr(`""`)}},
			&dumper.StructKV{"f3", &dumper.BeautifyNode{
				Ptr:  sPtr("0xADDR"),
				Type: "testStruct",
				StructValues: []*dumper.StructKV{
					&dumper.StructKV{"f1", &dumper.BeautifyNode{Type: "int", Value: sPtr("2")}},
					&dumper.StructKV{"f2", &dumper.BeautifyNode{Type: "string", Value: sPtr(`""`)}},
					&dumper.StructKV{"f3", &dumper.BeautifyNode{
						Ptr:  sPtr("0xADDR"),
						Type: "testStruct",
						StructValues: []*dumper.StructKV{
							&dumper.StructKV{"f1", &dumper.BeautifyNode{Type: "int", Value: sPtr("3")}},
							&dumper.StructKV{"f2", &dumper.BeautifyNode{Type: "string", Value: sPtr(`""`)}},
							&dumper.StructKV{"f3", &dumper.BeautifyNode{Ptr: sPtr("0xADDR"), Type: "testStruct"}},
							&dumper.StructKV{"f4", &dumper.BeautifyNode{Type: "<INVALID>"}},
						},
					},
					},
					&dumper.StructKV{"f4", &dumper.BeautifyNode{Type: "<INVALID>"}},
				},
			},
			},
			&dumper.StructKV{"f4", &dumper.BeautifyNode{Type: "int", Value: sPtr("100")}},
		},
	}, "Struct with cycle beautify")
}

func TestDumpArray(t *testing.T) {
	check(t, [3]int{1, 2, 3}, "[3]int{int(1),int(2),int(3)}", "Array")
	check(t, &[3]int{1, 2, 3}, "&(0xADDR)[3]int{int(1),int(2),int(3)}", "Ptr to array")
	check(t, [3]interface{}{1, 2, 3}, "[3]interface {}{int(1),int(2),int(3)}", "Array of interface")
	check(t, ArrayAlias{1, 2, 3}, "dumper_test.ArrayAlias/*array*/{int(1),int(2),int(3)}", "Array alias")
	check(t, &ArrayAlias{1, 2, 3}, "&(0xADDR)dumper_test.ArrayAlias/*array*/{int(1),int(2),int(3)}", "Ptr to array alias")

	checkBeautify(t, [3]int{1, 2, 3}, &dumper.BeautifyNode{
		Type: "[3]int",
		ArrayValues: []*dumper.BeautifyNode{
			&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
		},
	}, "Array beautify")

	checkBeautify(t, &[3]int{1, 2, 3}, &dumper.BeautifyNode{
		Ptr:  sPtr("0xADDR"),
		Type: "[3]int",
		ArrayValues: []*dumper.BeautifyNode{
			&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
		},
	}, "Ptr to array beautify")

	checkBeautify(t, [3]interface{}{1, 2, 3}, &dumper.BeautifyNode{
		Type: "[3]interface {}",
		ArrayValues: []*dumper.BeautifyNode{
			&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
		},
	}, "Array of interface beautify")

	checkBeautify(t, ArrayAlias{1, 2, 3}, &dumper.BeautifyNode{
		Type: "dumper_test.ArrayAlias",
		ArrayValues: []*dumper.BeautifyNode{
			&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
		},
	}, "Array alias beautify")

	checkBeautify(t, &ArrayAlias{1, 2, 3}, &dumper.BeautifyNode{
		Ptr:  sPtr("0xADDR"),
		Type: "dumper_test.ArrayAlias",
		ArrayValues: []*dumper.BeautifyNode{
			&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
		},
	}, "Ptr to array beautify")
}

func TestDumpSlice(t *testing.T) {
	check(t, []int{1, 2, 3}, "[]int{int(1),int(2),int(3)}", "Slice")
	check(t, &[]int{1, 2, 3}, "&(0xADDR)[]int{int(1),int(2),int(3)}", "Ptr to slice")
	check(t, []interface{}{1, 2, 3}, "[]interface {}{int(1),int(2),int(3)}", "Slice of interface{}")
	check(t, SliceAlias{1, 2, 3}, "dumper_test.SliceAlias/*slice*/{int(1),int(2),int(3)}", "Slice alias")
	check(t, &SliceAlias{1, 2, 3}, "&(0xADDR)dumper_test.SliceAlias/*slice*/{int(1),int(2),int(3)}", "Ptr to slice alias")

	checkBeautify(t, []int{1, 2, 3}, &dumper.BeautifyNode{
		Type: "[]int",
		ArrayValues: []*dumper.BeautifyNode{
			&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
		},
	}, "Slice beautify")

	checkBeautify(t, &[]int{1, 2, 3}, &dumper.BeautifyNode{
		Ptr:  sPtr("0xADDR"),
		Type: "[]int",
		ArrayValues: []*dumper.BeautifyNode{
			&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
		},
	}, "Ptr to slice beautify")

	checkBeautify(t, []interface{}{1, 2, 3}, &dumper.BeautifyNode{
		Type: "[]interface {}",
		ArrayValues: []*dumper.BeautifyNode{
			&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
		},
	}, "Slice of interface beautify")

	checkBeautify(t, SliceAlias{1, 2, 3}, &dumper.BeautifyNode{
		Type: "dumper_test.SliceAlias",
		ArrayValues: []*dumper.BeautifyNode{
			&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
		},
	}, "Slice alias beautify")

	checkBeautify(t, &SliceAlias{1, 2, 3}, &dumper.BeautifyNode{
		Ptr:  sPtr("0xADDR"),
		Type: "dumper_test.SliceAlias",
		ArrayValues: []*dumper.BeautifyNode{
			&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
			&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
		},
	}, "Ptr to slice alias beautify")
}

func TestDumpMap(t *testing.T) {
	check(t, map[int]int{1: 10, 2: 20, 3: 30}, "map[int]int{int(1):int(10),int(2):int(20),int(3):int(30)}", "Map")
	check(t, &map[int]int{1: 10, 2: 20, 3: 30}, "&(0xADDR)map[int]int{int(1):int(10),int(2):int(20),int(3):int(30)}", "Ptr to Map")
	check(t, map[int]interface{}{1: 10, 2: 20, 3: 30}, "map[int]interface {}{int(1):int(10),int(2):int(20),int(3):int(30)}", "Map of interface{}")
	check(t, MapAlias{1: 10, 2: 20, 3: 30}, "dumper_test.MapAlias/*map*/{int(1):int(10),int(2):int(20),int(3):int(30)}", "Map alias")
	check(t, &MapAlias{1: 10, 2: 20, 3: 30}, "&(0xADDR)dumper_test.MapAlias/*map*/{int(1):int(10),int(2):int(20),int(3):int(30)}", "Ptr to map alias")

	checkBeautify(t, map[int]int{1: 10, 2: 20, 3: 30}, &dumper.BeautifyNode{
		Type: "map[int]int",
		HashValues: []*dumper.HashKV{
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("10")},
			},
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("20")},
			},
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("30")},
			},
		},
	}, "Map beautify")

	checkBeautify(t, &map[int]int{1: 10, 2: 20, 3: 30}, &dumper.BeautifyNode{
		Ptr:  sPtr("0xADDR"),
		Type: "map[int]int",
		HashValues: []*dumper.HashKV{
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("10")},
			},
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("20")},
			},
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("30")},
			},
		},
	}, "Ptr to map beautify")

	checkBeautify(t, map[int]interface{}{1: 10, 2: 20, 3: 30}, &dumper.BeautifyNode{
		Type: "map[int]interface {}",
		HashValues: []*dumper.HashKV{
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("10")},
			},
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("20")},
			},
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("30")},
			},
		},
	}, "Map of interface beautify")

	checkBeautify(t, MapAlias{1: 10, 2: 20, 3: 30}, &dumper.BeautifyNode{
		Type: "dumper_test.MapAlias",
		HashValues: []*dumper.HashKV{
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("10")},
			},
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("20")},
			},
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("30")},
			},
		},
	}, "Map alias beautify")

	checkBeautify(t, &MapAlias{1: 10, 2: 20, 3: 30}, &dumper.BeautifyNode{
		Ptr:  sPtr("0xADDR"),
		Type: "dumper_test.MapAlias",
		HashValues: []*dumper.HashKV{
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("1")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("10")},
			},
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("2")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("20")},
			},
			&dumper.HashKV{
				&dumper.BeautifyNode{Type: "int", Value: sPtr("3")},
				&dumper.BeautifyNode{Type: "int", Value: sPtr("30")},
			},
		},
	}, "Ptr to map alias beautify")

}
