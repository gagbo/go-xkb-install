package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func Append(src, dst string) (int, error) {
	f, err := os.OpenFile(dst, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	fileBytes, err := ioutil.ReadAll(source)
	if err != nil {
		return 0, err
	}

	return f.Write(fileBytes)
}
