package common

import (
	"os"
	"crypto/md5"
	"io"
)

func ComputeMD5(file *os.File) []byte {
	var result []byte
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result
	}
	return hash.Sum(result)
}
