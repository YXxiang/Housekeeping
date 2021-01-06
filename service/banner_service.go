/**
 * @Author: LOFTER
 * @Description:
 * @File:  banner_service
 * @Date: 2021/1/6 7:08 下午
 */
package service

import (
	"Housekeeping/model"
	"Housekeeping/serializer"
)

//定义前段传递的数据字段
type BannerParams struct {
}

//增加BannerParams
func (item *BannerParams) CreateBannerParams() serializer.Response {

	var info []*model.Banner

	if err := model.DB.Find(&info).Error; err != nil {
		return serializer.Response{
			Code:  0,
			Data:  nil,
			Msg:   err.Error(),
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data:  nil,
		Msg:   "添加成功",
		Error: "",
	}
}

//删 BannerParams
func (item *BannerParams) DelBannerParams() serializer.Response {

	return serializer.Response{
		Data: nil,
		Msg:  "删除成功",
	}
}

//查ALL BannerParams
func (item *BannerParams) GetALLBannerParams() serializer.Response {

	return serializer.Response{
		Data: nil,
		Msg:  "查询ALL成功",
	}
}

//查ONE BannerParams
func (item *BannerParams) GetONEBannerParams() serializer.Response {

	return serializer.Response{
		Data: nil,
		Msg:  "查询ONE成功",
	}
}

//更新 BannerParams
func (item *BannerParams) UpdateBannerParams() serializer.Response {

	return serializer.Response{
		Data: nil,
		Msg:  "更新成功",
	}
}
