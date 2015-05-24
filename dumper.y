%{
package dumper;

import (
	"errors"
	"unicode/utf8"
	"regexp"
	__yyfmt__ "fmt"
)
%}

%union {
	string	  string
	node	  *BeautifyNode
	node_list []*BeautifyNode
	structKV  *StructKV
	structKVs []*StructKV
	hashKV    *HashKV
	hashKVs   []*HashKV
}

%token  STRING
%token 	NUMBER
%token 	PTR
%token 	NIL
%token	BOOL
%token	MAP
%token	INVALID
%token 	IDENTIFIER

%type	<string>	STRING, NUMBER, PTR, NIL, BOOL, MAP, INVALID, IDENTIFIER, ptr, ptrs, scalar_value
%type	<node>		top, expr, scalar, struct, array, hash
%type	<node_list> expr_list
%type	<structKV>	struct_kv
%type	<structKVs> struct_kvs
%type	<hashKV>	hash_kv
%type	<hashKVs>   hash_kvs

%%

top:		expr
				{
					yylex.(*exprLex).result = $$
				}

expr:		scalar
		|	struct
		|	array
		|	hash
		
expr_list:	expr
				{
					if $1 == nil {
						$$ = nil
					} else {
						$$ = []*BeautifyNode{$1}
					}
				}
		|	expr_list ',' expr	
				{
					$$ = append($1, $3)
				}
	
ptr:		'&' '(' PTR ')'
				{
					$$ = $3
				}
				
ptrs:		ptr
		|	ptrs ptr
				{
					$$ = $1 + "," + $2
				}
					
scalar:		ptrs IDENTIFIER '(' scalar_value ')'
				{ 
					$$ =  &BeautifyNode{
						Ptr:   strPtr($1),
						Type:  $2,
						Value: strPtr($4),		
					}
				}
		|	IDENTIFIER '(' scalar_value ')'
				{ 
					$$ =  &BeautifyNode{
						Type:  $1,
						Value: strPtr($3),		
					}
				}
		|	ptrs INVALID
				{
					$$ =  &BeautifyNode{
						Ptr:   strPtr($1),
						Type:  $2,		
					}
				}
		|	INVALID
				{
					$$ =  &BeautifyNode{
						Type:  $1,		
					}
				}

scalar_value:	STRING
			|	NUMBER
			|	BOOL

struct:		ptrs IDENTIFIER '{' struct_kvs '}'
				{ 
					$$ = &BeautifyNode{
						Ptr:          strPtr($1),
						Type:         $2,
						StructValues: $4,
					}
				}
		|	IDENTIFIER '{' struct_kvs '}'
				{ 
					$$ = &BeautifyNode{
						Type:         $1,
						StructValues: $3,
					}
				}
		|	IDENTIFIER '(' ptr ')'
				{ 
					$$ = &BeautifyNode{
						Type:         $1,
						Ptr:          strPtr($3),
					}
				}
				
struct_kvs:	struct_kv
				{
					if $1 == nil {
						$$ = nil
					} else {
						$$ = []*StructKV{$1}
					}
				}
		|	struct_kvs ',' struct_kv
				{
					$$ = append($1, $3)
				}


struct_kv:	
				{
					$$ = nil
				}
		|	IDENTIFIER ':' scalar
				{
					$$ = &StructKV{$1, $3}
				}
		|	IDENTIFIER ':' NIL
				{
					$$ = &StructKV{$1, nil}
				}
		|	IDENTIFIER ':' struct
				{
					$$ = &StructKV{$1, $3}
				}
				
array:		ptrs '[' NUMBER ']' IDENTIFIER '{' expr_list '}'
				{
					$$ = &BeautifyNode{
						Ptr:         strPtr($1),
						Type:        "[" + $3 + "]" + $5,
						ArrayValues: $7,
					}
				}
		|	'[' NUMBER ']' IDENTIFIER '{' expr_list '}'
				{
					$$ = &BeautifyNode{
						Type:        "[" + $2 + "]" + $4,
						ArrayValues: $6,
					}
				}
		|	ptrs '[' ']' IDENTIFIER '{' expr_list '}'
				{
					$$ = &BeautifyNode{
						Ptr:         strPtr($1),
						Type:        "[]" + $4,
						ArrayValues: $6,
					}
				}
		|	'[' ']' IDENTIFIER '{' expr_list '}'
				{
					$$ = &BeautifyNode{
						Type:        "[]" + $3,
						ArrayValues: $5,
					}
				}
				
hash:		ptrs MAP '[' IDENTIFIER ']' IDENTIFIER '{' hash_kvs '}'
				{
					$$ = &BeautifyNode{
						Ptr:		strPtr($1),	
						Type:       "map[" + $4 + "]" + $6,
						HashValues: $8,
					}
				}
		|	MAP '[' IDENTIFIER ']' IDENTIFIER '{' hash_kvs '}'
				{
					$$ = &BeautifyNode{
						Type:       "map[" + $3 + "]" + $5,
						HashValues: $7,
					}
				}
				
