/*
@Copyright : Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
@File : dialect.go
@Time : 2023/6/6 19:27
@Author : junerwang@tencent.com
@Desc: this is a simple description...
*/

package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
