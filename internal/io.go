package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const token = "/&/"

// Set new value
func Set(key, value string) error {
	return nil
}

// Get existing value
func Get(key string) (string, error) {
	file, err := os.Open(baseFile)
	if err != nil {
		return "", ErrConfFileNotFound
	}

	defer func(file *os.File) {
		er := file.Close()
		if er != nil {
			log.Println(er)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), token)
		if parts[0] == key {
			return parts[1], nil
		}
	}

	if er := scanner.Err(); er != nil {
		return "", ErrScanner
	}

	return "", ErrKeyNotSet
}

// Remove existing key
func Remove(key string) error {
	file, err := os.Open(baseFile)
	if err != nil {
		return ErrConfFileNotFound
	}

	list := make(map[string]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), token)

		if parts[0] != key {
			list[parts[0]] = parts[1]
		}
	}

	if er := scanner.Err(); er != nil {
		return ErrScanner
	}

	_ = file.Close()

	exportFile, err := os.OpenFile(baseFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return ErrConfFileNotFound
	}

	datawriter := bufio.NewWriter(exportFile)

	for data := range list {
		_, _ = datawriter.WriteString(fmt.Sprintf("%s%s%s\n", data, token, list[data]))
	}

	_ = datawriter.Flush()
	_ = exportFile.Close()

	return nil
}

// GetKeys in conf
func GetKeys() ([]string, error) {
	file, err := os.Open(baseFile)
	if err != nil {
		return nil, ErrConfFileNotFound
	}

	defer func(file *os.File) {
		er := file.Close()
		if er != nil {
			log.Println(er)
		}
	}(file)

	list := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), token)

		list = append(list, parts[0])
	}

	if er := scanner.Err(); er != nil {
		return nil, ErrScanner
	}

	return list, nil
}
