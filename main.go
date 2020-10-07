/*
 * @Author: your name
 * @Date: 1970-01-01 08:00:00
 * @LastEditTime: 2020-10-07 12:23:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go/src/gin-blog/main.go
 */
package main

import(
	"fmt"
	"net/http"

	"gin-blog/pkg/setting"
	"gin-blog/routers"
)

func main(){
	router:=routers.InitRouter()
	//直接使用http.Server配置参数，而不是用router.Run()
	s:=&http.Server{
		Addr: fmt.Sprintf(":%d",setting.HTTPPort),
		Handler: router,
		ReadTimeout:setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}