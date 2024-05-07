package middleware

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type PermissionRoleResponse struct {
	ID           int    `json:"id"`
	PermissionID string `json:"permissionId"`
	Name         string `json:"name"`
	PermissionAccess
}

type PermissionAccess struct {
	View                bool `pg:"view,type:boolean, use_zero" json:"view" example:"true"`
	Edit                bool `pg:"edit,type:boolean, use_zero" json:"edit" example:"true"`
	Create              bool `pg:"create,type:boolean, use_zero" json:"create" example:"true"`
	Delete              bool `pg:"delete,type:boolean, use_zero" json:"delete" example:"true"`
	EditOnlyYouCreate   bool `pg:"edit_onlyyou,type:boolean, use_zero" json:"editOnlyYou" example:"true"`
	DeleteOnlyYouCreate bool `pg:"delete_onlyyou,type:boolean, use_zero" json:"deleteOnlyYou" example:"true" `
}

type PermissionRoleResponses []PermissionRoleResponse

type adminClaim struct {
	UserID       string                  `json:"userId"`
	Username     string                  `json:"username"`
	Email        string                  `json:"email"`
	DepartmentID int                     `json:"majorId"`
	RoleID       int                     `json:"roleId"`
	Role         string                  `json:"role"`
	RoleName     string                  `json:"roleName"`
	Permissions  PermissionRoleResponses `json:"permissions"`
	jwt.StandardClaims
}

func GetAdminClaims(tokenString string) (*adminClaim, error) {
	var parser jwt.Parser
	var c adminClaim
	token, _, err := parser.ParseUnverified(
		tokenString,
		&c,
	)

	if err != nil {
		return nil, err
	}

	_, ok := token.Claims.(*adminClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}

	return &c, nil
}
