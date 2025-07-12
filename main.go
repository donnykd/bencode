package main

import "fmt"

func decodeInt(s string) (string, error) {
	var lastIndex int
	for i := 1; i < len(s); i++ {
		if s[i] == 'e' {
			lastIndex = i
			break;
		}
	}
	if lastIndex == 0 {
		return "", fmt.Errorf("invalid encoded int %v", s)
	}
	return "", nil
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
