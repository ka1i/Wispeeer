package utils

import "strings"

func SafeFormat(origin string, spec string, join string, with string) string {
	spell := strings.Fields(origin)
	concat := strings.Join(spell, spec)
	newspell := []string{concat, join}
	r := strings.Join(newspell, with)
	return r
}
