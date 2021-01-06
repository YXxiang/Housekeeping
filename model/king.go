/**
 * @Author: LOFTER
 * @Description:
 * @File:  king
 * @Date: 2021/1/5 3:35 下午
 */
package model

//金刚位
type King struct {
	BaseInfo
	Icon   string `gorm:"size:20;not null;default:''"`
	Title  string `gorm:"size:10;not null;default:''"`
	Sort   int    `gorm:"size:10;not null;default:0"`
	ISShow int    `gorm:"size:10;not null;default:1"` //是否显示 1显示 0 不显示
}
