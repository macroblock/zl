package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/macroblock/zl/core/log"
	"github.com/macroblock/zl/core/loglevel"
	"github.com/macroblock/zl/text"
)

var logFilter = loglevel.Error.OrLower() //loglevel.All

func main() {
	log := log.Default()
	log.AddLogger(logFilter, nil, os.Stdout)
	log.Debug("log initialized")
	if len(os.Args) <= 1 {
		log.Warning(nil, "not enough arguments")
	}
	for _, filename := range os.Args[1:] {
		dir, name := filepath.Split(filename)
		ext := filepath.Ext(filename)
		name = strings.TrimSuffix(name, ext)
		name, _ = text.Translit(name)
		err := os.Rename(filename, dir+name+ext)
		if err != nil {
			log.Error(err, "cannot rename file")
		}
	}
}
