//line dumper.y:2
package dumper

import (
	"errors"
	__yyfmt__ "fmt"
	"regexp"
	"unicode/utf8"
)

//line dumper.y:12
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
const NIL = 57349
const BOOL = 57350
const MAP = 57351
const INVALID = 57352
const IDENTIFIER = 57353

var yyToknames = []string{
	"STRING",
	"NUMBER",
	"PTR",
	"NIL",
	"BOOL",
	"MAP",
	"INVALID",
	"IDENTIFIER",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line dumper.y:230

type exprLex struct {
	line   []byte
	peek   rune
	result *BeautifyNode
	err    error
}

func strPtr(s string) *string { return &s }

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

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 38
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 128

var yyAct = []int{

	3, 12, 7, 82, 83, 67, 68, 2, 4, 14,
	36, 35, 60, 30, 11, 9, 8, 27, 13, 21,
	54, 31, 44, 38, 10, 18, 16, 15, 29, 13,
	23, 90, 28, 91, 22, 17, 94, 90, 43, 42,
	74, 74, 89, 74, 74, 87, 80, 72, 78, 75,
	50, 51, 62, 50, 65, 57, 79, 25, 49, 26,
	64, 61, 19, 76, 20, 71, 70, 14, 66, 59,
	53, 56, 73, 55, 48, 47, 77, 84, 25, 85,
	84, 81, 85, 88, 93, 9, 86, 14, 13, 69,
	24, 84, 96, 85, 85, 95, 97, 37, 58, 9,
	86, 63, 13, 52, 9, 8, 46, 13, 32, 33,
	16, 92, 34, 13, 45, 16, 15, 13, 13, 40,
	39, 32, 33, 6, 41, 34, 5, 1,
}
var yyPact = []int{

	5, -1000, -1000, -1000, -1000, -1000, -1000, 16, 48, -1000,
	14, 11, -1000, 76, -1000, 43, -1000, 12, 9, 104,
	86, 3, 109, 108, 118, 117, 86, 2, 103, 95,
	60, 59, -1000, -1000, -1000, 41, -1000, 33, 92, 54,
	0, 58, 56, 38, 87, 53, -8, -1000, -1000, -1000,
	86, 94, 52, 5, 78, -1000, -1000, -1000, 50, 5,
	36, -1000, -1000, -1000, -1000, 105, 5, 32, -1000, 47,
	5, 31, 40, 29, 5, -1000, 75, 28, -1000, 75,
	-1000, -1000, 25, -1000, 15, 100, 70, -1000, 19, -1000,
	75, 89, 64, 117, -1000, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 1, 2, 13, 127, 6, 0, 8, 126, 123,
	5, 10, 11, 4, 3,
}
var yyR1 = []int{

	0, 4, 5, 5, 5, 5, 10, 10, 1, 2,
	2, 6, 6, 6, 6, 3, 3, 3, 7, 7,
	7, 12, 12, 11, 11, 11, 11, 8, 8, 8,
	8, 9, 9, 14, 14, 14, 13, 13,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 3, 4, 1,
	2, 5, 4, 2, 1, 1, 1, 1, 5, 4,
	4, 1, 3, 0, 3, 3, 3, 8, 7, 7,
	6, 9, 8, 0, 1, 3, 3, 3,
}
var yyChk = []int{

	-1000, -4, -5, -6, -7, -8, -9, -2, 11, 10,
	19, 9, -1, 13, -1, 11, 10, 19, 9, 14,
	16, 5, 20, 19, 14, 14, 16, 5, 20, 19,
	-3, -1, 4, 5, 8, -12, -11, 11, 20, 11,
	11, 6, -3, -12, 20, 11, 11, 15, 15, 17,
	12, 18, 11, 16, 20, 15, 15, 17, 11, 16,
	20, -11, -6, 7, -7, -2, 16, -10, -5, 11,
	16, -10, 11, -10, 12, 17, 16, -10, 17, 16,
	17, -5, -14, -13, -6, -2, 11, 17, -14, 17,
	12, 18, 11, 14, 17, -13, -6, 7,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 0, 0, 14,
	0, 0, 9, 0, 10, 0, 13, 0, 0, 0,
	23, 0, 0, 0, 0, 0, 23, 0, 0, 0,
	0, 0, 15, 16, 17, 0, 21, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 12, 20, 19,
	23, 0, 0, 0, 0, 8, 11, 18, 0, 0,
	0, 22, 24, 25, 26, 0, 0, 0, 6, 0,
	0, 0, 0, 0, 0, 30, 33, 0, 29, 33,
	28, 7, 0, 34, 0, 0, 0, 27, 0, 32,
	0, 0, 0, 0, 31, 35, 36, 37,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 13, 3,
	14, 15, 3, 3, 12, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 18, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 19, 3, 20, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 16, 3, 17,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
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
		//line dumper.y:42
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
		//line dumper.y:52
		{
			if yyS[yypt-0].node == nil {
				yyVAL.node_list = nil
			} else {
				yyVAL.node_list = []*BeautifyNode{yyS[yypt-0].node}
			}
		}
	case 7:
		//line dumper.y:60
		{
			yyVAL.node_list = append(yyS[yypt-2].node_list, yyS[yypt-0].node)
		}
	case 8:
		//line dumper.y:65
		{
			yyVAL.string = yyS[yypt-1].string
		}
	case 9:
		yyVAL.string = yyS[yypt-0].string
	case 10:
		//line dumper.y:71
		{
			yyVAL.string = yyS[yypt-1].string + "," + yyS[yypt-0].string
		}
	case 11:
		//line dumper.y:76
		{
			yyVAL.node = &BeautifyNode{
				Ptr:   strPtr(yyS[yypt-4].string),
				Type:  yyS[yypt-3].string,
				Value: strPtr(yyS[yypt-1].string),
			}
		}
	case 12:
		//line dumper.y:84
		{
			yyVAL.node = &BeautifyNode{
				Type:  yyS[yypt-3].string,
				Value: strPtr(yyS[yypt-1].string),
			}
		}
	case 13:
		//line dumper.y:91
		{
			yyVAL.node = &BeautifyNode{
				Ptr:  strPtr(yyS[yypt-1].string),
				Type: yyS[yypt-0].string,
			}
		}
	case 14:
		//line dumper.y:98
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
		//line dumper.y:109
		{
			yyVAL.node = &BeautifyNode{
				Ptr:          strPtr(yyS[yypt-4].string),
				Type:         yyS[yypt-3].string,
				StructValues: yyS[yypt-1].structKVs,
			}
		}
	case 19:
		//line dumper.y:117
		{
			yyVAL.node = &BeautifyNode{
				Type:         yyS[yypt-3].string,
				StructValues: yyS[yypt-1].structKVs,
			}
		}
	case 20:
		//line dumper.y:124
		{
			yyVAL.node = &BeautifyNode{
				Type: yyS[yypt-3].string,
				Ptr:  strPtr(yyS[yypt-1].string),
			}
		}
	case 21:
		//line dumper.y:132
		{
			if yyS[yypt-0].structKV == nil {
				yyVAL.structKVs = nil
			} else {
				yyVAL.structKVs = []*StructKV{yyS[yypt-0].structKV}
			}
		}
	case 22:
		//line dumper.y:140
		{
			yyVAL.structKVs = append(yyS[yypt-2].structKVs, yyS[yypt-0].structKV)
		}
	case 23:
		//line dumper.y:146
		{
			yyVAL.structKV = nil
		}
	case 24:
		//line dumper.y:150
		{
			yyVAL.structKV = &StructKV{yyS[yypt-2].string, yyS[yypt-0].node}
		}
	case 25:
		//line dumper.y:154
		{
			yyVAL.structKV = &StructKV{yyS[yypt-2].string, nil}
		}
	case 26:
		//line dumper.y:158
		{
			yyVAL.structKV = &StructKV{yyS[yypt-2].string, yyS[yypt-0].node}
		}
	case 27:
		//line dumper.y:163
		{
			yyVAL.node = &BeautifyNode{
				Ptr:         strPtr(yyS[yypt-7].string),
				Type:        "[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 28:
		//line dumper.y:171
		{
			yyVAL.node = &BeautifyNode{
				Type:        "[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 29:
		//line dumper.y:178
		{
			yyVAL.node = &BeautifyNode{
				Ptr:         strPtr(yyS[yypt-6].string),
				Type:        "[]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 30:
		//line dumper.y:186
		{
			yyVAL.node = &BeautifyNode{
				Type:        "[]" + yyS[yypt-3].string,
				ArrayValues: yyS[yypt-1].node_list,
			}
		}
	case 31:
		//line dumper.y:194
		{
			yyVAL.node = &BeautifyNode{
				Ptr:        strPtr(yyS[yypt-8].string),
				Type:       "map[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 32:
		//line dumper.y:202
		{
			yyVAL.node = &BeautifyNode{
				Type:       "map[" + yyS[yypt-5].string + "]" + yyS[yypt-3].string,
				HashValues: yyS[yypt-1].hashKVs,
			}
		}
	case 33:
		//line dumper.y:210
		{
			yyVAL.hashKVs = nil
		}
	case 34:
		//line dumper.y:214
		{
			yyVAL.hashKVs = []*HashKV{yyS[yypt-0].hashKV}
		}
	case 35:
		//line dumper.y:218
		{
			yyVAL.hashKVs = append(yyS[yypt-2].hashKVs, yyS[yypt-0].hashKV)
		}
	case 36:
		//line dumper.y:223
		{
			yyVAL.hashKV = &HashKV{yyS[yypt-2].node, yyS[yypt-0].node}
		}
	case 37:
		//line dumper.y:227
		{
			yyVAL.hashKV = &HashKV{yyS[yypt-2].node, nil}
		}
	}
	goto yystack /* stack new state and value */
}
