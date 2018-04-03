package log

import (
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

// LoggerLineContextHook ...
// 为 logrus 增加行号显示
// https://github.com/sirupsen/logrus/issues/63
type LoggerLineContextHook struct{}

// Levels ...
func (hook LoggerLineContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire ...
// 每一次都要调用一次，性能开销较大
func (hook LoggerLineContextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/Sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["code_file"] = path.Base(file)
			entry.Data["code_func"] = path.Base(name)
			entry.Data["code_line"] = line
			break
		}
	}
	return nil
}

// InitLogrus ...
func InitLogrus(level string, debug bool) {
	logrus.SetLevel(logLevel(level))
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetOutput(os.Stdout)
	if debug {
		// only allow line record for debug model
		logrus.AddHook(LoggerLineContextHook{})
	}
}

func logLevel(level string) logrus.Level {
	l, err := logrus.ParseLevel(string(level))
	if err != nil {
		l = logrus.InfoLevel
		logrus.Warnf("error parsing level %q: %v, using %q	", level, err, l)
	}

	return l
}
