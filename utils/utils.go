package utils

import (
	"encoding/base64"
)

func EncodeBase64(in string) (out string) {
	out = base64.StdEncoding.EncodeToString([]byte(in))
	return
}