package service

import (
	"go_mall/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// UploadAvatarStatic
// @Description: 上传头像至本地
// @param file multipart.File
// @param filename string
// @param uid uint
// @return filePath string
// @return err error
func UploadAvatarStatic(file multipart.File, filename string, uid uint) (filePath string, err error) {
	basePath := "." + conf.AvatarPath + "user" + strconv.Itoa(int(uid)) + "/"
	if err = MkDirIfNotExist(basePath); err != nil {
		return
	}
	avatarPath := basePath + filename
	content, err := io.ReadAll(file)
	if err != nil {
		return
	}
	err = os.WriteFile(avatarPath, content, os.ModePerm)
	if err != nil {
		return
	}
	filePath = "user" + strconv.Itoa(int(uid)) + "/" + filename
	return
}

// MkDirIfNotExist
// @Description: 递归创建文件夹
// @param path string
// @return err error
func MkDirIfNotExist(path string) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	return
}
