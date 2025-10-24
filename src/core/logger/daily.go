package logger

import (
	"gbase/src/core/helper"
	"io"
	"log"
	"os"
)

var Log *log.Logger

func init() {
	date := helper.GetToday()
	file, err := getLogFile(("log-" + date))

	if err != nil {
		panic(err)
	}

	var writers []io.Writer
	writers = append(writers, file) // file log

	if !helper.IsRelease() {
		writers = append(writers, os.Stdout) // console log
	}

	multiWriter := io.MultiWriter(writers...)
	Log = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)
}
