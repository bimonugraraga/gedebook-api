package controllers

import (
	"net/http"

	"gedebook.com/api/constants"
	"gedebook.com/api/dto/requests"
	"gedebook.com/api/dto/responses"
	"gedebook.com/api/errs"
	"gedebook.com/api/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController interface {
	UserRegister(ctx *gin.Context)
	UserLogin(ctx *gin.Context)
	UserProfile(ctx *gin.Context)
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
	if err == nil {
		ctx.JSON(http.StatusCreated, responses.R{
			Code:    http.StatusCreated,
			Message: "Success Register",
		})
		return
	}

}

func (ctl *userController) ParseRequestRegisterEntity(ctx *gin.Context, src *requests.UserRegisterRequest) error {
	if err := ctx.ShouldBindBodyWith(src, binding.JSON); err != nil {
		return err
	}
	return nil
}

func (ctl *userController) UserLogin(ctx *gin.Context) {
	var src requests.UserLoginRequest
	if err := ctl.ParseRequestLoginEntity(ctx, &src); err != nil {
		errs.ErrorHandler(ctx, 400, "Email and Password are Required")
		return
	}
	jwt, err := ctl.userSrv.Login(ctx, &src)

	if err == nil && jwt != nil {
		ctx.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Success Login",
			Data:    jwt,
		})
	}
}

func (ctl *userController) ParseRequestLoginEntity(ctx *gin.Context, src *requests.UserLoginRequest) error {
	if err := ctx.ShouldBindBodyWith(src, binding.JSON); err != nil {
		return err
	}
	return nil
}

func (ctl *userController) UserProfile(ctx *gin.Context) {
	value, exists := ctx.Get("user")
	if !exists {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
	}
	user, ok := value.(constants.AuthnPayload)
	if !ok {
		errs.ErrorHandler(ctx, 401, "Login Is Needed")
	}
	targetUser, err := ctl.userSrv.UserProfile(ctx, user)
	if err == nil {
		ctx.JSON(http.StatusOK, responses.R{
			Code:    http.StatusOK,
			Message: "Success Fetch Data",
			Data:    targetUser,
		})
	}
}
