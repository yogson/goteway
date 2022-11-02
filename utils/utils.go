package utils

import (
	"log"
	"os"
)

func ReadFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
