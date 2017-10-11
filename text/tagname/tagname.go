package tagname

import (
	"github.com/macroblock/zl/core/zlog"
)

var log = zlog.Instance("tagname")

// Tagname - struct tagname
type Tagname struct {
	name, terminator, mstring, form, ext, qw, a, sdhd string
	year, age                                         int
}

const (
	tagUnknown = iota
	tagSeparator
	tagQuality
	tagAudio
	tagVideoDefinition
	tagYear
	tagAge
	tagMeta
	tagTerminator
	tagName
)

func (o TTagType) String() string {
	switch o {
	case tagUnknown:
		return "Unknown"
	case tagSeparator:
		return "Separator"
	case tagQuality:
		return "Quality"
	case tagAudio:
		return "Audio"
	case tagVideoDefinition:
		return "VideoDefinition"
	case tagYear:
		return "Year"
	case tagAge:
		return "Age"
	case tagMeta:
		return "Meta"
	case tagTerminator:
		return "Termination"
	case tagName:
		return "Name"
	}
	return "Invalid"
}

// TTagType -
type TTagType int

// TTag -
type TTag struct {
	Type  TTagType
	Value string
}

const separators string = "_\x00\x2E"
const runeEOF rune = 0

var list []string

// Parse -
func Parse(s string) []TTag {
	p := NewParser(s)
	for {
		ok := false
		p.lexer.AcceptWhile("_.")
		p.Emit(tagSeparator)
		if p.lexer.Peek() == runeEOF {
			return p.Result()
		}
		switch p.lexer.Peek() {
		case '0', '1', '3', '6':
			// ok = p.Accept(tagAge, p.SubAccept(p.IsDiggit(), p.IsDiggit()), p.Check(separators))
			ok = p.Accept(tagAge, p.OrAccept(p.MustIs("00"), p.MustIs("03"), p.MustIs("06"), p.MustIs("12"), p.MustIs("16"), p.MustIs("18")), p.Check(separators))
		} //end switch

		switch p.lexer.Peek() {
		case 'q':
			ok = p.Accept(tagQuality, p.Is("q"), p.Is("0123"), p.Is("ws"), p.Is("0123"), p.Check(separators))
		case 's', 'h':
			ok = p.Accept(tagVideoDefinition, p.Is("sh"), p.Is("d"), p.Check(separators))
		case 'a':
			ok = p.Accept(tagAudio, p.Is("a"), p.SubAccept(p.Is("er"), p.Is("26")), p.Is("n"), p.Check(separators))
		case 'm':
			ok = p.Accept(tagMeta, p.Is("m"), p.WhileNotSeparator(), p.Check(separators))
		case '1', '2':
			ok = p.Accept(tagYear, p.Is("12"), p.IsDiggit(), p.IsDiggit(), p.IsDiggit(), p.Check(separators))
		case 'f':
			ok = p.Accept(tagTerminator, p.MustIs("film"), p.Check(separators))
		case 't':
			ok = p.Accept(tagTerminator, p.MustIs("trailer"), p.Check(separators))
			// default:
			// 	ok = p.Accept(tagName, p.SubAccept(p.WhileNotSeparator()), p.Check(separators))
		} //end switch

		if !ok {
			ok = p.Accept(tagUnknown, p.WhileNot(separators))
		}
		// log.Warning(nil, fmt.Sprintf("%v %v", l.pos, l.Peek()))
	}
}

// func ParseT(l *TLexer) bool {
// 	ok := l.Accept("m")
// for l.Peek() != '_' || l.Peek() != '.' {
// 	ok = ok && l.AcceptAnyEnglish()
// }
// if ok {
// 	l.Emit(tagT)
// 	return true
// }
// 	l.Ignore()
// return false
// }
