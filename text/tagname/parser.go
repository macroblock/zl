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
	if CheckSeq(funcs...) {
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

//MustIs - must compare all rune
func (o *TParser) MustIs(s string) TTestFn {
	return func() bool {
		// pos := o.lexer.Pos()
		// log.Info("pos ", pos)
		// o.WhileNotSeparator()
		// pos1 := o.lexer.Pos()
		// log.Info("pos1 ", pos1)
		// if pos != pos1 {
		// 	return false
		// }

		for i := range s {
			log.Info(s[i])
			if !o.lexer.Accept(string(s[i])) {
				return false
			}
		}
		return true
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

// CheckSeq -
func CheckSeq(funcs ...TTestFn) bool {
	for _, fn := range funcs {
		if !fn() {
			return false
		}
	}
	return true
}

// SubAccept - accept if all fn true
func (o *TParser) SubAccept(funcs ...TTestFn) TTestFn {
	return func() bool {
		ok := false
		for {
			p := o.lexer.Pos()
			if !CheckSeq(funcs...) {
				o.lexer.SetPos(p)
				return ok
			}
			ok = true
		}
	}
}

//OrAccept - accept if any fn true
func (o *TParser) OrAccept(funcs ...TTestFn) TTestFn {
	return func() bool {
		for _, fn := range funcs {
			f := fn()
			if f {
				log.Info(f)
				return true
			}
		}
		log.Info("false")
		return false
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

// func (o *TParser) ParseName() {
// 	o.lexer.AcceptWhileNot(separators)
// 	o.Emit(tagName)
// }

//Result -
func (o *TParser) Result() []TTag {
	return o.taglist
}
