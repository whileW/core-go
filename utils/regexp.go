package utils

import "regexp"

//正则表达式支持

func FindAllString(exp string,str string) []string {
	re :=regexp.MustCompile(exp)
	return re.FindAllString(str,-1)
}
func FindString(exp string,str string) string {
	re :=regexp.MustCompile(exp)
	return re.FindString(str)
}
func MatchString(exp string,str string) bool {
	re :=regexp.MustCompile(exp)
	return re.MatchString(str)
}
func MatchPhone(phone string) bool {
	exp := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	return MatchString(exp,phone)
}
func MatchEmail(email string) bool {
	exp := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	return MatchString(exp,email)
}