package day01b

import (
	"unicode"
	"unicode/utf8"
)

// EvalCaptchaMid evaluates the result of the passed captcha string
func EvalCaptchaMid(captcha string) int {
	len := utf8.RuneCountInString(captcha)
	var runeList []rune
	runeList = make([]rune, len+1, len+1)
	pos := 0
	for _, rune := range captcha {
		runeList[pos] = rune
		pos++
	}
	runeList[len] = runeList[0]
	sum := 0
	for i := 0; i < len; i++ {
		if unicode.IsDigit(runeList[i]) && (runeList[i] == runeList[(i+(len/2))%len]) {
			sum += int(runeList[i] - '0')
		}
	}
	return sum
}
