package common

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// ComputeFileMD5 hash sum for file
func ComputeFileMD5(file *os.File) []byte {
	var result []byte
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result
	}
	return hash.Sum(result)
}

// ConvertToMD5 receive MD5 sring
func ConvertToMD5(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	hash := hasher.Sum(nil)
	inMD5 := hex.EncodeToString(hash)
	return inMD5
}