hash_kvs:	
			{
				$$ = nil
			}
		|	hash_kv
			{
				$$ = []*HashKV{$1}
			}
		|	hash_kvs ',' hash_kv
			{
				$$ = append($1, $3)
			}
		
hash_kv:	scalar ':' scalar
			{
				$$ = &HashKV{$1, $3}
			}
		|	scalar ':' NIL
			{
				$$ = &HashKV{$1, nil}
			}
%%

type exprLex struct {
        line   []byte
        peek   rune
        result *BeautifyNode
        err error
}

func strPtr(s string) *string {return &s}

var stringRe, numberRe, ptrRe, nilRe, boolRe, mapRe, invalidRe, identifierRe *regexp.Regexp

func init() {
	var err error

	stringRe, err = regexp.Compile(`^(?:\")(?:[^\\\"]*(?:\\.[^\\\"]*)*)(?:\")`)
	if err != nil {
		panic(err)
	}

	numberRe, err = regexp.Compile(`^-?\d+(?:[.,]\d+)?`)
	if err != nil {
		panic(err)
	}

	ptrRe, err = regexp.Compile(`^0x[0-9a-f]+`)
	if err != nil {
		panic(err)
	}
	
	nilRe, err = regexp.Compile(`^\(nil\)`)
	if err != nil {
		panic(err)
	}
	
	nilRe, err = regexp.Compile(`^\(nil\)`)
	if err != nil {
		panic(err)
	}

	boolRe, err = regexp.Compile(`^(?:true|false)`)
	if err != nil {
		panic(err)
	}

	mapRe, err = regexp.Compile(`^map`)
	if err != nil {
		panic(err)
	}
	
	invalidRe, err = regexp.Compile(`^<INVALID>`)
	if err != nil {
		panic(err)
	}
	
	identifierRe, err = regexp.Compile(`^[a-zA-Z_][a-zA-Z0-9\._]+`)
	if err != nil {
		panic(err)
	}
}

func (x *exprLex) Lex(yylval *yySymType) int {
	for {
		if yyDebug >= 1 {
			__yyfmt__.Printf("Lex: %s\n", x.line)
		}
		if len(x.line) == 0 {
			return 0
		}
	
		if m := stringRe.Find(x.line); m != nil {
			if yyDebug >= 1 {
				__yyfmt__.Print("Matched String\n")
			}
			x.line = x.line[len(m):]
			yylval.string = string(m)
			return STRING
		} else if m := ptrRe.Find(x.line); m != nil {
			if yyDebug >= 1 {
				__yyfmt__.Print("Matched PTR\n")
			}
			x.line = x.line[len(m):]
			yylval.string = string(m)
			return PTR
		} else if m := nilRe.Find(x.line); m != nil {
			if yyDebug >= 1 {
				__yyfmt__.Print("Matched NIL\n")
			}
			x.line = x.line[len(m):]
			yylval.string = string(m)
			return NIL
		} else if m := numberRe.Find(x.line); m != nil {
			if yyDebug >= 1 {
				__yyfmt__.Print("Matched Number\n")
			}
			x.line = x.line[len(m):]
			yylval.string = string(m)
			return NUMBER
		} else if m := boolRe.Find(x.line); m != nil {
			if yyDebug >= 1 {
				__yyfmt__.Print("Matched BOOL\n")
			}
			x.line = x.line[len(m):]
			yylval.string = string(m)
			return BOOL
		} else if m := mapRe.Find(x.line); m != nil {
			if yyDebug >= 1 {
				__yyfmt__.Print("Matched MAP\n")
			}
			x.line = x.line[len(m):]
			yylval.string = string(m)
			return MAP
		} else if m := invalidRe.Find(x.line); m != nil {
			if yyDebug >= 1 {
				__yyfmt__.Print("Matched INVALID\n")
			}
			x.line = x.line[len(m):]
			yylval.string = string(m)
			return INVALID
		} else if m := identifierRe.Find(x.line); m != nil {
			if yyDebug >= 1 {
				__yyfmt__.Print("Matched IDENTIFIER\n")
			}
			x.line = x.line[len(m):]
			yylval.string = string(m)
			return IDENTIFIER
		} else {
			c, size := utf8.DecodeRune(x.line)
			x.line = x.line[size:]
			if c == utf8.RuneError && size == 1 {
				x.Error("Invalid utf8")
				return 0
			}
			if yyDebug >= 1 {
				__yyfmt__.Print("Matched char\n")
			}
			switch c {
			case ' ', '\t', '\n', '\r':
				continue
			default:
				return int(c)
			}
		}
	}
}

func (x *exprLex) Error(s string) {
	x.err = errors.New(s)
}
