package util

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

func YamlLoad(path string, source interface{}) error {
	if ok := FileExist(path); ok {
		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(file, source)
		if err != nil {
			return err
		}
	}
	return errors.New("not found file")
}

func JsonLoad(path string, source interface{}) error {
	return nil
}

func FileExist(path string) bool {
	return false
}

func CurrentAbsPath() string {
	return ""
}
