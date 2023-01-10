// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	repo3 "go_project/scaffold/apiserver/apis/authz/question/repo"
	"sync"

	"github.com/rebirthmonkey/go/pkg/mysql"
)

// repo defines the APIServer storage.
type repo struct {
	questionRepo repo3.QuestionRepo
}

var (
	r    repo
	once sync.Once
)

var _ repo3.Repo = (*repo)(nil)

// Repo creates and returns the store client instance.
func Repo(cfg *mysql.CompletedConfig) (repo3.Repo, error) {
	once.Do(func() {
		r = repo{
			questionRepo: newQuestionRepo(cfg),
		}
	})

	return r, nil
}

// QuestionRepo returns the question store client instance.
func (r repo) QuestionRepo() repo3.QuestionRepo {
	return r.questionRepo
}

// Close closes the repo.
func (r repo) Close() error {
	return r.Close()
}
