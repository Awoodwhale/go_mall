package utils

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

// Encryption
// @Description: AES 加密算法
type Encryption struct {
	key string
}

var Encrypt *Encryption

// init
// @Description: 初始化enc
func init() {
	Encrypt = NewEncryption()
}

// NewEncryption
// @Description: 获取enc对象
// @return *Encryption
func NewEncryption() *Encryption {
	return &Encryption{}
}

// PadPwd
// @Description: 填充密码长度
// @param srcByte []byte
// @param blockSize int
// @return []byte
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// unPadPwd
// @Description: 去掉填充的部分
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
// @receiver k *Encryption
// @param src string
// @return string
func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		return src
	}
	// 密码填充
	NewSrcByte := PadPwd(srcByte, block.BlockSize()) //由于字节长度不够，所以要进行字节的填充
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	// base64 编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

// AesDecoding
// @Description: aes解密
// @receiver k *Encryption
// @param pwd string
// @return string
func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte := []byte(pwd)
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd
	}
	block, errBlock := aes.NewCipher([]byte(k.key))
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

// SetKey
// @Description: 设置aes密钥
// @receiver k *Encryption
// @param key string
func (k *Encryption) SetKey(key string) {
	k.key = key
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
