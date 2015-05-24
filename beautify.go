package dumper

import ()

type BeautifyNode struct {
	Ptr          *string
	Type         string
	Value        *string
	StructValues []*StructKV
	ArrayValues  []*BeautifyNode
	HashValues   []*HashKV
}

type StructKV struct {
	Key   string
	Value *BeautifyNode
}

type HashKV struct {
	Key   *BeautifyNode
	Value *BeautifyNode
}

func (n *BeautifyNode) String() string {
	return n.identString(0)
}

func (n *BeautifyNode) identString(ntabs uint) string {
	spaces := ""
	for i := uint(0); i < ntabs; i++ {
		spaces += "\t"
	}
	result := "{\n"

	if n.Ptr != nil {
		result += spaces + "\tPtr: " + *(n.Ptr) + "\n"
	}

	result += spaces + "\tType: " + n.Type + "\n"

	if n.Value != nil {
		result += spaces + "\tValue: " + *n.Value + "\n"
	}

	if n.StructValues != nil {
		result += spaces + "\tStructValues: {\n"
		for _, sv := range n.StructValues {
			result += spaces + "\t\t" + sv.Key + ": "
			if sv.Value == nil {
				result += "nil\n"
			} else {
				result += sv.Value.identString(ntabs+2) + "\n"
			}
		}
		result += spaces + "\t}\n"
	}

	if n.ArrayValues != nil {
		result += spaces + "\tArrayValues: {\n"
		for _, v := range n.ArrayValues {
			if v == nil {
				result += "nil\n"
			} else {
				result += spaces + "\t\t" + v.identString(ntabs+2) + "\n"
			}
		}
		result += spaces + "\t}\n"
	}

	if n.HashValues != nil {
		result += spaces + "\tHashValues: {\n"
		for _, v := range n.HashValues {
			result += spaces + "\t\t" + v.Key.identString(ntabs+2) + ": "
			if v.Value == nil {
				result += "nil\n"
			} else {
				result += v.Value.identString(ntabs+2) + "\n"
			}
		}
		result += spaces + "\t}\n"
	}

	return result + spaces + "}"
}

func GetBeautifyTree(dump []byte) (*BeautifyNode, error) {
	//yyDebug = 1
	lexer := &exprLex{line: dump}
	yyParse(lexer)
	if lexer.err == nil {
		return lexer.result, nil
	} else {
		return nil, lexer.err
	}
}
