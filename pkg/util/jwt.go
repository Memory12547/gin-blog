/*
 * @Author: Matt Meng
 * @Date: 1970-01-01 08:00:00
 * @LastEditors: Matt Meng
 * @LastEditTime: 2020-10-09 23:20:58
 * @Description: file content
 */
package util

import(
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"gin-blog/pkg/setting"
)

//自定义密码
var jwtSecret = []byte(setting.JwtSecret)

//自定义Claims，除jwt.StandardClaims官方字段，补充用户即密码字段
type Claims struct{
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username,password string)(string,error){
	nowTime:=time.Now()
	expireTime:=nowTime.Add(3*time.Hour)

	claims:=Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt : expireTime.Unix(),  //过期时间
			Issuer : "gin-blog",            //签发人
		},
	}

	//指定使用的签名方法，生成token对象
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err:=tokenClaims.SignedString(jwtSecret)

	return token,err
}

func ParseToken(token string)(*Claims,error){
	tokenClaims,err:=jwt.ParseWithClaims(token,&Claims{},func(token *jwt.Token)(interface{},error){
		return jwtSecret,nil
	})
	if tokenClaims!=nil{
		if claims,ok:=tokenClaims.Claims.(*Claims);ok&&tokenClaims.Valid{
			return claims,nil
		}
	}

	return nil,err
}