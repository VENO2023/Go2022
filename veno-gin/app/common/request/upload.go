package request

import "mime/multipart"

/*
 * @Author: ych
 * @Description: ...
 * @File: upload
 * @Version: ...
 * @Date: 2022-11-01 16:43:24
 */

type ImageUpload struct {
	Business string                `form:"business" json:"business" binding:"required"`
	Image    *multipart.FileHeader `form:"image" json:"image" binding:"required"`
}

func (imageUpload ImageUpload) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"business.required": "业务类型不能为空",
		"image.required":    "请选择图片",
	}
}
