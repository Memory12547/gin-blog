/*
 * @Author: Matt Meng
 * @Date: 1970-01-01 08:00:00
 * @LastEditors: Matt Meng
 * @LastEditTime: 2021-08-28 15:25:46
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

	"gin-blog/pkg/setting"
	"gin-blog/routers"
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

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
