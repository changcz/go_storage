package models

import "goproject/short_chain/global"

type Crud struct {
}

// Add 添加
func (c *Crud) Add(data interface{}) error {
	db := global.Db
	return db.Create(data).Error
}

// QueryOne 查询
func (c *Crud) QueryOne(query interface{}) error {
	db := global.Db
	return db.Where(*&query).First(query).Error
}
