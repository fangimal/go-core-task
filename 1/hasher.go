package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

type TypeValue struct {
	Value interface{}
	Type  string
}

func inspectVariables(vars ...interface{}) []TypeValue {
	result := make([]TypeValue, len(vars))
	for i, v := range vars {
		result[i] = TypeValue{
			Value: v,
			Type:  reflect.TypeOf(v).String(),
		}
	}

	return result
}

func valuesToStrings(infos []TypeValue) []string {
	strs := make([]string, len(infos))
	for i, v := range infos {
		strs[i] = fmt.Sprintf("%v", v.Value)
	}
	return strs
}

func stringToRunes(s string) []rune {
	return []rune(s)
}

func hashRunesWithSaltInMiddle(runes []rune, salt string) string {
	s := string(runes)
	mid := len(runes) / 2
	salted := s[:mid] + salt + s[mid:]

	hashBytes := sha256.Sum256([]byte(salted))

	return fmt.Sprintf("%x", hashBytes)
}
