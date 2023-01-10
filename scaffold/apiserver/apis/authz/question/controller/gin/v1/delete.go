// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

// Delete deletes an question by the question identifier.
// Only administrator can call this function.
func (u *controller) Delete(c *gin.Context) {
	log.L(c).Info("[GinServer] questionController: delete")

	if err := u.srv.NewQuestionService().Delete(c.Param("name")); err != nil {
		util.WriteResponse(c, err, nil)

		return
	}

	var msg string = "deleted question " + c.Param("name")
	util.WriteResponse(c, nil, msg)
}
