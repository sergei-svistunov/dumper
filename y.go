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
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line dumper.y:263

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

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 43
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 143

var yyAct = []int{

	67, 2, 68, 69, 66, 43, 42, 8, 37, 101,
	79, 12, 9, 14, 11, 17, 15, 71, 54, 32,
	45, 82, 10, 12, 9, 14, 11, 24, 15, 36,
	28, 89, 13, 46, 10, 105, 50, 33, 53, 52,
	16, 55, 63, 102, 58, 25, 12, 9, 14, 11,
	70, 15, 91, 64, 88, 38, 83, 10, 70, 111,
	78, 77, 75, 30, 81, 31, 76, 22, 80, 23,
	84, 21, 18, 14, 20, 65, 15, 57, 88, 90,
	85, 92, 19, 110, 73, 108, 98, 95, 96, 70,
	100, 56, 99, 85, 85, 85, 103, 88, 106, 104,
	97, 72, 94, 70, 85, 107, 70, 88, 109, 93,
	85, 62, 87, 62, 29, 86, 74, 60, 61, 49,
	48, 39, 40, 59, 14, 47, 44, 41, 39, 40,
	34, 35, 51, 15, 41, 26, 27, 6, 5, 4,
	3, 1, 7,
}
var yyPact = []int{

	35, -1000, -1000, -1000, -1000, -1000, -1000, 60, 50, -1000,
	22, 128, 8, -1000, -1000, 97, -1000, 46, -1000, 14,
	123, 7, 117, 112, -3, 111, 101, 100, 111, 126,
	124, 112, -5, 111, 72, 58, 111, 105, 99, -1000,
	-1000, -1000, 98, -1000, 21, 111, 56, -1000, 35, 35,
	-6, 83, 66, 96, 111, 47, 35, 35, -13, -1000,
	-1000, -1000, 112, 12, 37, 35, 95, -1000, 92, -1000,
	10, 111, -1000, -1000, -1000, 33, 35, 89, 82, 111,
	-1000, -1000, -1000, 35, 80, 35, -1000, -1000, 35, 0,
	24, 35, 79, -1000, -1000, 16, 78, -1000, -1000, -1000,
	-1000, -1000, 35, 65, -1000, 35, -1000, 63, -1000, 39,
	-1000, -1000,
}
var yyPgo = []int{

	0, 32, 142, 8, 7, 141, 0, 140, 139, 138,
	137, 4, 5, 6, 3, 2,
}
var yyR1 = []int{

	0, 5, 6, 6, 6, 6, 11, 11, 1, 2,
	2, 7, 7, 7, 7, 3, 3, 3, 8, 8,
	8, 13, 13, 12, 12, 12, 9, 9, 9, 9,
	9, 9, 10, 10, 10, 10, 15, 15, 15, 14,
	14, 4, 4,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 3, 4, 1,
	2, 5, 4, 2, 1, 1, 1, 1, 5, 4,
	4, 1, 3, 0, 3, 3, 8, 7, 7, 6,
	6, 5, 9, 8, 6, 5, 0, 1, 3, 3,
	3, 1, 1,
}
var yyChk = []int{

	-1000, -5, -6, -7, -8, -9, -10, -2, -4, 12,
	22, 14, 11, -1, 13, 16, -1, -4, 12, 22,
	14, 11, 17, 19, 5, 23, 7, 8, 22, 17,
	17, 19, 5, 23, 7, 8, 22, -3, -1, 4,
	5, 10, -13, -12, 14, 23, -4, 14, 19, 19,
	-4, 6, -3, -13, 23, -4, 19, 19, -4, 18,
	18, 20, 15, 21, -4, 19, -11, -6, -15, -14,
	-6, 23, 18, 18, 20, -4, 19, -11, -15, 23,
	-12, -6, 9, 19, -11, 15, 20, 20, 15, 21,
	-4, 19, -11, 20, 20, -4, -11, 20, -6, -14,
	-6, 9, 19, -11, 20, 19, 20, -15, 20, -15,
	20, 20,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 0, 0, 14,
	0, 41, 0, 9, 42, 0, 10, 0, 13, 0,
	41, 0, 0, 23, 0, 0, 0, 0, 0, 0,
	0, 23, 0, 0, 0, 0, 0, 0, 0, 15,
	16, 17, 0, 21, 0, 0, 0, 41, 0, 36,
	0, 0, 0, 0, 0, 0, 0, 36, 0, 12,
	20, 19, 23, 0, 0, 0, 0, 6, 0, 37,
	0, 0, 8, 11, 18, 0, 0, 0, 0, 0,
	22, 24, 25, 0, 0, 0, 31, 35, 0, 0,
	0, 0, 0, 30, 34, 0, 0, 29, 7, 38,
	39, 40, 36, 0, 28, 36, 27, 0, 26, 0,
	33, 32,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 16, 3,
	17, 18, 3, 3, 15, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 21, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 22, 3, 23, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 19, 3, 20,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14,
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
		//line dumper.y:46
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
		//line dumper.y:56
		{
			if yyS[yypt-0].node == nil {
				yyVAL.node_list = nil
			} else {
				yyVAL.node_list = []*BeautifyNode{yyS[yypt-0].node}
			}
		}
	case 7:
		//line dumper.y:64
		{
			yyVAL.node_list = append(yyS[yypt-2].node_list, yyS[yypt-0].node)
		}
	case 8:
		//line dumper.y:69
		{
			yyVAL.string = yyS[yypt-1].string
		}
	case 9:
		yyVAL.string = yyS[yypt-0].string
	case 10:
		//line dumper.y:75
		{
			yyVAL.string = yyS[yypt-1].string + "," + yyS[yypt-0].string
		}
	case 11:
		//line dumper.y:80
		{
			yyVAL.node = &BeautifyNode{
				Ptr:   strPtr(yyS[yypt-4].string),
				Type:  yyS[yypt-3].string,
				Value: strPtr(yyS[yypt-1].string),
			}
		}
	case 12:
		//line dumper.y:88
		{
			yyVAL.node = &BeautifyNode{
				Type:  yyS[yypt-3].string,
				Value: strPtr(yyS[yypt-1].string),
			}
		}
	case 13:
		//line dumper.y:95
		{
			yyVAL.node = &BeautifyNode{
				Ptr:  strPtr(yyS[yypt-1].string),
				Type: yyS[yypt-0].string,
			}
		}
	case 14:
		//line dumper.y:102
		{
			yyVAL.node = &BeautifyNode{
				Type: yyS[yypt-0].string,
			}
		}
	case 15:
		yyVAL.string = yyS[yypt-0].string
	case 16:
		yyVAL.string = yyS[yypt-0].string
	case 17:
		yyVAL.string = yyS[yypt-0].string
	case 18:
		//line dumper.y:113
		{
			yyVAL.node = &BeautifyNode{
				Ptr:          strPtr(yyS[yypt-4].string),
				Type:         yyS[yypt-3].string,
				StructValues: yyS[yypt-1].structKVs,
			}
		}
	case 19:
		//line dumper.y:121
		{
			yyVAL.node = &BeautifyNode{
				Type:         yyS[yypt-3].string,
				StructValues: yyS[yypt-1].structKVs,
			}
		}
	case 20:
		//line dumper.y:128
		{
			yyVAL.node = &BeautifyNode{
				Type: yyS[yypt-3].string,
				Ptr:  strPtr(yyS[yypt-1].string),
			}
		}
	case 21:
		//line dumper.y:136
		{
			if yyS[yypt-0].structKV == nil {
				yyVAL.structKVs = nil
			} else {
				yyVAL.structKVs = []*StructKV{yyS[yypt-0].structKV}
			}
		}
	case 22:
		//line dumper.y:144
		{
			yyVAL.structKVs = append(yyS[yypt-2].structKVs, yyS[yypt-0].structKV)
		}
	case 23:
		//line dumper.y:150
		{
			yyVAL.structKV = nil
		}
	case 24:
		//line dumper.y:154
		{
			yyVAL.structKV = &StructKV{yyS[yypt-2].string, yyS[yypt-0].node}
		}
	case 25:
		//line dumper.y:158
		{
			yyVAL.structKV = &StructKV{yyS[yypt-2].string, nil}
		}
	case 26:
		//line dumper.y:163
		{
			yyVAL.node = &BeautifyNode{
				Ptr:         strPtr(yyS[yypt-7].string),
				Type:        "[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 27:
		//line dumper.y:171
		{
			yyVAL.node = &BeautifyNode{
				Type:        "[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 28:
		//line dumper.y:178
		{
			yyVAL.node = &BeautifyNode{
				Ptr:         strPtr(yyS[yypt-6].string),
				Type:        "[]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 29:
		//line dumper.y:186
		{
			yyVAL.node = &BeautifyNode{
				Type:        "[]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 30:
		//line dumper.y:193
		{
			yyVAL.node = &BeautifyNode{
				Ptr:         strPtr(yyS[yypt-5].string),
				Type:        yyS[yypt-4].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 31:
		//line dumper.y:201
		{
			yyVAL.node = &BeautifyNode{
				Type:        yyS[yypt-4].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 32:
		//line dumper.y:209
		{
			yyVAL.node = &BeautifyNode{
				Ptr:        strPtr(yyS[yypt-8].string),
				Type:       "map[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 33:
		//line dumper.y:217
		{
			yyVAL.node = &BeautifyNode{
				Type:       "map[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 34:
		//line dumper.y:224
		{
			yyVAL.node = &BeautifyNode{
				Ptr:        strPtr(yyS[yypt-5].string),
				Type:       yyS[yypt-4].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 35:
		//line dumper.y:232
		{
			yyVAL.node = &BeautifyNode{
				Type:       yyS[yypt-4].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 36:
		//line dumper.y:240
		{
			yyVAL.hashKVs = nil
		}
	case 37:
		//line dumper.y:244
		{
			yyVAL.hashKVs = []*HashKV{yyS[yypt-0].hashKV}
		}
	case 38:
		//line dumper.y:248
		{
			yyVAL.hashKVs = append(yyS[yypt-2].hashKVs, yyS[yypt-0].hashKV)
		}
	case 39:
		//line dumper.y:253
		{
			yyVAL.hashKV = &HashKV{yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 40:
		//line dumper.y:257
		{
			yyVAL.hashKV = &HashKV{yyS[yypt-2].node, nil}
		}
	case 41:
		yyVAL.string = yyS[yypt-0].string
	case 42:
		yyVAL.string = yyS[yypt-0].string
	}
	goto yystack /* stack new state and value */
}
