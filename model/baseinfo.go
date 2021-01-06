/**
 * @Author: LOFTER
 * @Description:
 * @File:  baseinfo
 * @Date: 2021/1/5 3:26 下午
 */
package model

import (
	"gorm.io/gorm"
	"time"
)

//基础字段common
type BaseInfo struct {
	ID        int            `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
