/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-09-10 16:20
* Description:
*****************************************************************/

package xentity

import (
	"github.com/go-xe2/x/xf/ef/xq/xdatabase"
)

func (ent *TEntity) BeginTrans() error {
	db := xdatabase.DB(ent.dbName...)
	return db.BeginTran()
}

func (ent *TEntity) CommitTrans() error {
	db := xdatabase.DB(ent.dbName...)
	return db.Commit()
}

func (ent *TEntity) RollbackTrans() error {
	db := xdatabase.DB(ent.dbName...)
	return db.Rollback()
}
