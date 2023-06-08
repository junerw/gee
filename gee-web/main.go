/*
@Copyright : Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
@File : main.go
@Time : 2023/5/31 16:43
@Author : junerwang@tencent.com
@Desc: this is a simple description...
*/

package main

import (
	"geeweb/geeweb"
	"net/http"
)

func main() {
	r := geeweb.Default()
	r.GET("/", func(c *geeweb.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *geeweb.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
