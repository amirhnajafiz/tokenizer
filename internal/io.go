package internal

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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
		parts := strings.Split(scanner.Text(), "/&/")
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
	return nil
}

// GetKeys in conf
func GetKeys() ([]string, error) {
	return nil, nil
}
