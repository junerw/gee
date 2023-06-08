/*
@Copyright : Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
@File : logger.go
@Time : 2023/6/1 17:38
@Author : junerwang@tencent.com
@Desc: this is a simple description...
*/

package geeweb

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
