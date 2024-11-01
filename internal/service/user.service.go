package service

import (
	"example.com/go-ecommerce-backend-api/internal/repo"
	"example.com/go-ecommerce-backend-api/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// // user repo
// func (us *UserService) GetInfoUserService() string {
// 	return us.userRepo.GetInfoUser()
// }

// INTERFACE_VERSION
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo *repo.IUserRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: &userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	//1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.
	}
	return response.ErrCodeSuccess
}