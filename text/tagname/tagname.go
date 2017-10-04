package tagname

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/macroblock/zl/core/zlog"
)

// Tagname - struct tagname
type Tagname struct {
	name, terminator, mstring, form, ext string
	qw, a, sdhd                          []byte
	year, age                            int
}

var list []string
var log = zlog.Instance("tagname")

//Something - do smthng
func Something() {
	log.Warning(nil, "Something")
}

// New -
func New(res string) *Tagname {
	var name, terminator, mstring, f string
	var qw, a, sdhd []byte
	var year, age int
	s := strings.Split(string(res), ".")
	original := strings.Split(s[0], "_")
	ext := s[1]

	if original[0] == "sd" || original[0] == "hd" {
		f = "rt"
	} else {
		f = "normal"
	}
	for i := range original {
		if strings.Contains(original[i], "19") || strings.Contains(original[i], "20") {
			year, _ = strconv.Atoi(original[i])
			log.Info(year)
			original[i] = ""
		}
		if strings.Contains(original[i], "q") && strings.Contains(original[i], "w") {
			qw = []byte(original[i])
			fmt.Printf("qw: %s\n", qw)
			original[i] = ""
		}
		if strings.Contains(original[i], "trailer") || strings.Contains(original[i], "film") {
			terminator = original[i]
			log.Info(terminator)
			original[i] = ""
			if strings.HasPrefix(original[i-1], "m") {
				mstring = original[i-1]
				log.Info("mstring: " + mstring)
				original[i-1] = ""
			}
		}
		if original[i] >= strconv.Itoa(00) && original[i] <= strconv.Itoa(18) {
			age, _ = strconv.Atoi(original[i])
			log.Info(age)
			original[i] = ""
		}
		if strings.HasPrefix(original[i], "a") && strings.HasSuffix(original[i], "2") {
			a = []byte(original[i])
			log.Info("a: " + string(a))
			original[i] = ""
		}
		if strings.HasPrefix(original[i], "a") && strings.HasSuffix(original[i], "6") {
			a = []byte(original[i])
			log.Info("a: " + string(a))
			original[i] = ""
		}
	}

	name = strings.Join(original, " ")
	return &Tagname{name: name, form: f, qw: qw, terminator: terminator, year: year, mstring: mstring, age: age, a: a, sdhd: sdhd, ext: ext}
}
