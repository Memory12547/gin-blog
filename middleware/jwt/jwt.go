/*
 * @Author: Matt Meng
 * @Date: 1970-01-01 08:00:00
 * @LastEditors: Matt Meng
 * @LastEditTime: 2020-10-11 11:58:42
 * @Description: file content
 */
package jwt

import(
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gin-blog/pkg/util"
	"gin-blog/pkg/e"
)

func JWT() gin.HandlerFunc{
	return func(c *gin.Context){
		var code int
		var data interface{}

		code = e.SUCCESS
		//token可以放在Header、Body或URL中，这里是在URL中解析
		token:=c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		}else{
			claims,err := util.ParseToken(token)
			if err!=nil{
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}else if time.Now().Unix() > claims.ExpiresAt { //在上一步ParseToken中已经有token.Valid的检查，如果超期，会返回err不为nil，这里是二次检查是否为超期，以返回对应错误码
                code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
            }
		}

		if code!=e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":code,
				"msg":e.GetMsg(code),
				"data":data,
			})
			//当鉴权失败，调用 Abort 防止挂起的其他handler执行，因为Abort不会终止当前handler，所以还需要调用return
			c.Abort()
			return
		}
		//Next只能在中间件middleware中调用，执行handler链中的下一个函数
		c.Next()
	}
}