/*
 * @Author: Matt Meng
 * @Date: 1970-01-01 08:00:00
 * @LastEditors: Matt Meng
 * @LastEditTime: 2020-10-11 11:44:46
 * @Description: file content
 */
package models

type Auth struct{
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username,password string)bool{
	var auth Auth
	db.Select("id").Where(Auth{Username:username,Password:password}).First(&auth)
	if auth.ID>0{
		return true
	}
	return false
}