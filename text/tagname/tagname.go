package tagname

import (
	"github.com/macroblock/zl/core/zlog"
	"github.com/macroblock/zl/text/tagname/lexer"
)

var log = zlog.Instance("tagname")

// Tagname - struct tagname
type Tagname struct {
	name, terminator, mstring, form, ext, qw, a, sdhd string
	year, age                                         int
}

func (o *TTagType) String() string {
	switch *o {
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

// TTagType -
type TTagType int

const separators string = "_\x00"
const runeEOF rune = 0

var list []string

//Something - do smthng
func Something() {
	log.Warning(nil, "Something")
}

// ParseA - audio
func ParseA(l *lexer.TLexer) bool {
	ok := l.Accept("a")
	for l.Peek() != '_' && l.Peek() != runeEOF && ok {
		// ok = ok && l.AcceptAnyEnglish()
		ok = ok && l.AcceptFn(IsEnglishLetter)
		log.Info(ok)
		ok = ok && l.Accept("12345678")
		log.Info(ok, "\n")
	}
	if ok {
		l.Emit(tagAudio)
		return true
	}
	l.Ignore()
	return false
}

//ParseQ - quality
func ParseQ(l *lexer.TLexer) bool {
	ok := l.Accept("q")
	ok = ok && l.Accept("0123")
	ok = ok && l.Accept("sw")
	ok = ok && l.Accept("0123")
	ok = ok && IsEndOfTag(l)
	if ok {
		l.Emit(tagQuality)
		return true
	}
	l.Ignore()
	return false
}

//ParseD - hdsd
func ParseD(l *lexer.TLexer) bool {
	ok := l.Accept("sh")
	ok = ok && l.Accept("d")
	ok = ok && IsEndOfTag(l)
	if ok {
		l.Emit(tagVideoDefinition)
		return true
	}
	l.Ignore()
	return false
}

//ParseM -
func ParseM(l *lexer.TLexer) bool {
	ok := l.Accept("m")
	for l.Peek() != '_' && l.Peek() != runeEOF && ok {
		ok = ok && l.AcceptFn(IsEnglishLetter)
	}
	if ok {
		l.Emit(tagMeta)
		return true
	}
	l.Ignore()
	return false
}

// ParseY -
func ParseY(l *lexer.TLexer) bool {
	ok := l.Accept("12")
	ok = ok && l.AcceptFn(IsDiggit)
	ok = ok && l.AcceptFn(IsDiggit)
	ok = ok && l.AcceptFn(IsDiggit)
	ok = ok && IsEndOfTag(l)
	if ok {
		l.Emit(tagYear)
		return true
	}
	l.Ignore()
	return false
}

//IsEnglishLetter - asdfasdf
func IsEnglishLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// IsDiggit - check rune is diggit
func IsDiggit(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsEndOfTag -
func IsEndOfTag(l *lexer.TLexer) bool {
	if l.Peek() == '_' || l.Peek() == runeEOF {
		return true
	}
	return false
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

// ParseUnknown -
func ParseUnknown(l *lexer.TLexer) {
	l.AcceptWhileNot(separators)
	l.Emit(tagUnknown)
}

// Parse -
func Parse(s string) []lexer.TTag {
	l := lexer.New(s)
	for {
		ok := false
		l.AcceptWhile("_")
		l.Emit(tagSeparator)
		if l.Peek() == runeEOF {
			return l.Result()
		}

		switch l.Peek() {
		case 'q':
			ok = ParseQ(l)
		case 's', 'h':
			ok = ParseD(l)
		case 'a':
			ok = ParseA(l)
		case 'm':
			ok = ParseM(l)
		case '1', '2':
			ok = ParseY(l)
		} //end switch
		if !ok {
			ParseUnknown(l)
		}
		// log.Warning(nil, fmt.Sprintf("%v %v", l.pos, l.Peek()))

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
