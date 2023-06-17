package handler

import (
	"regexp"
	"unicode/utf8"
)

const LENGTH_LIMIT = 100

var (
	ikaRegexp   = regexp.MustCompile(`(い|イ|ｲ)(か|カ|ｶ)`)
	shikaRegexp = regexp.MustCompile(`(し|シ|ｼ)(か|カ|ｶ)`)
	mekaRegexp  = regexp.MustCompile(`(め|メ|ﾒ)(か|カ|ｶ)`)
)

func checkIkaShikaMeka(content string, reg *regexp.Regexp) bool {
	return reg.MatchString(content)
}

func checkLength(content string) bool {
	return utf8.RuneCountInString(content) <= LENGTH_LIMIT
}
