/**
 * @Author: LOFTER
 * @Description:
 * @File:  banner
 * @Date: 2021/1/5 3:43 下午
 */
package model

type Banner struct {
	BaseInfo
	Title   string `json:"title" gorm:"size:10;not null;default:''"`
	Type    string `json:"type" gorm:"size:10;not null;default:'H5'"` //banner 类型 是商品good  还是h5
	Icon    string `json:"icon" gorm:"size:20;not null;default:''"`
	GoodsID int    `json:"goods_id" gorm:"size:20;not null;"`
}
