/*
 * @Author: Matt Meng
 * @Date: 1970-01-01 08:00:00
 * @LastEditTime: 2020-10-07 14:26:12
 * @LastEditors: Matt Meng
 * @Description: In User Settings Edit
 * @FilePath: /go/src/gin-blog/pkg/util/pagination.go
 */
package util

import (
    "github.com/gin-gonic/gin"
    "github.com/unknwon/com"

    "gin-blog/pkg/setting"
)

func GetPage(c *gin.Context) int {
    result := 0
    page, _ := com.StrTo(c.Query("page")).Int()
    if page > 0 {
        result = (page - 1) * setting.PageSize
    }

    return result
}
