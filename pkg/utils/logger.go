package utils

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var Logger *logrus.Logger

// InitLog
// @Description: logger init
func InitLog(logPath string) {
	// 存在就去更新文件
	if Logger != nil {
		src, err := setOutPutFile(logPath)
		if err != nil {
			panic(err)
		}
		Logger.Out = src
		return
	}
	// 不存在就实例化
	logger := logrus.New()
	src, err := setOutPutFile(logPath)
	if err != nil {
		panic(err)
	}
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// TODO: elk体系
	Logger = logger
}

// setOutPutFile
// @Description: 获取输出log的文件
// @return src *os.File
// @return err error
func setOutPutFile(logPath string) (src *os.File, err error) {
	var logFilePath string
	// 获取log文件夹路径
	if cwd, err := os.Getwd(); err == nil {
		logFilePath = cwd + logPath
	} else {
		log.Println(err.Error())
		return nil, err
	}
	// 创建log文件夹
	if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	// 查看文件是否存在，不存在就去创建
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if _, err = os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	// 写入文件
	src, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return src, nil
}
