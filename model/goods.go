/**
 * @Author: LOFTER
 * @Description:
 * @File:  good
 * @Date: 2021/1/5 3:48 下午
 */
package model

type Goods struct {
	BaseInfo
	Title    string  `gorm:"size:20;not null;default:''"`
	Price    float64 `gorm:"size:10;not null;default:0.0"`
	Label    string  `gorm:"size:10;not null;default:''"`  //标签
	Sold     int     `gorm:"size:100;not null;default:0"`  //已售
	Stock    int     `gorm:"size:100;not null;default:0"`  //库存
	Praise   float64 `gorm:"size:10;not null;default:0.0"` //好评
	Describe string  `gorm:"size:10;not null;default:''"`  //服务描述
	KingID   int     `gorm:"size:10;not null;"`            //属于那个king
}
