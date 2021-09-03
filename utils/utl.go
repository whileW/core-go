package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//三元运算符
func IF(condition bool,trueval,falseval interface{}) interface{} {
	if condition {
		return trueval
	}
	return falseval
}

//slice contains
func SliceStringContains(a []string,b string) bool {
	for _,t := range a {
		if t == b {
			return true
		}
	}
	return false
}

func Md5v(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}