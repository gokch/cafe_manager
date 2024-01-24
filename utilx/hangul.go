package utilx

import (
	"unicode"
)

func IsHangulInitial(ch rune) bool {
	const (
		startRune = 0x3131 // 'ㄱ'
		endRune   = 0x314E // 'ㅎ'
	)

	return startRune <= ch && ch <= endRune
}

func IsHangulInitialsOnly(s string) bool {
	for _, ch := range s {
		if !IsHangulInitial(ch) {
			return false
		}
	}
	return true
}

func IsHangul(ch rune) bool {
	// 한글에 해당하는 Unicode 코드 포인트 범위
	return unicode.Is(unicode.Hangul, ch)
}

func IsHangulOnly(s string) bool {
	for _, ch := range s {
		if !IsHangul(ch) {
			return false
		}
	}
	return true
}
