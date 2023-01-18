package requests

import (
	"gedebook.com/api/domain"
	"github.com/jinzhu/copier"
)

type UserRegisterRequest struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
}

func (src UserRegisterRequest) AssignedUserRegister() (res domain.User, err error) {
	if err := copier.Copy(&res, &src); err != nil {
		return domain.User{}, err
	}
	return
}
