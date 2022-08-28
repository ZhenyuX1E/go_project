// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"

	"github.com/rebirthmonkey/go/80_standards/30_code/80_server/apiserver/user/repo"
	srv "github.com/rebirthmonkey/go/80_standards/30_code/80_server/apiserver/user/service/v1"
)

type Controller interface {
	ListUsers(ctx context.Context, r *ListUsersRequest) (*ListUsersResponse, error)
}

type controller struct {
	srv srv.Service
	UnimplementedUserServer
}

func NewController(repo repo.Repo) *controller {
	return &controller{
		srv: srv.NewService(repo),
	}
}
