/*
 * @Author: your name
 * @Date: 1970-01-01 08:00:00
 * @LastEditTime: 2020-10-07 11:28:03
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go/src/gin-blog/main.go
 */
package main

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-blog/pkg/setting"
)

func main(){
	//
	router:=gin.Default()
	//路由会被注册进radix tree(简化的trie树)，每种方法(GET、POST)有一颗单独的树，提高查询路由效率
	router.GET("/test",func(c *gin.Context){
		c.JSON(200,gin.H{             //gin.H{...}：就是一个map[string]interface{}
			"messgae":"test",
		})
	})

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