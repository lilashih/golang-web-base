package logger

import (
	"gbase/src/core/helper"
	"io"
	"log"
	"os"
)

var Migrate *log.Logger

func init() {
	if helper.IsTest() {
		// 直接寫入io.Discard，避免測試時產生不必要的log
		Migrate = log.New(io.Discard, "", log.LstdFlags|log.Lshortfile) // No-op logger for tests
	} else {
		date := helper.GetToday()
		file, err := getLogFile(("migrate-" + date))

		if err != nil {
			panic(err)
		}

		var writers []io.Writer
		writers = append(writers, file)      // file log
		writers = append(writers, os.Stdout) // print log

		multiWriter := io.MultiWriter(writers...)
		Migrate = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)
	}
}
