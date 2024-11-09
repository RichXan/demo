package log

type LoggerConfig struct {
	Level       string
	OnlyConsole bool
	Directory   string // log file path = Director + ProjectName + LoggerName + .log
	ProjectName string
	LoggerName  string
	MaxSize     int
	MaxBackups  int
}
