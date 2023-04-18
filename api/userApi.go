package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// Login @Tag 用户管理
// @Summary 用户登录
// @Description 用户登录详情
// @Accept  json
// @Produce  json
// @Param   name  		body    string     true     "用户名"
// @Param   password    body    string     true     "登录密码"
// @Success 200 {string} string	"登录成功"
// @Failure 400 {string} sting "登录失败"
// @Router /api/v1/public/user/login/ [post]
func (u UserApi) Login(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "Login Success",
	})
}
