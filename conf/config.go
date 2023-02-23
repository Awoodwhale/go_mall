package conf

import (
	"go_mall/cache"
	"go_mall/dao"
	"go_mall/pkg/utils"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	// service
	AppMode  string
	HttpPort string
	// mysql
	DBHost string
	DBPort string
	DBUser string
	DBPwd  string
	DBName string
	// redis
	RedisHost string
	RedisPort string
	RedisPwd  string
	RedisName string
	// email
	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpToken  string
	// img
	ImgHost     string
	ImgPort     string
	ProductPath string
	AvatarPath  string
	// log
	LogPath string
)

// Init
// @Description: 从config.ini读取配置
func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	// load config file
	loadServer(file)
	loadMySQL(file)
	loadRedis(file)
	loadEmail(file)
	loadImage(file)
	loadLog(file)
	// logger init
	utils.InitLog(LogPath)
	// mysql read（主）
	pathRead := strings.Join([]string{DBUser, ":", DBPwd, "@tcp(", DBHost, ":", DBPort, ")/", DBName, "?charset=utf8mb4&parseTime=true"}, "")
	// mysql write（从）
	pathWrite := strings.Join([]string{DBUser, ":", DBPwd, "@tcp(", DBHost, ":", DBPort, ")/", DBName, "?charset=utf8mb4&parseTime=true"}, "")
	// mysql init
	dao.InitDatabase(pathRead, pathWrite)
	// redis init
	cache.InitDatabase(RedisHost+":"+RedisPort, RedisName, RedisPwd)
}

// loadServer
// @Description: 获取server的config
// @param file *ini.File
func loadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = ":" + file.Section("service").Key("HttpPort").String() // 加上:前缀
}

// loadMySQL
// @Description: 获取MySQL的config
// @param file *ini.File
func loadMySQL(file *ini.File) {
	DBHost = file.Section("mysql").Key("DBHost").String()
	DBPort = file.Section("mysql").Key("DBPort").String()
	DBUser = file.Section("mysql").Key("DBUser").String()
	DBPwd = file.Section("mysql").Key("DBPwd").String()
	DBName = file.Section("mysql").Key("DBName").String()
}

// loadRedis
// @Description: 获取redis的config
// @param file *ini.File
func loadRedis(file *ini.File) {
	RedisHost = file.Section("redis").Key("RedisHost").String()
	RedisPort = file.Section("redis").Key("RedisPort").String()
	RedisPwd = file.Section("redis").Key("RedisPwd").String()
	RedisName = file.Section("redis").Key("RedisName").String()
}

// loadEmail
// @Description: 获取email的config
// @param file *ini.File
func loadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpToken = file.Section("email").Key("SmtpToken").String()
}

// loadImage
// @Description: 获取image的config
// @param file *ini.File
func loadImage(file *ini.File) {
	ImgHost = file.Section("image").Key("ImgHost").String()
	ImgPort = file.Section("image").Key("ImgPort").String()
	ProductPath = file.Section("image").Key("ProductPath").String()
	AvatarPath = file.Section("image").Key("AvatarPath").String()
}

// loadLog
// @Description: 获取log配置
// @param file *ini.File
func loadLog(file *ini.File) {
	LogPath = file.Section("log").Key("LogPath").String()
}
