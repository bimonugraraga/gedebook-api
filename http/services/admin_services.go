package services

import (
	"gedebook.com/api/domain/repository"
)

type AdminService interface {
	//!Attach Your Function Here
}
type adminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(adminRepo repository.AdminRepository) AdminService {
	return &adminService{
		adminRepo: adminRepo,
	}
}

//!Code Your Function Here
