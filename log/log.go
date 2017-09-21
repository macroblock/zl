package log

// TLogLevel -
type TLogLevel uint

// General loglevel flags
const (
	Quiet TLogLevel = 0
	Fatal           = 1 << (iota * 2)
	Error
	Warning
	Notice
	Info
	Debug
)

// TLogger -
type TLogger struct {
	logger ILog
}

func (o *TLogger) Close() {
}

func (o *TLogger) Log(level TLogLevel, err error, text ...interface{}) {
}

// ILog -
type ILog interface {
	Log(level TLogLevel, err error, text ...interface{})
}

type TLog struct {
}
