// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	model "go_project/scaffold/apiserver/apis/authz/question/model/v1"
	"go_project/scaffold/apiserver/apis/authz/question/repo"

	"github.com/rebirthmonkey/go/pkg/metamodel"

)

// QuestionService defines functions used to handle question request.
type QuestionService interface {
	Create(question *model.Question) error
	Delete(questionname string) error
	Update(question *model.Question) error
	Get(questionname string) (*model.Question, error)
	List() (*model.QuestionList, error)
}

// questionService is the QuestionService instance to handle question request.
type questionService struct {
	repo repo.Repo
}

var _ QuestionService = (*questionService)(nil)

// newQuestionService creates and returns the question service instance.
func newQuestionService(repo repo.Repo) QuestionService {
	return &questionService{repo}
}

// Create creates a new question account.
func (u *questionService) Create(question *model.Question) error {

	question.Status = 1

	return u.repo.QuestionRepo().Create(question)
}

// Delete deletes the question by the question identifier.
func (u *questionService) Delete(questionname string) error {
	return u.repo.QuestionRepo().Delete(questionname)
}

// Update updates a question account information.
func (u *questionService) Update(question *model.Question) error {
	/*updateQuestion, err := u.Get(question.Name)
	if err != nil {
		return err
	}

	updateQuestion.Nickname = question.Nickname
	updateQuestion.Email = question.Email
	updateQuestion.Phone = question.Phone
	updateQuestion.Extend = question.Extend*/

	return u.repo.QuestionRepo().Update(question)
}

// Get returns a question's info by the question identifier.
func (u *questionService) Get(questionname string) (*model.Question, error) {
	return u.repo.QuestionRepo().Get(questionname)
}

// List returns all the related questions.
func (u *questionService) List() (*model.QuestionList, error) {
	questions, err := u.repo.QuestionRepo().List()
	if err != nil {
		return nil, err
	}

	infos := make([]*model.Question, 0)
	for _, question := range questions.Items {
		infos = append(infos, &model.Question{
			ObjectMeta: metamodel.ObjectMeta{
				ID:        question.ID,
				Name:      question.Name,
				CreatedAt: question.CreatedAt,
				UpdatedAt: question.UpdatedAt,
			},
			Studentname: question.Studentname,
			Content:    question.Content,
		})
	}

	return &model.QuestionList{ListMeta: questions.ListMeta, Items: infos}, nil
}
