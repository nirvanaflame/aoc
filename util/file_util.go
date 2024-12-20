package util

import (
	"log"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	content := string(file)
	return strings.Split(content, "\n")
}
