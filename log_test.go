package logger

import (
	"testing"
	"time"
)

func initLogger(name, logPath, logName string, level string) (err error) {
	m := make(map[string]string, 8)
	m["log_path"] = logPath
	m["log_name"] = logName
	m["log_level"] = level
	m["log_split_type"] = "size"
	err = InitLogger(name, m)
	if err != nil {
		return
	}
	Debug("init logger success")
	return
}

func Run() {
	for {
		Debug("user server is running, :\\project\\src\\github.com\\pingguoxueyuan\\gostudy\\listen17\\user_server")
		time.Sleep(time.Second)
	}
}

func TestLog(t *testing.T) {
	initLogger("file", "c:/logs/", "user_server", "debug")
	Run()
	return
}
