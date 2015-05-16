# dumper - simple data printer

## Why?

For example, we have simple structure
```go
type t struct {
	f1 int
	f2 *t
}
```
and variable of type `*t`
```go
v := &t{
	f1: 1,
	f2: &t{
		f1: 2,
	},
}
v.f2.f2 = v //cycle
```
###Task: Save dump of `v` to log file
Ways:
######1. `fmt` - no information about field f2, only address:
```go
fmt.FPrintf(w, "%#v\n", v) //&t{f1:1, f2:(*t)(0xc20800a270)}
```
######2. JSON - all fields private, JSON prints `{}`
######3. `dumper`
```go
fmt.Printf("%s\n", dumper.DumpToString(v)) //&(0xc20800a260)t{f1:int(1),f2:&(0xc20800a270)t{f1:int(2),f2:t(&(0xc20800a260))}}
```

##Description
`dumper` doesn't produce valid Go code

##How to install
`go get github.com/sergei-svistunov/dumper`
