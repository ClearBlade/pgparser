/*
PostgreSQL grammar.
The MIT License (MIT).
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package parser

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type PostgreSQLParserBase struct {
	*antlr.BaseParser
}

func NewPostgreSQLParserBase(input antlr.TokenStream) *PostgreSQLParserBase {
	return &PostgreSQLParserBase{
		BaseParser: antlr.NewBaseParser(input),
	}
}

func (receiver *PostgreSQLParserBase) ParseRoutineBody() {
	// localContext, ok := localContextInterface.(*Createfunc_opt_listContext)
	//
	//	if !ok {
	//	    return
	//	}
	//
	// var lang string
	//
	//	for _, coi := range localContext.AllCreatefunc_opt_item() {
	//	    createFuncOptItemContext, ok := coi.(*Createfunc_opt_itemContext)
	//	    if !ok || createFuncOptItemContext.LANGUAGE() == nil {
	//	        continue
	//	    }
	//	    nonReservedWordOrSConstContextInterface := createFuncOptItemContext.Nonreservedword_or_sconst()
	//	    if nonReservedWordOrSConstContextInterface == nil {
	//	        continue
	//	    }
	//	    nonReservedWordOrSConstContext, ok := nonReservedWordOrSConstContextInterface.(*Nonreservedword_or_sconstContext)
	//	    if !ok {
	//	        continue
	//	    }
	//	    nonReservedWordContextInterface := nonReservedWordOrSConstContext.Nonreservedword()
	//	    if nonReservedWordContextInterface == nil {
	//	        continue
	//	    }
	//	    nonReservedWordContext, ok := nonReservedWordContextInterface.(*NonreservedwordContext)
	//	    if !ok {
	//	        continue
	//	    }
	//	    identifierInterface := nonReservedWordContext.Identifier()
	//	    if identifierInterface == nil {
	//	        continue
	//	    }
	//	    identifier, ok := identifierInterface.(*IdentifierContext)
	//	    if !ok {
	//	        continue
	//	    }
	//	    node := identifier.Identifier()
	//	    if node == nil {
	//	        continue
	//	    }
	//	    lang = node.GetText()
	//	    break
	//	}
	//
	//	if lang == "" {
	//	    return
	//	}
	//
	// var funcAs *Createfunc_opt_itemContext
	//
	//	for _, coi := range localContext.AllCreatefunc_opt_item() {
	//	    ctx, ok := coi.(*Createfunc_opt_itemContext)
	//	    if !ok || ctx.LANGUAGE() == nil {
	//	        continue
	//	    }
	//	    as := ctx.Func_as()
	//	    if as != nil {
	//	        funcAs = ctx
	//	        break
	//	    }
	//	}
	//
	//	if funcAs == nil {
	//	    return
	//	}
	//
	// funcAsContextInterface := funcAs.Func_as()
	//
	//	if funcAsContextInterface == nil {
	//	    return
	//	}
	//
	// funcAsContext, ok := funcAsContextInterface.(*Func_asContext)
	//
	//	if !ok {
	//	    return
	//	}
	//
	// sConstContextInterface := funcAsContext.Sconst(0)
	//
	//	if sConstContextInterface == nil {
	//	    return
	//	}
	//
	// sConstContext, ok := sConstContextInterface.(*SconstContext)
	//
	//	if !ok {
	//	    return
	//	}
	//
	// text := GetRoutineBodyString(sConstContext)
	// line := sConstContext.GetStart().GetLine()
	// parser := getPostgreSQLParser(text)
	// switch lang {
	// case "plpgsql":
	//
	//	funcAs.Func_as().(*Func_asContext).Definition = parser.Plsqlroot()
	//
	// case "sql":
	//
	//	    funcAs.Func_as().(*Func_asContext).Definition = parser.Root()
	//	}
	//
	//	for _, err := range parser.parseErrors {
	//	    receiver.parseErrors = append(receiver.parseErrors, &PostgreSQLParseError{
	//	        Number:  err.Number,
	//	        Offset:  err.Offset,
	//	        Line:    err.Line + line,
	//	        Column:  err.Column,
	//	        Message: err.Message,
	//	    })
	//	}
}

func TrimQuotes(s string) string {
	if s == "" {
		return s
	}
	return s[1 : len(s)-2]
}

func unquote(s string) string {
	result := strings.Builder{}
	length := len(s)
	index := 0
	for index < length {
		c := s[index]
		result.WriteByte(c)
		if c == '\'' && index < length-1 && (s[index+1] == '\'') {
			index++
		}
		index++
	}
	return result.String()
}

func GetRoutineBodyString(rule *SconstContext) string {
	if rule.Anysconst() == nil {
		return ""
	}
	anySConstContext := rule.Anysconst().(*AnysconstContext)

	stringConstant := anySConstContext.StringConstant()
	if stringConstant != nil {
		return unquote(TrimQuotes(stringConstant.GetText()))
	}

	unicodeEscapeStringConstant := anySConstContext.UnicodeEscapeStringConstant()
	if unicodeEscapeStringConstant != nil {
		return TrimQuotes(unicodeEscapeStringConstant.GetText())
	}

	escapeStringConstant := anySConstContext.EscapeStringConstant()
	if escapeStringConstant != nil {
		return TrimQuotes(escapeStringConstant.GetText())
	}

	result := strings.Builder{}
	for _, node := range anySConstContext.AllDollarText() {
		result.WriteString(node.GetText())
	}
	return result.String()
}

func (p *PostgreSQLParserBase) OnlyAcceptableOps() bool {
	stream := p.GetTokenStream()
	c := stream.LT(1)
	text := c.GetText()
	return text == "!" || text == "!!" || text == "!=-"
}
