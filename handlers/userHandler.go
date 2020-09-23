package handlers

import (
	"GoProj1.0/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func UserSaveService(ctx *gin.Context) {
	var user models.User
	if err1 := ctx.ShouldBind(&user); err1 != nil {
		log.Panicln("模型绑定错误", err1.Error())
		return
	}
	log.Println("user name:", user.Username)
	if err2 := user.UserSave(); err2 != nil {
		log.Panicln("err->", err2.Error())
		return
	}
}

func UserDeleteService(ctx *gin.Context) {
	fmt.Println("???")
	var user models.User
	if err1 := ctx.ShouldBind(&user); err1 != nil {
		log.Panicln("模型绑定错误", err1.Error())
		return
	}
	log.Println("user id:", user.Id)
	log.Println("user name:", user.Username)
	if err2 := user.UserDelete(); err2 != nil {
		log.Panicln("err->", err2.Error())
		return
	}
}

//func UserSave(ctx *gin.Context){
//	//uname,upwd:=ctx.PostForm("username"),ctx.PostForm("password")
//	//log.Println("username:",uname)
//	//log.Println("password:",upwd)
//
//
//	var user models.User
//	fmt.Println("测试开始")
//	ctx.ShouldBind(&user)
//	fmt.Println("测试进行中")
//	fmt.Println(user.Username)
//	fmt.Println(user.Password)
//	err:=(&user).UserSave()
//	if err!=nil{
//		log.Println(err)
//	}
//
//	//if err1 := ctx.ShouldBind(&user);err1 != nil{
//	//	log.Println(user.UserName)
//	//
//	//
//	//	ctx.String(http.StatusBadRequest,"发生错误")
//	//	log.Println("???????",err1.Error())
//	//}else{
//	//	err2:=user.UserSave()
//	//	if err2!=nil{
//	//		log.Println("数据存储失败：",err2.Error())
//	//	}else{
//	//		log.Println("数据存储成功：")
//	//	}
//	//}
//}
