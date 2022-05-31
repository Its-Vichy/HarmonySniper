package utils

import (
	"bufio"
	"os"
	"regexp"
)

func LoadTokensFromFile(FilePath string) []string {
	File, _ := os.Open(FilePath)
	scanner := bufio.NewScanner(File)
	var FoundTokens []string

	for scanner.Scan() {
		for _, Match := range regexp.MustCompile(`[\w-]{24}\.[\w-]{6}\.[\w-]{27}|mfa\.[\w-]{84}`).FindAllString(scanner.Text(), -1) {
			FoundTokens = append(FoundTokens, Match)
		}
	}

	return FoundTokens
}

func IncludeStr(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}