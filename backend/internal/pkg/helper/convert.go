package helper

import (
	"cozeos/internal/model"
	"cozeos/internal/types"
)

func ModelUserToTypesUser(u *model.User) *types.User {
	if u == nil {
		return nil
	}

	return &types.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Phone:       u.Phone,
		Points:      u.Points,
		WeChatID:    u.WeChatID,
		Role:        u.Role,
		VIPLevel:    u.VIPLevel,
		VIPExpireAt: u.VIPExpireAt,
		ReferrerID:  u.ReferrerID,
		Ext:         u.Ext,
	}
}

func TypesUserToModelUser(u *types.User) *model.User {
	if u == nil {
		return nil
	}

	res := &model.User{}
	res.ID = u.ID
	res.Name = u.Name
	res.Email = u.Email
	res.Phone = u.Phone
	res.Points = u.Points
	res.WeChatID = u.WeChatID
	res.Role = u.Role
	res.VIPLevel = u.VIPLevel
	res.VIPExpireAt = u.VIPExpireAt
	res.ReferrerID = u.ReferrerID
	res.Ext = u.Ext

	return res
}

func ModelUploadToTypesUpload(u *model.Upload) *types.Upload {
	if u == nil {
		return nil
	}

	res := &types.Upload{}
	res.ID = u.ID
	res.UserID = u.UserID
	res.OriginalName = u.OriginalName
	res.DownloadURL = u.DownloadURL
	res.FileFormat = u.FileFormat
	res.FileSize = u.FileSize
	res.UploadTime = u.UploadTime
	res.Remarks = u.Remarks

	return res
}

func TypesUploadToModelUpload(u *types.Upload) *model.Upload {
	if u == nil {
		return nil
	}

	res := &model.Upload{}
	res.ID = u.ID
	res.UserID = u.UserID
	res.OriginalName = u.OriginalName
	res.DownloadURL = u.DownloadURL
	res.FileFormat = u.FileFormat
	res.FileSize = u.FileSize
	res.UploadTime = u.UploadTime
	res.Remarks = u.Remarks

	return res
}
