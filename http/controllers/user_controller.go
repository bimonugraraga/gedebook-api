package controllers

import (
	"net/http"

	"gedebook.com/api/dto/requests"
	"gedebook.com/api/dto/responses"
	"gedebook.com/api/errs"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController interface {
	UserRegister(ctx *gin.Context)
}

type userController struct {
	userSrv services.UserService
}

func NewUserController(userSrv services.UserService) UserController {
	return &userController{
		userSrv: userSrv,
	}
}

func (ctl *userController) UserRegister(ctx *gin.Context) {
	var src requests.UserRegisterRequest
	if err := ctl.ParseRequestRegisterEntity(ctx, &src); err != nil {
		errs.ErrorHandler(ctx, 400, "Email, Password, and Name are Required")
		return
	}
	newUser, err := src.AssignedUserRegister()
	err = ctl.userSrv.Register(ctx, &newUser)
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To Register")
		return
	}

	ctx.JSON(http.StatusCreated, responses.R{
		Code:    http.StatusCreated,
		Message: "Success Register",
	})
}

func (ctl *userController) ParseRequestRegisterEntity(ctx *gin.Context, src *requests.UserRegisterRequest) error {
	if err := ctx.ShouldBindBodyWith(src, binding.JSON); err != nil {
		return err
	}
	return nil
}