/*
@Copyright : Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
@File : sqlite3_test.go
@Time : 2023/6/7 11:14
@Author : junerwang@tencent.com
@Desc: this is a simple description...
*/

package dialect

import (
	"reflect"
	"testing"
)

func TestDataTypeOf(t *testing.T) {
	dial := &sqlite3{}
	cases := []struct {
		Value interface{}
		Type  string
	}{
		{"Tom", "text"},
		{123, "integer"},
		{1.2, "real"},
		{[]int{1, 2, 3}, "blob"},
	}

	for _, c := range cases {
		if typ := dial.DataTypeOf(reflect.ValueOf(c.Value)); typ != c.Type {
			t.Fatalf("expect %s, but got %s", c.Type, typ)
		}
	}
}
