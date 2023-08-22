package helpers

import (
	"net/http"
	"strings"
)

func UpperCase(s string) string {
	return strings.ToUpper(s)
}

func LowerCase(s string) string {
	return strings.ToLower(s)
}

func FullName(n, s string) string {
	return s + n
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
