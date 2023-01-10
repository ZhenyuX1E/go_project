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

func (u *controller) List(c *gin.Context) {
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
		log.L(c).Info("[GinServer] questionController: list")
		questions, err := u.questionsrv.NewQuestionService().List()
		if err != nil {
			util.WriteResponse(c, err, nil)

			return
		}

		util.WriteResponse(c, nil, questions)
	}

	
}