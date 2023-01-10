// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"fmt"
	model "go_project/scaffold/apiserver/apis/authz/question/model/v1"
	questionRepoInterface "go_project/scaffold/apiserver/apis/authz/question/repo"
	"regexp"

	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/log"
	"github.com/rebirthmonkey/go/pkg/mysql"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// userRepo stores the user's info.
type questionRepo struct {
	dbEngine *gorm.DB
}

var _ questionRepoInterface.QuestionRepo = (*questionRepo)(nil)

// newQuestionRepo creates and returns a question storage.
func newQuestionRepo(cfg *mysql.CompletedConfig) questionRepoInterface.QuestionRepo {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Database,
		true,
		"Local")

	db, err := gorm.Open(mysqlDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Mysql connection fails %+v\n", err)
		return nil
	}

	return &questionRepo{dbEngine: db}
}

// close closes the repo's DB engine.
func (u *questionRepo) close() error {
	dbEngine, err := u.dbEngine.DB()
	if err != nil {
		return errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return dbEngine.Close()
}

// Create creates a new question account.
func (u *questionRepo) Create(question *model.Question) error {
	tmpQuestion := model.Question{}
	u.dbEngine.Where("name = ?", question.Name).Find(&tmpQuestion)
	if tmpQuestion.Name != "" {
		err := errors.WithCode(errcode.ErrRecordAlreadyExist, "the created question already exit")

		log.Errorf("%+v", err)
		return err
	}

	err := u.dbEngine.Create(&question).Error
	if err != nil {
		if match, _ := regexp.MatchString("Duplicate entry", err.Error()); match {
			return errors.WrapC(err, errcode.ErrRecordAlreadyExist, "duplicate entry.")
		}

		return err
	}

	return nil
}

// Delete deletes the question by the question identifier.
func (u *questionRepo) Delete(questionname string) error {
	//tmpQuestion := model.Question{}
	//u.dbEngine.Where("name = ?", questionname).Find(&tmpQuestion)
	//if tmpQuestion.Name == "" {
	//	err := errors.WithCode(errcode.ErrRecordNotFound, "the delete question not found")
	//	log.Errorf("%s\n", err)
	//	return err
	//}

	if err := u.dbEngine.Where("name = ?", questionname).Delete(&model.Question{}).Error; err != nil {
		return err
	}

	return nil
}

// Update updates a question account information.
func (u *questionRepo) Update(question *model.Question) error {
	tmpQuestion := model.Question{}
	u.dbEngine.Where("name = ?", question.Name).Find(&tmpQuestion)
	if tmpQuestion.Name == "" {
		err := errors.WithCode(errcode.ErrRecordNotFound, "the update question not found")
		log.Errorf("%s\n", err)
		return err
	}

	if err := u.dbEngine.Save(question).Error; err != nil {
		return err
	}

	return nil
}

// Get returns a question's info by the question identifier.
func (u *questionRepo) Get(questionname string) (*model.Question, error) {
	question := &model.Question{}
	err := u.dbEngine.Where("name = ?", questionname).First(&question).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(errcode.ErrRecordNotFound, "the get question not found.")
		}

		return nil, errors.WithCode(errcode.ErrDatabase, err.Error())
	}

	return question, nil
}

// List returns all the related questions.
func (u *questionRepo) List() (*model.QuestionList, error) {
	ret := &model.QuestionList{}

	d := u.dbEngine.
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
