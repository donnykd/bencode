package main

import (
	"fmt"
	"strconv"
)

func decodeInt(s string) (int, error) {
	var lastIndex int
	for i := 1; i < len(s); i++ {
		if s[i] == 'e' {
			lastIndex = i
			break
		}
	}
	if lastIndex == 0 {
		return 0, fmt.Errorf("invalid encoded int %v", s)
	}
	result, err := strconv.Atoi(string(s[1:lastIndex]))
	if err != nil{
		return 0, fmt.Errorf("Failed to parse %v as integer", s)
	}
	return result, nil
}

func decodeByteString(s string) (string, error) {
	return "", nil
}

func decodeList(s string) (string, error) {
	return "", nil
}

func decodeDict(s string) (string, error) {
	return "", nil
}

func decode(s string) {
	switch firstChar := s[0]; {
	case firstChar == 'i':
		decodeInt(s)
	case firstChar > '0' && firstChar <= '9':
		decodeByteString(s)
	case firstChar == 'l':
		decodeList(s)
	case firstChar == 'd':
		decodeDict(s)
	}
}

func main() {
	fmt.Println("hello, world")
}
