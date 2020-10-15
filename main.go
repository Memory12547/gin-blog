/*
 * @Author: your name
 * @Date: 1970-01-01 08:00:00
 * @LastEditTime: 2020-10-15 23:05:55
 * @LastEditors: Matt Meng
 * @Description: In User Settings Edit
 * @FilePath: /go/src/gin-blog/main.go
 */
package main

import(
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"

	"gin-blog/pkg/setting"
	"gin-blog/routers"
)

func main(){
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1<<20
	endPoint := fmt.Sprintf(":%d",setting.HTTPPort)
	
	server := endless.NewServer(endPoint,routers.InitRouter())
	server.BeforeBegin = func(add string){
		log.Printf("Article pis is %d",syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err!=nil {
		log.Printf("Server err: %v",err)
	}
}