/**
 * @Author: LOFTER
 * @Description:
 * @File:  user
 * @Date: 2021/1/5 4:22 下午
 */
package model

type User struct {
	BaseInfo
	Name string `gorm:"size:10;not null;default:''"`
	Icon string `gorm:"size:10;not null;default:''"`
}
