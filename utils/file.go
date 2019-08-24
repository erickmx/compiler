package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

// IsValidFile verify if the file extension is valid
func IsValidFile(file string) bool {
	matched, err := regexp.MatchString("(?im)^\\w+.cod$", file)
	if err != nil {
		return false
	}
	return matched
}

// OpenFile open the gived file and returns it to work with it
func OpenFile(path string) (string, error) {
	base := filepath.Base(path)
	if !IsValidFile(base) {
		return "", errors.New("The file extension is not valid")
	}
	fmt.Println(base)
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	byteContent, err := ioutil.ReadFile(absPath)
	return string(byteContent), err
}
