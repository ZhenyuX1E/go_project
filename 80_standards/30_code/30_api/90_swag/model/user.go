// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

type User struct {
	Username 	string	`json:"name" form:"name"`
	Age  		uint8	`json:"age" form:"age"`
	Mobile 		string	`json:"mobile" form:"mobile"`
	Sex			string	`json:"sex" form:"sex"`
	Address 	string	`json:"address" form:"address"`
	Id          uint16  `json:"id" form:"id"`
}


