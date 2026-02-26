package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertBytes(bytes []byte) []byte {
	str := string(bytes)
	if strings.HasPrefix(str, ".") || strings.HasPrefix(str, "-") {
		return []byte(morse.ToText(str))
	}

	return []byte(morse.ToMorse(str))
}
