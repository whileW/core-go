package stringx

import (
	"bytes"
	"strings"
	"unicode"
)

func ToTitle(s string) string {
	if IsEmptyOrSpace(s) {
		return s
	}
	return strings.Title(s)
}

// ToCamel converts the input text into camel case
func ToCamel(s string) string {
	list := splitBy(s,func(r rune) bool {
		return r == '_'
	}, true)
	var target []string
	for _, item := range list {
		target = append(target, ToTitle(item))
	}
	return strings.Join(target, "")
}

// ToSnake converts the input text into snake case
func ToSnake(s string) string {
	list := splitBy(s,unicode.IsUpper, false)
	var target []string
	for _, item := range list {
		target = append(target, strings.ToLower(item))
	}
	return strings.Join(target, "_")
}

// it will not ignore spaces
func splitBy(s string,fn func(r rune) bool, remove bool) []string {
	if IsEmptyOrSpace(s) {
		return nil
	}
	var list []string
	buffer := new(bytes.Buffer)
	for _, r := range s {
		if fn(r) {
			if buffer.Len() != 0 {
				list = append(list, buffer.String())
				buffer.Reset()
			}
			if !remove {
				buffer.WriteRune(r)
			}
			continue
		}
		buffer.WriteRune(r)
	}
	if buffer.Len() != 0 {
		list = append(list, buffer.String())
	}
	return list
}

func IsEmptyOrSpace(s string) bool {
	if len(s) == 0 {
		return true
	}
	if strings.TrimSpace(s) == "" {
		return true
	}
	return false
}