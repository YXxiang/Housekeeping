/**
 * @Author: LOFTER
 * @Description:
 * @File:  review
 * @Date: 2021/1/5 4:01 下午
 */
package model

//商品评论
type Review struct {
	BaseInfo
	Content  string `gorm:"size:10;not null;default:''"`
	Speed    int    `gorm:"size:10;not null;default:0"`
	Attitude int    `gorm:"size:10;not null;default:0"`
	Quality  int    `gorm:"size:10;not null;default:0"`
	GoodsID  int    `gorm:"size:10;not null"` //那个商品的评论
	UserID   uint   `gorm:"size:10;not null"`
}
