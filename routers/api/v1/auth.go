/*
 * @Author: Matt Meng
 * @Date: 1970-01-01 08:00:00
 * @LastEditors: Matt Meng
 * @LastEditTime: 2020-10-15 22:19:24
 * @Description: file content
 */
package v1

import(
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"

	"gin-blog/pkg/util"
	"gin-blog/pkg/e"
	"gin-blog/models"
    "gin-blog/pkg/logging"
)

type auth struct{
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context){
	username:=c.Query("username")
	password:=c.Query("password")

	valid:=validation.Validation{}
	a:=auth{Username:username,Password:password}
	//使用auth结构体的定义中的tag进行Valid验证
	ok,_:=valid.Valid(&a)

	data:=make(map[string]interface{})
	code:=e.INVALID_PARAMS
	if ok{
		isExist:=models.CheckAuth(username,password)
		if isExist {
			token,err:=util.GenerateToken(username,password)
			if err!=nil{
				code=e.ERROR_AUTH_TOKEN
			}else{
				data["token"]=token
				code = e.SUCCESS
			}
		}else{
			code=e.ERROR_AUTH
		}
	}else{
		for _,err:=range valid.Errors{
			logging.Info(err.Key,err.Message)
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})

}