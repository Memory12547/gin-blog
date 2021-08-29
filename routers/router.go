/*
 * @Author: Matt Meng
 * @Date: 2020-10-07 12:11:22
 * @LastEditTime: 2021-08-28 15:18:58
 * @LastEditors: Matt Meng
 * @Description: router configuration
 * @FilePath: /go/src/gin-blog/routers/router.go
 */
package routers

import (
	_ "gin-blog/docs"
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	//创建一个没有中间件的Engine实例
	r := gin.New()
	//添加Logger中间件
	r.Use(gin.Logger())
	//添加Recovery中间件
	r.Use(gin.Recovery())
	//根据配置设置运行模式
	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//获取token路由绑定
	r.GET("/auth", v1.GetAuth)
	//
	apiv1 := r.Group("api/v1")
	//指定Group apiv1使用自定义的JWT中间件
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新增标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定id的标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定id的标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定id的文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新增文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定id的文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定id的文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}

	return r
}
