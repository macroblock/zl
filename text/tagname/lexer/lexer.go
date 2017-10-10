package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/macroblock/zl/core/zlog"
)

var log = zlog.Instance("lexer")

// TTagType -
type TTagType int

// TLexer -
type TLexer struct {
	start int
	pos   int
	width int
	src   string
	canRB bool
}

const runeEOF rune = 0

//New - create new lexer
func New(s string) *TLexer {
	return &TLexer{src: s}
}

//Next - go to next rune
func (o *TLexer) Next() rune {
	o.canRB = true
	if o.pos >= len(o.src) {
		o.width = 0
		return runeEOF
	}
	r, width := utf8.DecodeRuneInString(o.src[o.pos:])
	o.width = width
	o.pos += width
	log.Debug("---", fmt.Sprintf("%c", r))
	return r
}

//RollBack - reset pos
func (o *TLexer) RollBack() {
	if o.canRB {
		o.pos -= o.width
		o.canRB = false
	} else {
		log.Panic(nil, "Cannot RollBack")
	}
}

//Peek - look to next rune
func (o *TLexer) Peek() rune {
	r := o.Next()
	o.RollBack()
	return r
}

//Accept - peakup next rune
func (o *TLexer) Accept(s string) bool {
	if strings.IndexRune(s, o.Next()) >= 0 {
		return true
	}
	o.RollBack()
	return false
}

// AcceptPeek -
func (o *TLexer) AcceptPeek(s string) bool {
	if strings.IndexRune(s, o.Peek()) >= 0 {
		return true
	}
	return false
}

//AcceptFn - asdfasdf
func (o *TLexer) AcceptFn(fn func(r rune) bool) bool {
	if fn(o.Next()) {
		return true
	}
	o.RollBack()
	return false
}

//AcceptAnyEnglish -
func (o *TLexer) AcceptAnyEnglish() bool {
	if r := o.Next(); r >= 'a' && r <= 'z' {
		return true
	}
	o.RollBack()
	return false
}

//AcceptWhile -
func (o *TLexer) AcceptWhile(s string) {
	for strings.IndexRune(s, o.Next()) >= 0 {
	}
	o.RollBack()

}

// AcceptWhileNot -
func (o *TLexer) AcceptWhileNot(s string) {
	for strings.IndexRune(s, o.Next()) < 0 {
	}
	o.RollBack()
}

//Emit -
func (o *TLexer) Emit() string {
	// o.result = append(o.result, TTag{Type: tagid, Value: o.src[o.start:o.pos]})
	ret := o.src[o.start:o.pos]
	o.start = o.pos
	return ret
}

//Ignore -
func (o *TLexer) Ignore() {
	o.pos = o.start
	o.width = 0
}

//Pos - return pos of lexer
func (o *TLexer) Pos() int {
	return o.pos
}

//SetPos  - set pos to i
func (o *TLexer) SetPos(i int) {
	o.pos = i
	o.width = 0
}
