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

func GetEnv(key string, dflt ...string) string {
	result, success := os.LookupEnv(key)
	if success {
		return result
	}
	if len(dflt) > 0 {
		return dflt[0]
	}
	return ""
}
