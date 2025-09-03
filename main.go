package main

import (
	"fmt"
	"strconv"
)

func decodeInt(s string) (int, int, error) {
	// i4e
	var lastIndex int
	for i := 1; i < len(s); i++ {
		if s[i] == 'e' {
			lastIndex = i
			break
		}
	}
	if lastIndex == 0 {
		return 0, 0, fmt.Errorf("no end character found in s: %v", s)
	}
	result, err := strconv.Atoi(s[1:lastIndex])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse %v as integer", s)
	}
	return result, lastIndex + 1, nil
}

func decodeByteString(s string) ([]byte, int, error) {
	// 8:elephant
	var colonIndex int
	for i := 0; i < len(s); i++ {
		if s[i] == ':' {
			colonIndex = i
			break
		}
	}

	if colonIndex == 0 {
		return nil, 0, fmt.Errorf("no length found before colon in s: %v", s)
	}

	length, err := strconv.Atoi(s[:colonIndex])
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get length of s: %v", s)
	}

	byteString := make([]byte, length)

	copy(byteString, s[colonIndex+1:colonIndex+1+length])

	return byteString, colonIndex + 1 + length, nil
}

func decodeList(s string) (interface{}, int, error) {
	// li4e8:elephante
	var result []interface{}
	i := 1
	for i < len(s) && s[i] != 'e' {
		var consumed int
		currResult, consumed, err := decode(s[i:])
		if err != nil {
			return nil, 0, err
		}
		result = append(result, currResult)
		i += consumed
	}

	if i >= len(s) || s[i] != 'e' {
		return nil, 0, fmt.Errorf("list not terminated with 'e'")
	}

	return result, i + 1, nil
}

func decodeDict(s string) (string, error) {
	//
	return "", nil
}

func decode(s string) (interface{}, int, error) {
	switch firstChar := s[0]; {

	case firstChar == 'i':
		int, consumed, err := decodeInt(s)
		if err != nil {
			return nil, 0, fmt.Errorf("Cannot decode: %v", s)
		}
		return int, consumed, nil

	case firstChar > '0' && firstChar <= '9':
		str, consumed, err := decodeByteString(s)
		if err != nil {
			return nil, 0, fmt.Errorf("Cannot decode: %v", s)
		}
		return str, consumed, nil

	case firstChar == 'l':
		return decodeList(s)

	case firstChar == 'd':
		decodeDict(s)
	}

	return nil, 0, fmt.Errorf("Cannot decode: %v", s)
}

func main() {
	fmt.Println("hello, world")
}
