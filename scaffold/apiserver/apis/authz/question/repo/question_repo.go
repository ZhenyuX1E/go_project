// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package repo

import (
	model "go_project/scaffold/apiserver/apis/authz/question/model/v1"
)

// QuestionRepo defines the Question resources.
type QuestionRepo interface {
	Create(Question *model.Question) error
	Delete(Questionname string) error
	Update(Question *model.Question) error
	Get(Questionname string) (*model.Question, error)
	List() (*model.QuestionList, error)
}
