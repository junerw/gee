/*
@Copyright : Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
@File : gee_test.go
@Time : 2023/6/1 16:50
@Author : junerwang@tencent.com
@Desc: this is a simple description...
*/

package geeweb

import "testing"

func TestNestedGroup(t *testing.T) {
	r := New()
	v1 := r.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")
	if v2.prefix != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
	if v3.prefix != "/v1/v2/v3" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
}
