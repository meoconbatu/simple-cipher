package cipher

import (
	"strings"
	"unicode"
)

const testVersion = 1

type Caesar struct {
	shiftAmount int
}
type Vigenere struct {
	key string
}

func NewCaesar() Cipher {
	c := Caesar{shiftAmount: 3}
	return c
}
func (c Caesar) Encode(in string) string {
	if len(in) == 0 {
		return ""
	}
	formattedStr := formatString(in)
	if len(formattedStr) == 0 {
		return ""
	}
	out := make([]byte, len(formattedStr))
	for i := 0; i < len(formattedStr); i++ {
		out[i] = formattedStr[i] + byte(c.shiftAmount)
		if out[i] >= 'a' {
			out[i] = ((out[i] - 'a') % 26) + 'a'
		} else {
			out[i] = byte(((int(out[i]) - 'a') % 26) + 123)
		}
	}
	return string(out)
}
func (c Caesar) Decode(in string) string {
	if len(in) == 0 {
		return ""
	}
	out := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[i] - byte(c.shiftAmount)
		if out[i] >= 'a' {
			out[i] = ((out[i] - 'a') % 26) + 'a'
		} else {
			out[i] = byte(((int(out[i]) - 'a') % 26) + 123)
		}
	}
	return string(out)
}

func NewShift(in int) Cipher {
	if !((in >= 1 && in <= 25) || (in <= -1 && in >= -25)) {
		return nil
	}
	c := Caesar{shiftAmount: in}
	return c
}
func NewVigenere(k string) Cipher {
	if !isValidVigenereKey(k) {
		return nil
	}
	c := Vigenere{key: k}
	return c
}
func isValidVigenereKey(key string) bool {
	if key == "" {
		return false
	}
	if key == strings.Repeat("a", len(key)) {
		return false
	}
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
func (c Vigenere) Encode(in string) string {
	if len(in) == 0 {
		return ""
	}
	formattedStr := formatString(in)
	if len(formattedStr) == 0 {
		return ""
	}
	k := generateKey(formattedStr, c.key)
	cipher_text := make([]byte, 0)
	for i := 0; i < len(formattedStr); i++ {
		x := (formattedStr[i] + k[i] - 2*'a') % byte(26)
		x += 'a'
		cipher_text = append(cipher_text, x)
	}
	return string(cipher_text)
}
func (c Vigenere) Decode(in string) string {
	out := make([]byte, 0)
	k := generateKey(in, c.key)
	for i := 0; i < len(in); i++ {
		x := (in[i] - k[i] + 26) % 26
		x += 'a'
		out = append(out, x)
	}
	return string(out)
}
func generateKey(str string, key string) string {
	x := len(key)
	out := make([]byte, 0)
	for i := 0; ; i++ {
		if x == i {
			i = 0
		}
		if len(out) == len(str) {
			break
		}
		out = append(out, key[i])
	}
	return string(out)
}
func formatString(str string) string {
	out := make([]rune, 0)
	for _, s := range str {
		if unicode.IsLetter(s) {
			out = append(out, unicode.ToLower(s))
		}
	}
	return string(out)
}
