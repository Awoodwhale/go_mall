package service

import (
	"go_mall/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarStatic(file multipart.File, filename string, uid uint) (filePath string, err error) {
	/**
	 * UploadAvatarStatic
	 * @Description: 上传头像至本地
	 * @param file
	 * @param filename
	 * @param uid
	 * @return filePath
	 * @return err
	 */
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
	/**
	 * MkDirIfNotExist
	 * @Description: 递归创建文件夹
	 * @param path
	 * @return err
	 */
	err = os.MkdirAll(path, os.ModePerm)
	return
}
