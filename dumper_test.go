package dumper_test

import (
	"bytes"
	"github.com/sergei-svistunov/dumper"
	"regexp"
	"testing"
)

var addrRe, _ = regexp.Compile("\\&\\(0x[0-9a-f]+\\)")

type testStruct struct {
	f1 int
	f2 string
	f3 *testStruct
}

func check(t *testing.T, v interface{}, expected, description string) {
	buf := &bytes.Buffer{}
	dumper.Dump(buf, v)

	gotten := addrRe.ReplaceAllString(buf.String(), "&(0xADDR)")

	if gotten != expected {
		t.Errorf("%s: expected \"%s\", gotten \"%s\"", description, expected, gotten)
	}
}

func TestDumpInt(t *testing.T) {
	check(t, int(123), "int(123)", "Int")
	check(t, int8(123), "int8(123)", "Int8")
	check(t, int16(123), "int16(123)", "Int16")
	check(t, int32(123), "int32(123)", "Int32")
	check(t, int64(123), "int64(123)", "Int64")
}

func TestDumpUint(t *testing.T) {
	check(t, uint(123), "uint(123)", "Uint")
	check(t, uint8(123), "uint8(123)", "Uint8")
	check(t, uint16(123), "uint16(123)", "Uint16")
	check(t, uint32(123), "uint32(123)", "Uint32")
	check(t, uint64(123), "uint64(123)", "Uint64")
}

func TestDumpBool(t *testing.T) {
	check(t, true, "true", "Bool true")
	check(t, false, "false", "Bool false")
}

func TestDumpFloat(t *testing.T) {
	check(t, float32(10.5), "float32(10.500000)", "Float32")
	check(t, float64(10.5), "float64(10.500000)", "Float64")
}

func TestDumpString(t *testing.T) {
	check(t, string("Test\n,\t,\""), `string("Test\n,\t,\"")`, "String")
}

func TestDumpPtr(t *testing.T) {
	intVar := int(123)
	check(t, &intVar, "&(0xADDR)int(123)", "Ptr to int")
	intVarPtr := &intVar
	check(t, &intVarPtr, "&(0xADDR)&(0xADDR)int(123)", "Ptr to ptr to int")
}

func TestDumpInterface(t *testing.T) {
	var interfaceVar interface{}

	check(t, interfaceVar, "<INVALID>", "Empty interface")
	check(t, &interfaceVar, "&(0xADDR)<INVALID>", "Ptr to empty interface")

	interfaceVar = int(123)
	check(t, interfaceVar, "int(123)", "Int in interface")
	check(t, &interfaceVar, "&(0xADDR)int(123)", "Ptr to Int in interface")
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
	}
	check(t, s, `testStruct{f1:int(1),f2:string(""),f3:&(0xADDR)testStruct{f1:int(2),f2:string(""),f3:&(0xADDR)testStruct{f1:int(3),f2:string(""),f3:(nil)}}}`, "Struct")

	s.f3.f3.f3 = &s
	check(t, s, `testStruct{f1:int(1),f2:string(""),f3:&(0xADDR)testStruct{f1:int(2),f2:string(""),f3:&(0xADDR)testStruct{f1:int(3),f2:string(""),f3:&(0xADDR)testStruct{f1:int(1),f2:string(""),f3:testStruct(&(0xADDR))}}}}`, "Struct with cycle")
}

func TestDumpArray(t *testing.T) {
	check(t, [3]int{1, 2, 3}, "[3]int{int(1),int(2),int(3)}", "Array")
}

func TestDumpSlice(t *testing.T) {
	check(t, []int{1, 2, 3}, "[]int{int(1),int(2),int(3)}", "Slice")
}

func TestDumpMap(t *testing.T) {
	t.Skip("Need keys sorting")
	check(t, map[int]int{1: 10, 2: 20, 3: 30}, "map[int]int{int(1):int(10),int(2):int(20),int(3):int(30)}", "Map")
}
