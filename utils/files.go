package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

// GetSha1ViaPath 通过文件路径 filePath 来获取文件的 sha1
func GetSha1ViaPath(filePath string) (sha1Str string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return sha1Str, err
	}
	hashInBytes := hash.Sum(nil)[:20]
	sha1Str = hex.EncodeToString(hashInBytes)

	return
}

// GetSha1ViaFile 通过文件对象 file 来获取文件的 sha1
func GetSha1ViaFile(file *os.File) (sha1Str string, err error) {
	file.Seek(0, 0)
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return sha1Str, err
	}
	hashInBytes := hash.Sum(nil)[:20]
	sha1Str = hex.EncodeToString(hashInBytes)

	return
}
