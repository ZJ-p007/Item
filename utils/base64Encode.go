package utils

import "encoding/base64"

func Base64Str(mes string) string {
	return base64.StdEncoding.EncodeToString([]byte(mes))
}
