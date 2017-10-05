package tagname

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/macroblock/zl/core/zlog"
)

// Tagname - struct tagname
type Tagname struct {
	name, terminator, mstring, form, ext, qw, a, sdhd string
	year, age                                         int
}

// TLexer -
type TLexer struct {
	start  int
	pos    int
	width  int
	src    string
	result []TTag
}

// TTag -
type TTag struct {
	Type  int
	Value string
}

const runeEOF rune = 0
const separators string = "_\x00"
const (
	tagUnknown = iota
	tagSeparator
	tagQ
	tagA
	tagD
	tagY
	tagN
	tagAge
	tagM
	tagT
	tagF
)

var list []string
var log = zlog.Instance("tagname")

//Something - do smthng
func Something() {
	log.Warning(nil, "Something")
}

//NewLexer - create new lexer
func NewLexer(s string) *TLexer {
	return &TLexer{src: s}
}

//Next - go to next rune
func (o *TLexer) Next() rune {
	if o.pos >= len(o.src) {
		o.width = 0
		return runeEOF
	}
	r, width := utf8.DecodeRuneInString(o.src[o.pos:])
	o.width = width
	o.pos += width
	return r
}

//RollBack - reset pos
func (o *TLexer) RollBack() {
	o.pos -= o.width
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
func (o *TLexer) Emit(tagid int) {
	o.result = append(o.result, TTag{Type: tagid, Value: o.src[o.start:o.pos]})
	o.start = o.pos
}

//Ignore -
func (o *TLexer) Ignore() {
	o.pos = o.start
	o.width = 0
}

//Result -
func (o *TLexer) Result() []TTag {
	return o.result
}

//ParseQ - quality
func ParseQ(lexer *TLexer) bool {
	ok := lexer.Accept("q")
	ok = ok && lexer.Accept("0123")
	ok = ok && lexer.Accept("sw")
	ok = ok && lexer.Accept("0123")
	if lexer.Peek() != '_' {
		ok = false
	}

	if ok {
		lexer.Emit(tagQ)
		return true
	}
	lexer.Ignore()
	return false
}

//ParseD - hdsd
func ParseD(lexer *TLexer) bool {
	ok := lexer.Accept("sh")
	ok = ok && lexer.Accept("d")
	if lexer.Peek() != '_' {
		ok = false
	}
	if ok {
		lexer.Emit(tagD)
		return true
	}
	lexer.Ignore()
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

//IsEnglishLetter - asdfasdf
func IsEnglishLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// ParseA - audio
func ParseA(lexer *TLexer) bool {

	ok := lexer.Accept("a")
	for lexer.Peek() != '_' && lexer.Peek() != runeEOF && ok {
		// ok = ok && lexer.AcceptAnyEnglish()
		ok = ok && lexer.AcceptFn(IsEnglishLetter)
		log.Info(ok)
		ok = ok && lexer.Accept("12345678")
		log.Info(ok, "\n")
	}
	if ok {
		lexer.Emit(tagA)
		return true
	}
	lexer.Ignore()
	return false
}

// func ParseT(lexer *TLexer) bool {
// 	ok := lexer.Accept("m")
// 	for lexer.Peek() != '_' || lexer.Peek() != '.' {
// 		ok = ok && lexer.AcceptAnyEnglish()
// 	}
// 	if ok {
// 		lexer.Emit(tagT)
// 		return true
// 	}
// 	lexer.Ignore()
// 	return false
// }

// ParseUnknown -
func ParseUnknown(lexer *TLexer) {
	lexer.AcceptWhileNot(separators)
	lexer.Emit(tagUnknown)
}

// Parse -
func Parse(s string) []TTag {
	lexer := NewLexer(s)

	for {
		ok := false
		lexer.AcceptWhile("_")
		lexer.Emit(tagSeparator)
		if lexer.Peek() == runeEOF {
			return lexer.Result()
		}

		switch lexer.Peek() {
		case 'q':
			ok = ParseQ(lexer)
		case 's', 'h':
			ok = ParseD(lexer)
		case 'a':
			ok = ParseA(lexer)
		} //end switch
		if !ok {
			ParseUnknown(lexer)
		}
		log.Warning(nil, fmt.Sprintf("%v %v", lexer.pos, lexer.Peek()))

	}
}

// New -
// func New(res string) *Tagname {
// var name, terminator, mstring, f string
// var qw, a, sdhd []byte
// var year, age int
// 	s := strings.Split(string(res), ".")
// 	original := strings.Split(s[0], "_")
// 	ext := s[1]

// 	if original[0] == "sd" || original[0] == "hd" {
// 		f = "rt"
// 	} else {
// 		f = "normal"
// 	}
// 	for i := range original {
// 		if strings.HasPrefix(original[i], "19") || strings.HasPrefix(original[i], "20") {
// 			year, _ = strconv.Atoi(original[i])
// 			log.Info(year)
// 			original[i] = ""
// 		}
// 		if strings.HasPrefix(original[i], "q") && strings.Contains(original[i], "w") {
// 			qw = []byte(original[i])
// 			log.Info("qw: " + string(qw))
// 			original[i] = ""
// 		}
// 		if strings.Contains(original[i], "trailer") || strings.Contains(original[i], "film") {
// 			terminator = original[i]
// 			log.Info(terminator)
// 			original[i] = ""
// 			if strings.HasPrefix(original[i-1], "m") {
// 				mstring = original[i-1]
// 				log.Info("mstring: " + mstring)
// 				original[i-1] = ""
// 			}
// 		}
// 		if original[i] >= strconv.Itoa(00) && original[i] <= strconv.Itoa(18) {
// 			age, _ = strconv.Atoi(original[i])
// 			log.Info(age)
// 			original[i] = ""
// 		}
// 		if strings.HasPrefix(original[i], "a") && strings.HasSuffix(original[i], "2") {
// 			a = []byte(original[i])
// 			log.Info("a: " + string(a))
// 			original[i] = ""
// 		}
// 		if strings.HasPrefix(original[i], "a") && strings.HasSuffix(original[i], "6") {
// 			a = []byte(original[i])
// 			log.Info("a: " + string(a))
// 			original[i] = ""
// 		}
// 	}
// 	name = strings.Join(original, " ")
// 	return &Tagname{name: name, form: f, qw: string(qw), terminator: terminator, year: year, mstring: mstring, age: age, a: string(a), sdhd: string(sdhd), ext: ext}
// }

// New -
// func New(res string) *Tagname {
// 	var name, terminator, mstring, f string
// 	var qw, a, sdhd []byte
// 	var year, age int
