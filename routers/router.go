/*
 * @Author: Matt Meng
 * @Date: 2020-10-07 12:11:22
 * @LastEditTime: 2020-10-07 12:22:53
 * @LastEditors: Please set LastEditors
 * @Description: router configuration
 * @FilePath: /go/src/gin-blog/routers/router.go
 */
package routers

import (
	"github.com/gin-gonic/gin"
	"gin-blog/pkg/setting"
)



func InitRouter()*gin.Engine{
	//创建一个没有中间件的Engine实例
	r := gin.New()
	//添加Logger中间件
	r.Use(gin.Logger())
	//添加Recovery中间件
	r.Use(gin.Recovery())
	//根据配置设置运行模式
	gin.SetMode(setting.RunMode)
	//路由会被注册进radix tree(简化的trie树)，每种方法(GET、POST)有一颗单独的树，提高查询路由效率
	r.GET("/test",func(c *gin.Context){
		c.JSON(200,gin.H{             //gin.H{...}：就是一个map[string]interface{}
			"messgae":"test",
		})
	})
	return r
}
