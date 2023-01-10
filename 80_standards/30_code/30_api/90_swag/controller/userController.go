// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
	"wukong/go/50_web/20_gin/90_swag/model"
	)

// @Summary UserController 接口
// @Accept json
// @Tags Name
// @Produce  json
// @Param id path int true "ID"
// @Param name path string true "NAME"
// @Resource User
// @Router /user/{id}/{name} [get]
// @Success 200 {object} model.User
func QueryById(context *gin.Context) {
	println(">>>> get user by id and name action start <<<<")

	name := context.Param("username")

	var user model.User
	user.Username = name

	context.JSON(200, gin.H{
		"name": user.Username,
	})
}
