// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/24.

package organization

import (
	"goa/model"

	"github.com/lauvinson/gogo/gogo"
)

type PositionSerializer struct {
	model.BaseModel
	ParentID uint                `json:"parent_id"` // 上级ID
	OfficeID uint                `json:"office_id"` // 组织ID
	Name     string              `json:"name"`      // 职位名称
	Code     string              `json:"code"`      //职位编码
	Children []gogo.ForkTreeNode `json:"children"`
}

func (receiver *PositionSerializer) GetID() uint {
	return receiver.ID
}

func (receiver *PositionSerializer) GetPID() uint {
	return receiver.ParentID
}

func (receiver *PositionSerializer) Following(v []gogo.ForkTreeNode) {
	receiver.Children = append(receiver.Children, v...)
}
