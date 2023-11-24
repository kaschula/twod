package file

import (
	"os"
	"strings"
)

func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateDirIfNotExists(filePath string) {
	err := os.Mkdir(filePath, os.ModePerm)
	if err != nil {
		if !strings.ContainsAny(err.Error(), "file exists") {
			panic("wd error " + err.Error())
		}
	}
}
