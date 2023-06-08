/*
@Copyright : Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
@File : table_test.go
@Time : 2023/6/7 11:20
@Author : junerwang@tencent.com
@Desc: this is a simple description...
*/

package session

import "testing"

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	s := NewSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}

func TestSession_Model(t *testing.T) {
	s := NewSession().Model(&User{})
	table := s.RefTable()
	s.Model(&Session{})
	if table.Name != "User" || s.RefTable().Name != "Session" {
		t.Fatal("Failed to change model")
	}
}
