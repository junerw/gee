/*
@Copyright : Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
@File : log_test.go
@Time : 2023/6/6 19:57
@Author : junerwang@tencent.com
@Desc: this is a simple description...
*/

package log

import (
	"os"
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel(ErrorLevel)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() != os.Stdout {
		t.Fatal("failed to set log level")
	}
	SetLevel(Disabled)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() == os.Stdout {
		t.Fatal("failed to set log level")
	}
}
