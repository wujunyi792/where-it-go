package logger

import (
	"github.com/wujunyi792/gin-template-new/config"
	"io"
	"log"
	"os"
)

type debugDefault struct {
	Debug *log.Logger
}

func (d *debugDefault) Println(v ...interface{}) {
	if config.GetConfig().MODE == "debug" {
		d.Debug.Println(v)
	}
}

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Debug   *debugDefault
)

func init() {
	errFile, err := os.OpenFile(ERROR_LOG_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败！")
	}

	Info = log.New(os.Stdout, "[Info] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "[Warning] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "[Error] ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = &debugDefault{
		Debug: log.New(os.Stdout, "[Debug] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
