package tagname

import "github.com/macroblock/zl/core/zlog"

var log = zlog.Instance("tagname")

func Something() {
	log.Warning(nil, "Something")
}
