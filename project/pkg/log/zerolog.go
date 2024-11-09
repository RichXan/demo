package log

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"xproject/pkg/utils"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newZeroLogger(cfg LoggerConfig) zerolog.Logger {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.000"

	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		panic(fmt.Errorf("Fatal error, init logger error: Parse log level error %s \n", err.Error()))
	}
	zerolog.SetGlobalLevel(level)

	var writers []io.Writer
	writers = append(writers, os.Stdout)
	if !cfg.OnlyConsole {
		writers = append(writers, newRollingFile(cfg.Directory, cfg.ProjectName, cfg.LoggerName, cfg.MaxSize, cfg.MaxBackups))
	}
	mw := io.MultiWriter(writers...)
	l := zerolog.New(mw).With().Timestamp().Logger()
	return l
}

// 创建文件
func newRollingFile(dir, projectName, loggerName string, maxSize, maxBackups int) io.Writer {
	if dir == "" || projectName == "" || loggerName == "" {
		panic(fmt.Errorf("Fatal error, init logger error: log director or project name is nil \n"))
	}
	loggerNameLen := len(loggerName)
	if loggerName[loggerNameLen-4:loggerNameLen-1] != ".log" {
		loggerName = loggerName + ".log"
	}

	filename := path.Join(dir, projectName, loggerName)

	// make sure the log file permission is 644
	err := utils.SetFileModeWithCreating(filename, fs.FileMode(0644))
	if err != nil {
		panic(fmt.Errorf("Fatal error, init logger error: Set log file mode error %s \n", err))
	}

	return &lumberjack.Logger{
		Filename:   path.Join(dir, projectName, loggerName), //日志文件
		MaxBackups: maxBackups,                              //保留旧文件的最大数量
		MaxSize:    maxSize,                                 //单文件最大容量(单位MB)
		Compress:   false,
	}
}
