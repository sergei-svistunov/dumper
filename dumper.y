%{
package dumper;

import (
	"errors"
	"unicode/utf8"
	"regexp"
	"bytes"
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
%token	HINTARRAY
%token	HINTMAP
%token 	NIL
%token	BOOL
%token	MAP
%token	INVALID
%token	INTERFACE
%token 	IDENTIFIER

%type	<string>	STRING, NUMBER, PTR, HINTARRAY, HINTMAP, NIL, BOOL, MAP, INVALID, INTERFACE, IDENTIFIER, ptr, ptrs, scalar_value, type_name
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
					
scalar:		ptrs type_name '(' scalar_value ')'
				{ 
					$$ =  &BeautifyNode{
						Ptr:   strPtr($1),
						Type:  $2,
						Value: strPtr($4),		
					}
				}
		|	type_name '(' scalar_value ')'
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

struct:		ptrs type_name '{' struct_kvs '}'
				{ 
					$$ = &BeautifyNode{
						Ptr:          strPtr($1),
						Type:         $2,
						StructValues: $4,
					}
				}
		|	type_name '{' struct_kvs '}'
				{ 
					$$ = &BeautifyNode{
						Type:         $1,
						StructValues: $3,
					}
				}
		|	type_name '(' ptr ')'
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
		|	IDENTIFIER ':' expr
				{
					$$ = &StructKV{$1, $3}
				}
		|	IDENTIFIER ':' NIL
				{
					$$ = &StructKV{$1, nil}
				}
				
array:		ptrs '[' NUMBER ']' type_name '{' expr_list '}'
				{
					$$ = &BeautifyNode{
						Ptr:         strPtr($1),
						Type:        "[" + $3 + "]" + $5,
						ArrayValues: $7,
					}
				}
		|	'[' NUMBER ']' type_name '{' expr_list '}'
				{
					$$ = &BeautifyNode{
						Type:        "[" + $2 + "]" + $4,
						ArrayValues: $6,
					}
				}
		|	ptrs '[' ']' type_name '{' expr_list '}'
				{
					$$ = &BeautifyNode{
						Ptr:         strPtr($1),
						Type:        "[]" + $4,
						ArrayValues: $6,
					}
				}
		|	'[' ']' type_name '{' expr_list '}'
				{
					$$ = &BeautifyNode{
						Type:        "[]" + $3,
						ArrayValues: $5,
					}
				}
		|	ptrs IDENTIFIER HINTARRAY '{' expr_list '}'
				{
					$$ = &BeautifyNode{
						Ptr:         strPtr($1),
						Type:        $2,
						ArrayValues: $5,
					}
				}
		|	IDENTIFIER HINTARRAY '{' expr_list '}'
				{
					$$ = &BeautifyNode{
						Type:        $1,
						ArrayValues: $4,
					}
				}
				
hash:		ptrs MAP '[' type_name ']' type_name '{' hash_kvs '}'
				{
					$$ = &BeautifyNode{
						Ptr:		strPtr($1),	
						Type:       "map[" + $4 + "]" + $6,
						HashValues: $8,
					}
				}
		|	MAP '[' type_name ']' type_name '{' hash_kvs '}'
				{
					$$ = &BeautifyNode{
						Type:       "map[" + $3 + "]" + $5,
						HashValues: $7,
					}
				}		
		|	ptrs IDENTIFIER HINTMAP '{' hash_kvs '}'
				{
					$$ = &BeautifyNode{
						Ptr:		strPtr($1),	
						Type:       $2,
						HashValues: $5,
					}
				}
		|	IDENTIFIER HINTMAP '{' hash_kvs '}'
				{
					$$ = &BeautifyNode{
						Type:       $1,
						HashValues: $4,
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
		
hash_kv:	expr ':' expr
			{
				$$ = &HashKV{$1, $3}
			}
		|	expr ':' NIL
			{
				$$ = &HashKV{$1, nil}
			}
			
type_name:	IDENTIFIER
		|	INTERFACE
%%

type exprLex struct {
	line   []byte
	peek   rune
	result *BeautifyNode
	err    error
}

type simpleToken struct {
	token string
	value int
}

type reToken struct {
	re    string
	value int
}

type compiledReToken struct {
	re    *regexp.Regexp
	value int
}

var simpleTokens = []simpleToken{
	simpleToken{"(nil)", NIL},
	simpleToken{"true", BOOL},
	simpleToken{"false", BOOL},
	simpleToken{"map", MAP},
	simpleToken{"<INVALID>", INVALID},
	simpleToken{"/*array*/", HINTARRAY},
	simpleToken{"/*slice*/", HINTARRAY},
	simpleToken{"/*map*/", HINTMAP},
}

var reTokens = []reToken{
	reToken{`^(?:\")(?:[^\\\"]*(?:\\.[^\\\"]*)*)(?:\")`, STRING},
	reToken{`^0x[0-9a-f]+`, PTR},	
	reToken{`^-?\d+(?:[.,]\d+)?`, NUMBER},	
	reToken{`^interface\s*\{\}`, INTERFACE},
	reToken{`^[a-zA-Z_][a-zA-Z0-9\._]+`, IDENTIFIER},
}

var compiledReTokens = getCompiledReTokens()

func getCompiledReTokens() []compiledReToken {
	result := make([]compiledReToken, len(reTokens))
	var err error
	for i, _ := range reTokens {
		result[i].re, err = regexp.Compile(reTokens[i].re)
		if err != nil {
			panic(err)
		}
		result[i].value = reTokens[i].value
	}

	return result
}

func strPtr(s string) *string { return &s }

func (x *exprLex) Lex(yylval *yySymType) int {
	for {
		if yyDebug >= 1 {
			__yyfmt__.Printf("Lex: %s\n", x.line)
		}
		if len(x.line) == 0 {
			return 0
		}

		for _, token := range simpleTokens {
			if bytes.HasPrefix(x.line, []byte(token.token)) {
				if yyDebug >= 1 {
					__yyfmt__.Printf("Matched %s\n", token.token)
				}
				x.line = x.line[len(token.token):]
				yylval.string = token.token
				return token.value
			}
		}

		for _, token := range compiledReTokens {
			if m := token.re.Find(x.line); m != nil {
				if yyDebug >= 1 {
					__yyfmt__.Printf("Matched %s\n", yyTokname(token.value))
				}
				x.line = x.line[len(m):]
				yylval.string = string(m)
				return token.value
			}
		}

		c, size := utf8.DecodeRune(x.line)
		x.line = x.line[size:]
		if c == utf8.RuneError && size == 1 {
			x.Error("Invalid utf8")
			return 0
		}
		if yyDebug >= 1 {
			__yyfmt__.Printf("Matched char '%s'\n", string(c))
		}
		switch c {
		case ' ', '\t', '\n', '\r':
			continue
		default:
			return int(c)
		}

	}
}

func (x *exprLex) Error(s string) {
	x.err = errors.New(s)
}
