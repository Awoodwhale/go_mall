package service

import (
	"go_mall/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// MkDirIfNotExist
// @Description: 递归创建文件夹
// @param path string
// @return err error
func MkDirIfNotExist(path string) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	return
}

// UploadAvatarImage
// @Description: 上传头像至本地
// @param file multipart.File
// @param filename string
// @param uid uint
// @return filePath string
// @return err error
func UploadAvatarImage(file multipart.File, filename string, uid uint) (string, error) {
	// TODO: 补充七牛云的上传方式
	return uploadStatic(file, filename, uid, 1)
}

// UploadProductImage
// @Description: 上传产品图片到本地
// @param file multipart.File
// @param filename string
// @param uid uint
// @return string
// @return error
func UploadProductImage(file multipart.File, filename string, uid uint) (string, error) {
	// TODO: 补充七牛云的上传方式
	return uploadStatic(file, filename, uid, 2)
}

// UploadStatic
// @Description: 上传图片到本地，type为1表示上传avatar，为2表示上传product
// @param file multipart.File
// @param filename string
// @param uid uint
// @param typeof uint
func uploadStatic(file multipart.File, filename string, uid uint, typeof uint) (filePath string, err error) {
	var basePath, people string
	if typeof == 1 {
		people = "user"
		basePath = "." + conf.AvatarPath + people + strconv.Itoa(int(uid)) + "/"
	} else if typeof == 2 {
		people = "boss"
		basePath = "." + conf.ProductPath + people + strconv.Itoa(int(uid)) + "/"
	} else {
		return
	}

	if err = MkDirIfNotExist(basePath); err != nil {
		return
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return
	}
	err = os.WriteFile(basePath+filename, content, os.ModePerm)
	if err != nil {
		return
	}
	filePath = people + strconv.Itoa(int(uid)) + "/" + filename
	return
}
