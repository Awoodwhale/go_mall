package utils

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

// padPwd
// @Description: 填充密码长度
// @param srcByte []byte
// @param blockSize int
// @return []byte
func padPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// unPadPwd
// @Description: 掉填充的部分
// @param dst []byte
// @return []byte
// @return error
func unPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}
	// 去掉的长度
	unpadNum := int(dst[len(dst)-1])
	strErr := "error"
	op := []byte(strErr)
	if len(dst) < unpadNum {
		return op, nil
	}
	str := dst[:(len(dst) - unpadNum)]
	return str, nil
}

// AesEncoding
// @Description: aes加密
// @param src string
// @param key string
// @return string
func AesEncoding(src string, key string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return src
	}
	// 密码填充
	NewSrcByte := padPwd(srcByte, block.BlockSize()) //由于字节长度不够，所以要进行字节的填充
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	// base64 编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

// AesDecoding
// @Description: aes解密
// @param pwd string
// @param key string
// @return string
func AesDecoding(pwd string, key string) string {
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd
	}
	block, errBlock := aes.NewCipher([]byte(key))
	if errBlock != nil {
		return pwd
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, err = unPadPwd(dst) // 填充的要去掉
	if err != nil {
		return pwd
	}
	return string(dst)
}

// CheckKey
// @Description: 检测Key的合法性
// @param key string
// @return bool
func CheckKey(key string) bool {
	if key == "" || len(key) != 16 {
		return false
	}
	return true
}
