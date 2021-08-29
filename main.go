/*
 * @Author: Matt Meng
 * @Date: 1970-01-01 08:00:00
 * @LastEditors: Matt Meng
 * @LastEditTime: 2021-08-29 20:57:56
 * @Description: file content
 */
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/robfig/cron"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"gin-blog/models"
)

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅： 一起用 Go 做项目
// @termsofService https://github.com/go-programming-tour-book
func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	//增加定时删除无效tag和article
	c := cron.New()
	c.AddFunc("*/4 * * * * *", func() {
		// TODO:打印级别调整
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("*/4 * * * * *", func() {
		// TODO:打印级别调整
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
 
	c.Start()

	//转发os.Interrupt信号到quit
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	//通过context实现 goroutine 同步关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
