package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"myself/dao/mysql"
	"myself/dao/redis"
	"myself/logger"
	"myself/pkg/snowflake"
	"myself/router"
	"myself/setting"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//1.初始化设置
	err := setting.Init()
	if err != nil {
		fmt.Println("init setting failed")
		return
	}
	//2.初始化日志
	logger.InitLogger()
	//3.雪花算法生成用户ID
	if err := snowflake.Init(viper.GetString("snowflake.starttime"), viper.GetInt64("snowflake.machineID")); err != nil {
		fmt.Printf("init snowflake failed,err:%v\n", err)
	}
	//4.mysql数据库初始化
	err = mysql.MysqlInit()
	if err != nil {
		logger.Log.Error(err)
		return
	}
	defer mysql.DBclose()
	//5.redis初始化
	err = redis.RedisInit()
	if err != nil {
		logger.Log.Error(err)
		return
	}
	//6.生成路由
	r, err := router.Setup()
	if err != nil {
		logger.Log.Error(err)
		return
	}
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}
	go func() {
		err2 := srv.ListenAndServe()
		if err2 != nil {
			logger.Log.Error(err)
			return
		}
		defer srv.Close()
	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Info("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
