package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const token = "//&&//"

// Set new value
func Set(key, value string) error {
	file, err := os.Open(baseFile)
	if err != nil {
		return ErrConfFileNotFound
	}

	list := make(map[string]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), token)

		tmpKey, _ := DeCode(parts[0])
		tmpValue, _ := DeCode(parts[1])

		list[tmpKey] = tmpValue
	}

	list[key] = value

	if er := scanner.Err(); er != nil {
		return ErrScanner
	}

	_ = file.Close()

	exportFile, err := os.OpenFile(baseFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return ErrConfFileNotFound
	}

	datawriter := bufio.NewWriter(exportFile)

	for data := range list {
		tmpKey, _ := Code(data)
		tmpValue, _ := Code(list[data])

		_, _ = datawriter.WriteString(fmt.Sprintf("%s%s%s\n", tmpKey, token, tmpValue))
	}

	_ = datawriter.Flush()
	_ = exportFile.Close()

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

		tmpKey, _ := DeCode(parts[0])
		tmpValue, _ := DeCode(parts[1])

		if tmpKey == key {
			return tmpValue, nil
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

		tmpKey, _ := DeCode(parts[0])
		tmpValue, _ := DeCode(parts[1])

		if tmpKey != key {
			list[tmpKey] = tmpValue
		}
	}

	if er := scanner.Err(); er != nil {
		return ErrScanner
	}

	_ = file.Close()

	exportFile, err := os.OpenFile(baseFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return ErrConfFileNotFound
	}

	datawriter := bufio.NewWriter(exportFile)

	for data := range list {
		tmpKey, _ := Code(data)
		tmpValue, _ := Code(list[data])

		_, _ = datawriter.WriteString(fmt.Sprintf("%s%s%s\n", tmpKey, token, tmpValue))
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

		tmpKey, _ := DeCode(parts[0])

		list = append(list, tmpKey)
	}

	if er := scanner.Err(); er != nil {
		return nil, ErrScanner
	}

	return list, nil
}
