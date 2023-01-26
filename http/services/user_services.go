package services

import (
	"context"
	"database/sql"
	"fmt"

	"gedebook.com/api/db"
	"gedebook.com/api/domain"
	"gedebook.com/api/domain/repository"
	"gedebook.com/api/dto/requests"
	"gedebook.com/api/errs"
	"gedebook.com/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type UserService interface {
	Register(ctx *gin.Context, src *domain.User) error
	Login(ctx *gin.Context, src *requests.UserLoginRequest) error
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
		errs.ErrorHandler(ctx, 400, "Failed To Register")
		return err
	}
	return nil
}

func (srv *userService) Login(ctx *gin.Context, src *requests.UserLoginRequest) error {
	targetUser, err := srv.userRepo.GetOneUser(ctx, src.Email)
	fmt.Println(targetUser)
	fmt.Println(err, ">>>")
	return nil
}
