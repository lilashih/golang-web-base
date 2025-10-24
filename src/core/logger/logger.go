package logger

import (
	"fmt"
	"gbase/src/core/helper"
	"os"
	"path/filepath"
)

func getLogFile(name string, dirs ...string) (*os.File, error) {
	path, err := helper.GetStorageDirOrCreate(append([]string{"log"}, dirs...)...)
	if err != nil {
		return nil, err
	}

	path = filepath.Join(path, fmt.Sprintf("%s.log", name))
	newFile := !helper.IsFileExists(path)

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if newFile {
		file.Write([]byte{0xEF, 0xBB, 0xBF})
	}

	return file, err
}
