package tagname

import "github.com/macroblock/zl/text/tagname/lexer"
import "strings"

// TParser -
type TParser struct {
	lexer   *lexer.TLexer
	taglist []TTag
}

// NewParser -
func NewParser(s string) *TParser {
	ret := &TParser{}
	ret.lexer = lexer.New(s)
	return ret
}

// TTestFn -
type TTestFn func() bool

// Emit -
func (o *TParser) Emit(tid TTagType) {
	o.taglist = append(o.taglist, TTag{Type: tid, Value: o.lexer.Emit()})
}

// Accept -
func (o *TParser) Accept(tagid TTagType, funcs ...TTestFn) bool {
	ok := true
	for _, fn := range funcs {
		ok = ok && fn()
	}

	if ok {
		o.Emit(tagid)
		return true
	}
	o.lexer.Ignore()
	return false
}

// Is -
func (o *TParser) Is(s string) TTestFn {
	return func() bool {
		return o.lexer.Accept(s)
	}
}

//Check -
func (o *TParser) Check(s string) TTestFn {
	return func() bool {
		return o.lexer.AcceptPeek(s)
	}

}

//IsDiggit -
func (o *TParser) IsDiggit() TTestFn {
	return func() bool {
		return o.lexer.AcceptFn(IsDiggit)
	}
}

// IsLetter -
func (o *TParser) IsLetter() TTestFn {
	return func() bool {
		return o.lexer.AcceptFn(IsEnglishLetter)
	}
}

// SubAccept -
func (o *TParser) SubAccept(funcs ...TTestFn) TTestFn {
	return func() bool {
		ok := true
		for ok {
			for _, fn := range funcs {
				if f := fn(); !f {
					return ok
				}
			}
		}

		return ok
	}
}

// WhileNotSeparator -
func (o *TParser) WhileNotSeparator() TTestFn {
	return func() bool {
		ok := false
		for r := o.lexer.Next(); IsEnglishLetter(r) || IsDiggit(r); r = o.lexer.Next() {
			ok = true
		}
		o.lexer.RollBack()
		return ok
	}
}

// WhileNot -
func (o *TParser) WhileNot(s string) TTestFn {
	return func() bool {
		for strings.IndexRune(s, o.lexer.Next()) < 0 {
		}
		o.lexer.RollBack()
		return true
	}
}

//IsEnglishLetter - asdfasdf
func IsEnglishLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// IsDiggit - check rune is diggit
func IsDiggit(r rune) bool {
	return r >= '0' && r <= '9'
}

// ParseA - audio
func (o *TParser) ParseA() bool {
	ok := o.lexer.Accept("a")
	for o.lexer.Peek() != '_' && o.lexer.Peek() != runeEOF && ok {
		// ok = ok && l.AcceptAnyEnglish()
		ok = ok && o.lexer.AcceptFn(IsEnglishLetter)
		log.Info(ok)
		ok = ok && o.lexer.Accept("12345678")
		log.Info(ok, "\n")
	}
	if ok {
		o.Emit(tagAudio)
		return true
	}
	o.lexer.Ignore()
	return false
}

// IsEndOfTag -
func (o *TParser) IsEndOfTag() bool {
	if o.lexer.Peek() == '_' || o.lexer.Peek() == runeEOF {
		return true
	}
	return false
}

// ParseUnknown -
func (o *TParser) ParseUnknown() {
	o.lexer.AcceptWhileNot(separators)
	o.Emit(tagUnknown)
}

//Result -
func (o *TParser) Result() []TTag {
	return o.taglist
}
