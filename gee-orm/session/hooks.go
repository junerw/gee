/*
@Copyright : Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
@File : hooks.go
@Time : 2023/6/7 19:35
@Author : junerwang@tencent.com
@Desc: this is a simple description...
*/

package session

import (
	"geeorm/log"
	"reflect"
)

const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

// CallMethod calls the registered hooks
func (s *Session) CallMethod(method string, value interface{}) {
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
	return
}
