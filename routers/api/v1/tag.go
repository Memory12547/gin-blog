/*
 * @Author: Matt Meng
 * @Date: 1970-01-01 08:00:00
 * @LastEditors: Matt Meng
 * @LastEditTime: 2020-10-08 13:45:04
 * @Description: tap api
 */

package v1

import(
	"log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"
    "github.com/Unknwon/com"

    "gin-blog/pkg/e"
    "gin-blog/models"
    "gin-blog/pkg/util"
    "gin-blog/pkg/setting"
)

//获取多个文章标签
func GetTags(c *gin.Context){
	name:=c.Query("name")
	maps:=make(map[string]interface{})
	data:=make(map[string]interface{})

	if name != ""{
		maps["name"]=name
	}

	var state int =-1
	if arg:=c.Query("state");arg!=""{
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
    //调用models下的tag接口
	data["lists"]= models.GetTags(util.GetPage(c),setting.PageSize,maps)
	data["total"]=models.GetTagTotal(maps)

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}

//增加文章标签
func AddTag(c *gin.Context){
	name:=c.Query("name")
	state:=com.StrTo(c.DefaultQuery("state","0")).MustInt()
	createdBy:=c.Query("created_by")

	valid:=validation.Validation{}
	valid.Required(name,"name").Message("名称不能为空")
	valid.MaxSize(name,100,"name").Message("名称最长为100字符")
	valid.Required(createdBy,"created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy,100,"created_by").Message("创建人最长为100字符")
	valid.Range(state,0,1,"state").Message("状态只允许0或1")

	code:=e.INVALID_PARAMS
	if !valid.HasErrors(){
		if !models.ExistTagByName(name){
			code=e.SUCCESS
			models.AddTag(name,state,createdBy)
		}else{
			code=e.ERROR_EXIST_TAG
		}
	}else{
		// 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]string),
	})
}
//修改文章标签
func EditTag(c *gin.Context){
	id:=com.StrTo(c.Param("id")).MustInt()
	name:=c.Query("name")
	modifiedBy:=c.Query("modified_by")

	//表单数据验证
	valid:=validation.Validation{}

	var state int=-1
	if arg:=c.Query("state");arg!=""{
		state=com.StrTo(arg).MustInt()
		valid.Range(state,0,1,"state").Message("状态只允许为0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
    valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
    valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
    valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	
	code:=e.INVALID_PARAMS
	if ! valid.HasErrors(){
		//表单数据验证通过，调用models下封装的数据库操作函数
		code=e.SUCCESS
		if models.ExistTagByID(id){
			data:=make(map[string]interface{})
			data["modified_by"]=modifiedBy
			if name!=""{
				data["name"]=name
			}
			if state!=-1{
				data["state"]=state
			}
			models.EditTag(id,data)
		}else{
			code=e.ERROR_NOT_EXIST_TAG
		}
	}else{
		// 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
	}
	//构建返回的json数据
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]string),
	})
}
//删除文章标签
func DeleteTag(c *gin.Context){
	id:=com.StrTo(c.Param("id")).MustInt()

	//表单数据验证
	valid:=validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

    code := e.INVALID_PARAMS
    if ! valid.HasErrors() {
        code = e.SUCCESS
        if models.ExistTagByID(id) {
            models.DeleteTag(id)
        } else {
            code = e.ERROR_NOT_EXIST_TAG
        }
	}else{
		// 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
	}
	
	c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}