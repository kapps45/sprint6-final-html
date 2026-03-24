package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// ConvertAuto автоматически определяет, Морзе или входящий текст,
// и конвертирует его.
//
// Принимает:
//   - input string — входная строка (Морзе или текст).
//
// Возвращает:
//   - string — результат конвертации;
//   - error — ошибка, если вход пустой или Морзе некорректен.
func ConvertAuto(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("empty input")
	}

	for _, ch := range input {
		if ch != '.' && ch != '-' && !unicode.IsSpace(ch) {
			return morse.ToMorse(input), nil
		}
	}

	text := morse.ToText(input)
	if strings.TrimSpace(text) == "" {
		return "", errors.New("invalid morse code")
	}
	return text, nil
}
