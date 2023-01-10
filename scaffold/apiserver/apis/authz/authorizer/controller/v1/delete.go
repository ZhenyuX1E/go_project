package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/pkg/errors"
	"github.com/rebirthmonkey/pkg/gin/util"
	"github.com/gin-gonic/gin/binding"
	"github.com/rebirthmonkey/go/pkg/log"
)

// Delete deletes an question by the question identifier.
// Only administrator can call this function.
func (u *controller) Delete(c *gin.Context) {
	var request ladon.Request
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		util.WriteResponse(c, errors.WithCode(errcode.ErrBind, err.Error()), nil)
		return
	}

	if request.Context == nil {
		request.Context = ladon.Context{}
	}
	request.Context["username"] = c.GetString("username")

	res := u.srv.NewAuthorizerService().Authorize(&request)
	if res.Allowed == true{
		log.L(c).Info("[GinServer] questionController: delete")

		if err := u.questionsrv.NewQuestionService().Delete(c.Param("name")); err != nil {
			util.WriteResponse(c, err, nil)

			return
		}

		var msg string = "deleted question " + c.Param("name")
		util.WriteResponse(c, nil, msg)
	}

	
}
