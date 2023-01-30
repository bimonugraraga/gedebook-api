package services

import (
	"context"
	"database/sql"

	"gedebook.com/api/db"
	"gedebook.com/api/domain"
	"gedebook.com/api/domain/repository"
	"gedebook.com/api/dto/requests"
	"gedebook.com/api/dto/responses"
	"gedebook.com/api/errs"
	"gedebook.com/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type UserService interface {
	Register(ctx *gin.Context, src *domain.User) error
	Login(ctx *gin.Context, src *requests.UserLoginRequest) (*responses.UserLoginResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (srv *userService) Register(ctx *gin.Context, src *domain.User) error {
	c, cancel := repository.NewContext(ctx)
	defer cancel()
	if err := db.GetConn().RunInTx(c, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) (err error) {
		hashPassword, err := utils.HashPassword(src.Password)
		if err != nil {
			return err
		}
		src.Password = hashPassword
		if err = srv.userRepo.Register(c, src); err != nil {
			return err
		}
		return

	}); err != nil {
		errNum := errs.ParseSQLError(err)
		if errNum == "23505" {
			errs.ErrorHandler(ctx, 400, "Email Has Been Registered")
		} else {
			errs.ErrorHandler(ctx, 400, "Failed To Register")
		}
		return err
	}
	return nil
}

func (srv *userService) Login(ctx *gin.Context, src *requests.UserLoginRequest) (*responses.UserLoginResponse, error) {
	targetUser, err := srv.userRepo.GetOneUser(ctx, src.Email)
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Invalid Email Or Password")
		return nil, err
	}
	isPassword := utils.CheckPasswordHash(src.Password, targetUser.Password)
	if !isPassword {
		errs.ErrorHandler(ctx, 400, "Invalid Email Or Password")
		return nil, err
	}

	payload := domain.UserPayload{
		Email: targetUser.Email,
		ID:    targetUser.ID,
		Name:  targetUser.Name,
	}
	jwt, err := utils.SignToken(payload)
	if err != nil {
		errs.ErrorHandler(ctx, 400, "Failed To login")
		return nil, err
	}

	logged := responses.UserLoginResponse{
		ID:          targetUser.ID,
		Email:       targetUser.Email,
		Name:        targetUser.Name,
		AccessToken: jwt,
	}
	return &logged, nil
}
