package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
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

func BackupAndWrite(content, dst string) (int, error) {
	madeBackup, err := backupFile(dst)
	if err != nil {
		return 0, err
	}

	if madeBackup != "" {
		log.Printf("Backed up %s at %s", dst, madeBackup)
	}

	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.WriteString(content)
}

func BackupAndAppend(src, dst string) (int, error) {
	madeBackup, err := backupFile(dst)
	if err != nil {
		return 0, err
	}

	if madeBackup != "" {
		log.Printf("Backed up %s at %s", dst, madeBackup)
	}

	return Append(src, dst)
}

func backupFile(path string) (string, error) {
	ts := time.Now().UnixMilli()
	data, err := os.Stat(path)
	var backupPath string

	if !data.Mode().IsRegular() {
		return "", fmt.Errorf("%s is not a regular file", path)
	}

	// Check whether backup is needed or not
	if err == nil && data.Mode().IsRegular() {
		// Backup the file
		backupFileName := fmt.Sprintf("%s.%d.bak", path, ts)
		if _, err = Copy(path, backupFileName); err != nil {
			return "", fmt.Errorf("could not backup the file at %s: %s", path, err)
		}
		backupPath = backupFileName
	}

	// At this point, it is safe to override the file at path in any case
	return backupPath, nil
}
