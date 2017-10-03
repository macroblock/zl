package tagname

import "github.com/macroblock/zl/core/zlog"
import "strings"

// Tagname - struct tagname
type Tagname struct {
	original                        []string
	name, sdhd, terminator, mstring string
	year, age, w, h                 int
	qw, a                           []rune
}

var list []string
var log = zlog.Instance("tagname")

//Something - do smthng
func Something() {
	log.Warning(nil, "Something")
}

// New -
func New(res string) *Tagname {
	// &Tagname{original: strings.Split(res, "_")}
	return &Tagname{original: strings.Split(res, "_")}
}
