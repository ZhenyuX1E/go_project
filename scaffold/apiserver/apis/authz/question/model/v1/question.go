// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (

	"github.com/rebirthmonkey/go/pkg/metamodel"
	"github.com/rebirthmonkey/go/pkg/util"
	"gorm.io/gorm"
)

// Question represents a Question restful resource. It is also used as data model.
type Question struct {
	metamodel.ObjectMeta `json:"metadata,omitempty"`

	Status      int64     `json:"status"              gorm:"column:status"    validate:"omitempty"`
	Studentname string    `json:"studentname"         gorm:"column:studentname"  validate:"required"`
	Content     string    `json:"content"             gorm:"column:content"   validate:"required"`

}

type QuestionRequest struct {
	Questioninst   Question   `json:"question"`

}

// QuestionList is the whole list of all Questions which have been stored in the storage.
type QuestionList struct {
	// +optional
	metamodel.ListMeta `json:",inline"`

	Items []*Question `json:"items"`
}

// TableName maps to mysql table name.
func (u *Question) TableName() string {
	return "question"
}


// AfterCreate run after create database record.
func (u *Question) AfterCreate(tx *gorm.DB) error {
	u.InstanceID = util.GetInstanceID(u.ID, "question-")

	return tx.Save(u).Error
}
