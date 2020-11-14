/*
 * @Author: Matt Meng
 * @Date: 1970-01-01 08:00:00
 * @LastEditors: Matt Meng
 * @LastEditTime: 2020-10-15 23:20:21
 * @Description: file content
 */
package main

import (
    "fmt"
    "net/http"
    "context"
    "log"
    "os"
    "os/signal"
    "time"


    "gin-blog/routers"
    "gin-blog/pkg/setting"
)

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
    <- quit

    log.Println("Shutdown Server ...")

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    if err := s.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }

    log.Println("Server exiting")
}