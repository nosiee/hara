package controllers

import "github.com/gin-gonic/gin"

func SignIn(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")

	println(username, password)
}
