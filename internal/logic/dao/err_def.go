package dao

import "errors"

var (
	DAO_ERROR_RECORD_NOT_FOUND = errors.New("记录未找到")
	DAO_ERROR_DUPLICATE_RECORD = errors.New("不符合唯一约束")
)
