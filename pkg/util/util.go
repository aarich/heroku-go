package util

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func GetEnv(varName string) string {
	v := os.Getenv(varName)

	if v == "" {
		log.Fatalf("$%s must be set", varName)
	}

	return v
}
