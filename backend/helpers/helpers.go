package helpers

import (
	"os"
	b64 "encoding/base64"
)

func Getb64() string {
	strToEncode := os.Getenv("ENV_ZINC_USER") +":"+os.Getenv("ENV_ZINC_PASSWORD") 
	strEncoded := b64.StdEncoding.EncodeToString([]byte(strToEncode)) 
	return strEncoded
}

func GetUrlZinc() string {
	return os.Getenv("ENV_ZINC_URL")
}