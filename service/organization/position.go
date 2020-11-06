// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import (
	"errors"
	"oa-auth/initializer/db"
	"oa-auth/model/authorization"
	"oa-auth/model/organization"
	serializerorganization "oa-auth/serializer/organization"

	"github.com/lauvinson/gogo/gogo"
)

// PositionCreateService 职位添加服务
type PositionCreateService struct {
	ParentID int    `form:"parentId" json:"parentId" binding:"required"`
	OfficeID int    `form:"officeId" json:"officeId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

func (receiver PositionCreateService) Execute() error {
	var count int
	db.DB.Model(&organization.Office{}).Where("id = ?", receiver.OfficeID).Count(&count)
	if count < 1 {
		return errors.New("创建失败，组织不存在")
	}
	count = 0
	db.DB.Model(&organization.Position{}).Where("id = ?", receiver.ParentID).Count(&count)
	if count < 1 {
		return errors.New("创建失败，上级不存在")
	}
	position := organization.Position{
		ParentID: receiver.ParentID,
		OfficeID: receiver.OfficeID,
		Name:     receiver.Name,
		Code:     receiver.Code,
	}
	return db.DB.Model(&organization.Position{}).Save(&position).Error
}

// PositionAddService 职位更新服务
type PositionUpdateService struct {
	ID       int    `form:"id" json:"id" binding:"required"`
	ParentID int    `form:"parentId" json:"parentId" binding:"required"`
	OfficeID int    `form:"officeId" json:"officeId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

func (receiver PositionUpdateService) Execute() error {
	position := organization.Position{
		Name: receiver.Name,
	}
	return db.DB.Where("id = ?", receiver.ID).Updates(&position).Error
}

// PositionAddService 职位删除服务
type PositionDeleteService struct {
	ID int `form:"id" json:"id" binding:"required"`
}

func (receiver PositionDeleteService) Execute() error {
	return db.DB.Where("id = ?", receiver.ID).Unscoped().Delete(&organization.Position{}).Error
}

// PositionViewService 职位查看服务
type PositionViewService struct {
	OfficeID int `form:"officeId" json:"officeId" binding:"required"`
}

func (receiver PositionViewService) Execute() (interface{}, error) {
	var result []serializerorganization.PositionSerializer
	err := db.DB.Table("positions").Find(&result, "office_id = ?", receiver.OfficeID).Error
	if nil != err {
		return result, err
	}
	var se []gogo.ForkTreeNode
	for _, v := range result {
		temp := v
		se = append(se, &temp)
	}
	a := gogo.BuildTreeByRecursive(se)
	return a, nil
}

// PositionRoleMappingAddService 职位角色添加服务
type PositionRoleMappingAddService struct {
	PositionID int `form:"positionId" json:"positionId" binding:"required"`
	RoleID     int `form:"roleId" json:"roleId" binding:"required"`
}

func (receiver PositionRoleMappingAddService) Execute() error {
	var count = 0
	db.DB.Model(&organization.Position{}).Where("id = ?", receiver.PositionID).Count(&count)
	if count < 1 {
		return errors.New("创建失败，职位不存在")
	}
	count = 0
	db.DB.Model(&authorization.Role{}).Where("id = ?", receiver.RoleID).Count(&count)
	if count < 1 {
		return errors.New("创建失败，角色不存在")
	}
	positionRole := organization.PositionRoleMapping{
		PositionID: receiver.PositionID,
		RoleID:     receiver.RoleID,
	}
	return db.DB.Model(&organization.PositionRoleMapping{}).Save(&positionRole).Error
}

// PositionRoleMappingRemoveService 职位角色删除服务
type PositionRoleMappingRemoveService struct {
	ID int `form:"id" json:"id" binding:"required"`
}

func (receiver PositionRoleMappingRemoveService) Execute() error {
	return db.DB.Where("id = ?", receiver.ID).Unscoped().Delete(&organization.PositionRoleMapping{}).Error
}
