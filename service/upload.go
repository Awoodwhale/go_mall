package service

import (
	"go_mall/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

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

func MkDirIfNotExist(path string) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	return
}
