//line dumper.y:2
package dumper

import (
	"bytes"
	"errors"
	__yyfmt__ "fmt"
	"regexp"
	"unicode/utf8"
)

//line dumper.y:13
type yySymType struct {
	yys       int
	string    string
	node      *BeautifyNode
	node_list []*BeautifyNode
	structKV  *StructKV
	structKVs []*StructKV
	hashKV    *HashKV
	hashKVs   []*HashKV
}

const STRING = 57346
const NUMBER = 57347
const PTR = 57348
const HINTARRAY = 57349
const HINTMAP = 57350
const NIL = 57351
const BOOL = 57352
const MAP = 57353
const INVALID = 57354
const INTERFACE = 57355
const IDENTIFIER = 57356
const UNKNOWN = 57357

var yyToknames = []string{
	"STRING",
	"NUMBER",
	"PTR",
	"HINTARRAY",
	"HINTMAP",
	"NIL",
	"BOOL",
	"MAP",
	"INVALID",
	"INTERFACE",
	"IDENTIFIER",
	"UNKNOWN",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line dumper.y:298

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
	simpleToken{"<???>", UNKNOWN},
}

var reTokens = []reToken{
	reToken{`^(?:\")(?:[^\\\"]*(?:\\.[^\\\"]*)*)(?:\")`, STRING},
	reToken{`^0x[0-9a-f]+`, PTR},
	reToken{`^-?\d+(?:[.,]\d+)?`, NUMBER},
	reToken{`^interface\s*\{\}`, INTERFACE},
	reToken{`^\*?[a-zA-Z_][a-zA-Z0-9\._]+`, IDENTIFIER},
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
			l := len(x.line)
			if l > 100 {
				l = 100
			}
			__yyfmt__.Printf("Lex: %s\n", x.line[:l])
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

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 49
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 183

var yyAct = []int{

	71, 2, 72, 8, 73, 44, 43, 118, 70, 32,
	37, 17, 111, 113, 12, 9, 14, 11, 84, 87,
	15, 12, 9, 14, 11, 81, 10, 15, 33, 47,
	24, 75, 52, 10, 132, 69, 56, 57, 55, 46,
	61, 54, 12, 9, 14, 11, 14, 49, 15, 25,
	67, 36, 74, 28, 10, 129, 104, 94, 95, 66,
	79, 74, 136, 83, 14, 49, 117, 86, 82, 116,
	94, 85, 112, 90, 97, 135, 30, 89, 31, 96,
	107, 21, 18, 14, 20, 100, 13, 15, 103, 99,
	14, 49, 108, 19, 16, 74, 110, 105, 94, 109,
	58, 14, 49, 131, 22, 91, 23, 114, 29, 38,
	130, 48, 98, 74, 94, 121, 120, 122, 74, 128,
	125, 91, 126, 91, 91, 124, 127, 45, 123, 119,
	74, 91, 133, 74, 91, 134, 115, 94, 91, 106,
	94, 91, 102, 101, 65, 93, 92, 65, 77, 78,
	39, 40, 64, 88, 80, 68, 41, 60, 59, 51,
	50, 42, 76, 15, 63, 53, 62, 39, 40, 14,
	49, 34, 35, 41, 26, 27, 6, 5, 42, 4,
	3, 1, 7,
}
var yyPact = []int{

	31, -1000, -1000, -1000, -1000, -1000, -1000, 70, 86, -1000,
	25, 167, 30, -1000, -1000, 90, -1000, 58, -1000, 4,
	164, 28, 146, 113, 15, 88, 140, 139, 156, 159,
	163, 113, 12, 77, 138, 137, 156, 147, 145, -1000,
	-1000, -1000, -1000, 131, -1000, 37, 156, 135, 11, -1000,
	31, 31, 7, 143, 129, 128, 156, 134, 1, 31,
	31, -6, -1000, -1000, -1000, 113, 10, 133, 31, 156,
	125, -1000, 124, -1000, 36, 51, -1000, -1000, -1000, 92,
	31, 156, 122, 121, 33, -1000, -1000, -1000, 31, 118,
	60, 31, -1000, -1000, 31, 3, 52, -11, 31, 115,
	49, -1000, -1000, 46, -17, 108, -1000, 31, -1000, -1000,
	-1000, -1000, 31, 156, 107, -1000, 31, 31, 156, -1000,
	105, 98, 35, -1000, 89, 82, 14, -1000, -1000, 31,
	-1000, -1000, 31, 54, 41, -1000, -1000,
}
var yyPgo = []int{

	0, 86, 182, 10, 3, 181, 0, 180, 179, 177,
	176, 8, 5, 6, 4, 2,
}
var yyR1 = []int{

	0, 5, 6, 6, 6, 6, 11, 11, 11, 1,
	2, 2, 7, 7, 7, 7, 3, 3, 3, 3,
	8, 8, 8, 13, 13, 12, 12, 12, 9, 9,
	9, 9, 9, 9, 9, 9, 10, 10, 10, 10,
	10, 10, 15, 15, 15, 14, 14, 4, 4,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 0, 1, 3, 4,
	1, 2, 5, 4, 2, 1, 1, 1, 1, 1,
	5, 4, 4, 1, 3, 0, 3, 3, 8, 7,
	7, 6, 9, 8, 6, 5, 9, 8, 11, 10,
	6, 5, 0, 1, 3, 3, 3, 1, 1,
}
var yyChk = []int{

	-1000, -5, -6, -7, -8, -9, -10, -2, -4, 12,
	23, 14, 11, -1, 13, 17, -1, -4, 12, 23,
	14, 11, 18, 20, 5, 24, 7, 8, 23, 18,
	18, 20, 5, 24, 7, 8, 23, -3, -1, 4,
	5, 10, 15, -13, -12, 14, 24, -4, 23, 14,
	20, 20, -4, 6, -3, -13, 24, -4, 23, 20,
	20, -4, 19, 19, 21, 16, 22, -4, 20, 24,
	-11, -6, -15, -14, -6, 24, 19, 19, 21, -4,
	20, 24, -11, -15, 24, -12, -6, 9, 20, -11,
	-4, 16, 21, 21, 16, 22, -4, 23, 20, -11,
	-4, 21, 21, -4, 23, -11, 21, 20, -6, -14,
	-6, 9, 20, 24, -11, 21, 20, 20, 24, 21,
	-11, -15, -4, 21, -11, -15, -4, 21, 21, 20,
	21, 21, 20, -15, -15, 21, 21,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 0, 0, 15,
	0, 47, 0, 10, 48, 0, 11, 0, 14, 0,
	47, 0, 0, 25, 0, 0, 0, 0, 0, 0,
	0, 25, 0, 0, 0, 0, 0, 0, 0, 16,
	17, 18, 19, 0, 23, 0, 0, 0, 0, 47,
	6, 42, 0, 0, 0, 0, 0, 0, 0, 6,
	42, 0, 13, 22, 21, 25, 0, 0, 6, 0,
	0, 7, 0, 43, 0, 0, 9, 12, 20, 0,
	6, 0, 0, 0, 0, 24, 26, 27, 6, 0,
	0, 0, 35, 41, 0, 0, 0, 0, 6, 0,
	0, 34, 40, 0, 0, 0, 31, 6, 8, 44,
	45, 46, 42, 0, 0, 30, 6, 42, 0, 29,
	0, 0, 0, 28, 0, 0, 0, 33, 37, 42,
	32, 36, 42, 0, 0, 39, 38,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 17, 3,
	18, 19, 3, 3, 16, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 22, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 23, 3, 24, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 20, 3, 21,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line dumper.y:47
		{
			yylex.(*exprLex).result = yyVAL.node
		}
	case 2:
		yyVAL.node = yyS[yypt-0].node
	case 3:
		yyVAL.node = yyS[yypt-0].node
	case 4:
		yyVAL.node = yyS[yypt-0].node
	case 5:
		yyVAL.node = yyS[yypt-0].node
	case 6:
		//line dumper.y:57
		{
			yyVAL.node_list = nil
		}
	case 7:
		//line dumper.y:61
		{
			if yyS[yypt-0].node == nil {
				yyVAL.node_list = nil
			} else {
				yyVAL.node_list = []*BeautifyNode{yyS[yypt-0].node}
			}
		}
	case 8:
		//line dumper.y:69
		{
			yyVAL.node_list = append(yyS[yypt-2].node_list, yyS[yypt-0].node)
		}
	case 9:
		//line dumper.y:74
		{
			yyVAL.string = yyS[yypt-1].string
		}
	case 10:
		yyVAL.string = yyS[yypt-0].string
	case 11:
		//line dumper.y:80
		{
			yyVAL.string = yyS[yypt-1].string + "," + yyS[yypt-0].string
		}
	case 12:
		//line dumper.y:85
		{
			yyVAL.node = &BeautifyNode{
				Ptr:   strPtr(yyS[yypt-4].string),
				Type:  yyS[yypt-3].string,
				Value: strPtr(yyS[yypt-1].string),
			}
		}
	case 13:
		//line dumper.y:93
		{
			yyVAL.node = &BeautifyNode{
				Type:  yyS[yypt-3].string,
				Value: strPtr(yyS[yypt-1].string),
			}
		}
	case 14:
		//line dumper.y:100
		{
			yyVAL.node = &BeautifyNode{
				Ptr:  strPtr(yyS[yypt-1].string),
				Type: yyS[yypt-0].string,
			}
		}
	case 15:
		//line dumper.y:107
		{
			yyVAL.node = &BeautifyNode{
				Type: yyS[yypt-0].string,
			}
		}
	case 16:
		yyVAL.string = yyS[yypt-0].string
	case 17:
		yyVAL.string = yyS[yypt-0].string
	case 18:
		yyVAL.string = yyS[yypt-0].string
	case 19:
		yyVAL.string = yyS[yypt-0].string
	case 20:
		//line dumper.y:119
		{
			yyVAL.node = &BeautifyNode{
				Ptr:          strPtr(yyS[yypt-4].string),
				Type:         yyS[yypt-3].string,
				StructValues: yyS[yypt-1].structKVs,
			}
		}
	case 21:
		//line dumper.y:127
		{
			yyVAL.node = &BeautifyNode{
				Type:         yyS[yypt-3].string,
				StructValues: yyS[yypt-1].structKVs,
			}
		}
	case 22:
		//line dumper.y:134
		{
			yyVAL.node = &BeautifyNode{
				Type: yyS[yypt-3].string,
				Ptr:  strPtr(yyS[yypt-1].string),
			}
		}
	case 23:
		//line dumper.y:142
		{
			if yyS[yypt-0].structKV == nil {
				yyVAL.structKVs = nil
			} else {
				yyVAL.structKVs = []*StructKV{yyS[yypt-0].structKV}
			}
		}
	case 24:
		//line dumper.y:150
		{
			yyVAL.structKVs = append(yyS[yypt-2].structKVs, yyS[yypt-0].structKV)
		}
	case 25:
		//line dumper.y:156
		{
			yyVAL.structKV = nil
		}
	case 26:
		//line dumper.y:160
		{
			yyVAL.structKV = &StructKV{yyS[yypt-2].string, yyS[yypt-0].node}
		}
	case 27:
		//line dumper.y:164
		{
			yyVAL.structKV = &StructKV{yyS[yypt-2].string, nil}
		}
	case 28:
		//line dumper.y:169
		{
			yyVAL.node = &BeautifyNode{
				Ptr:         strPtr(yyS[yypt-7].string),
				Type:        "[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 29:
		//line dumper.y:177
		{
			yyVAL.node = &BeautifyNode{
				Type:        "[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 30:
		//line dumper.y:184
		{
			yyVAL.node = &BeautifyNode{
				Ptr:         strPtr(yyS[yypt-6].string),
				Type:        "[]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 31:
		//line dumper.y:192
		{
			yyVAL.node = &BeautifyNode{
				Type:        "[]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 32:
		//line dumper.y:199
		{
			yyVAL.node = &BeautifyNode{
				Ptr:         strPtr(yyS[yypt-8].string),
				Type:        "[]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 33:
		//line dumper.y:207
		{
			yyVAL.node = &BeautifyNode{
				Type:        "[]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 34:
		//line dumper.y:214
		{
			yyVAL.node = &BeautifyNode{
				Ptr:         strPtr(yyS[yypt-5].string),
				Type:        yyS[yypt-4].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 35:
		//line dumper.y:222
		{
			yyVAL.node = &BeautifyNode{
				Type:        yyS[yypt-4].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 36:
		//line dumper.y:230
		{
			yyVAL.node = &BeautifyNode{
				Ptr:        strPtr(yyS[yypt-8].string),
				Type:       "map[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 37:
		//line dumper.y:238
		{
			yyVAL.node = &BeautifyNode{
				Type:       "map[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 38:
		//line dumper.y:245
		{
			yyVAL.node = &BeautifyNode{
				Type:       "map[" + yyS[yypt-7].string + "]" + yyS[yypt-3].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 39:
		//line dumper.y:252
		{
			yyVAL.node = &BeautifyNode{
				Type:       "map[" + yyS[yypt-7].string + "]" + yyS[yypt-3].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 40:
		//line dumper.y:259
		{
			yyVAL.node = &BeautifyNode{
				Ptr:        strPtr(yyS[yypt-5].string),
				Type:       yyS[yypt-4].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 41:
		//line dumper.y:267
		{
			yyVAL.node = &BeautifyNode{
				Type:       yyS[yypt-4].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 42:
		//line dumper.y:275
		{
			yyVAL.hashKVs = nil
		}
	case 43:
		//line dumper.y:279
		{
			yyVAL.hashKVs = []*HashKV{yyS[yypt-0].hashKV}
		}
	case 44:
		//line dumper.y:283
		{
			yyVAL.hashKVs = append(yyS[yypt-2].hashKVs, yyS[yypt-0].hashKV)
		}
	case 45:
		//line dumper.y:288
		{
			yyVAL.hashKV = &HashKV{yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 46:
		//line dumper.y:292
		{
			yyVAL.hashKV = &HashKV{yyS[yypt-2].node, nil}
		}
	case 47:
		yyVAL.string = yyS[yypt-0].string
	case 48:
		yyVAL.string = yyS[yypt-0].string
	}
	goto yystack /* stack new state and value */
}
