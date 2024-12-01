package helpers

import (
	"log"
	"os"
	"strings"
)

func ReadFile(name string) string {
	data, err := os.ReadFile("inputs/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func Lines(file string) []string {
	return strings.Split(strings.Trim(file, "\n"), "\n")
}
