package cmd

import "C"
import (
	"EIM/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-micro.dev/v4"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConfigFileSwitch 配置文件选择本地config.yaml OR 项目缓存中的config.yaml
// true 则是选择项目缓存编译进去的，false 则是选择本地config.yaml
const ConfigFileSwitch = false

type Server struct {
	GinPort string
	Config  *config.Config // 配置参数
}

// RunConfig  启动获取配置程序
func (server *Server) RunConfig() {
	server.Config = new(config.Config)
	logrus.Info("check config type local OR cache file [config code :··cmd/RunServer.go·· ConfigFileSwitch]")
	if ConfigFileSwitch {
		server.Config.FileRead()
	} else {
		server.Config.FileCache = config.CacheFile
	}
	server.Config.RunFileConfig()
}

// RunDataBase 启动数据库
func (server *Server) RunDataBase() {
	sqlInit := fmt.Sprintf("%s:%s@tcp(%s)/%s?utf8mb4&parseTime=True&loc=Local",
		server.Config.DataBase.User,
		server.Config.DataBase.Password,
		server.Config.DataBase.Addr,
		server.Config.DataBase.Schema)
	SQL, err := gorm.Open(mysql.Open(sqlInit), &gorm.Config{
		// 设置打印gorm查询时的打印语句 https://gorm.io/zh_CN/docs/logger.html,需要指定打印级别
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Error("mysql socks error：", err)
		panic(err)
	}
	// 利用database/sql设置数据库连接池
	sql, err := SQL.DB()
	// 设置最大连接数 默认为0 也就是没有限制
	sql.SetMaxOpenConns(0)
	// 设置最大空闲连接 每次执行完语句都将连接放入连接池，默认为2
	sql.SetConnMaxIdleTime(100000)
	logrus.Info("gorm socks mysql success")
}

// RunGinEngine 启动gin接口服务
// 提供部分http接口
func (server *Server) RunGinEngine() {
	engine := gin.Default()
	engine.Run(":8888")
}

// RunMicro 配置微服务接口
// 学习文档 http://m1.topgoer.com/Guides/WritingaGoService.html
func (server *Server) RunMicro() {
	service := micro.NewService(
		micro.Name("HelloWorld"),
	)

	// initialise flags
	service.Init()

	// start the service
	service.Run()
}
