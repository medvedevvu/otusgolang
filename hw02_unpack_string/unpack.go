package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	ErrInvalidString         = errors.New("invalid string")
	twoDigitsInStringMatcher = regexp.MustCompile(`[^\d]+\d\d[^\d]*`)
	allNeedCombinations      = regexp.MustCompile(`\\*[^\d]\d`)
)

func Unpack(lstr string) (string, error) {
	// Все плохие вариенты строчек
	if _, e := strconv.Atoi(lstr); e == nil {
		return "", ErrInvalidString
	}
	r, _ := utf8.DecodeRuneInString(lstr)
	if unicode.IsDigit(r) {
		return "", ErrInvalidString
	}
	if twoDigitsInStringMatcher.MatchString(lstr) {
		return "", ErrInvalidString
	}
	// Все хорошие вариенты строчек
	allVarinats := allNeedCombinations.FindAll([]byte(lstr), -1)
	if len(allVarinats) == 0 {
		return lstr, nil
	}
	for _, v := range allVarinats {
		rebuildString(&lstr, string(v))
	}
	return lstr, nil
}

func rebuildString(inStr *string, pattern string) {
	if !strings.Contains(*inStr, pattern) {
		return
	}
	var replstring string
	switch len(pattern) {
	case 2: // SimbRep
		v := string(pattern[0])
		var k int
		k, _ = strconv.Atoi(string(pattern[1]))
		if k == 0 {
			replstring = ""
		} else {
			replstring = strings.Repeat(v, k)
		}
	case 3: // \\SimbRep
		v := string(pattern[1])
		var k int
		k, _ = strconv.Atoi(string(pattern[2]))
		if k == 0 {
			replstring = ""
		} else {
			replstring = strings.Repeat("\\"+v, k)
		}
	}
	*inStr = strings.ReplaceAll(*inStr, pattern, replstring)
}
