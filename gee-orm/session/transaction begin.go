/*
@Copyright : Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
@File : transaction begin.go
@Time : 2023/6/7 19:58
@Author : junerwang@tencent.com
@Desc: this is a simple description...
*/

package session

import "geeorm/log"

func (s *Session) Begin() (err error) {
	log.Info("transaction begin")
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Info("transaction commit")
	if err = s.tx.Commit(); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) Rollback() (err error) {
	log.Info("transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		log.Error(err)
	}
	return
}
